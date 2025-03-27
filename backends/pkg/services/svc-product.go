package services

import (
	"context"
	"fmt"

	"github.com/adrisongomez/pti-ecommerce-site/backends/databases/db"
	. "github.com/adrisongomez/pti-ecommerce-site/backends/internal/gen/svc_products"
	internalUtils "github.com/adrisongomez/pti-ecommerce-site/backends/internal/utils"
	"github.com/adrisongomez/pti-ecommerce-site/backends/pkg/utils"
	"github.com/google/uuid"
	"go.uber.org/zap"

	productGenhttp "github.com/adrisongomez/pti-ecommerce-site/backends/internal/gen/http/svc_products/server"
	productGen "github.com/adrisongomez/pti-ecommerce-site/backends/internal/gen/svc_products"
	goahttp "goa.design/goa/v3/http"
)

type ProductService struct {
	client *db.PrismaClient
	logger *zap.Logger
}

var Connections = []db.ProductRelationWith{
	db.Product.Medias.Fetch().With(db.ProductMedia.Media.Fetch()),
	db.Product.Variants.Fetch(),
	db.Product.MediasIn.Fetch(),
}

func MapFromProductDbToOut(model *db.ProductModel) *Product {
	logger := zap.L()
	response := Product{
		ID:          model.ID,
		Title:       model.Title,
		Description: model.Description,
		Handle:      model.Handler,
		Tags:        model.Tags,
		CreatedAt:   model.CreatedAt.String(),
		Status:      productGen.ProductStatus(model.Status),
	}
	if value, ok := model.UpdatedAt(); ok {
		updatedAtStr := value.String()
		response.UpdatedAt = &updatedAtStr
	}

	variants := []*ProductVariant{}
	medias := []*ProductMedia{}

	dbProductMedias := model.Medias()
	dbVariants := model.Variants()

	for _, productMedia := range dbProductMedias {
		dbMedia, ok := productMedia.Media()
		logger.Info("Db Media",
			zap.Any("check", ok),
			zap.Any("dbMedia", dbMedia),
			zap.Any("productMedia", productMedia),
		)
		if ok == false {
			continue
		}
		media := ProductMedia{
			ID:         productMedia.ID,
			URL:        utils.GetResourceURL(dbMedia.Bucket, "us-east-1", dbMedia.Key),
			MediaID:    productMedia.MediaID,
			SortNumber: productMedia.SortNumber,
			Alt:        nil,
			MediaType:  MediaType(dbMedia.Type),
			CreatedAt:  dbMedia.CreatedAt.String(),
			UpdatedAt:  nil,
		}
		if value, ok := dbMedia.UpdatedAt(); ok {
			media.UpdatedAt = internalUtils.StringRef(value.String())
		}
		logger.Info("data", zap.Any("data", dbMedia))

		if value, ok := productMedia.Alt(); ok {
			media.Alt = internalUtils.StringRef(value)

		}
		medias = append(medias, &media)
	}

	for _, dbVariant := range dbVariants {
		variant := ProductVariant{
			ID:             dbVariant.ID,
			ColorName:      dbVariant.ColorName,
			Price:          dbVariant.Price,
			ColorHex:       nil,
			FeatureMediaID: nil,
			CreatedAt:      dbVariant.CreatedAt.String(),
		}

		if value, ok := dbVariant.ColorHex(); ok {
			variant.ColorHex = internalUtils.StringRef(value)
		}

		if value, ok := dbVariant.FeatureMediaID(); ok {
			variant.FeatureMediaID = &value
		}

		if value, ok := dbVariant.UpdatedAt(); ok {
			variant.UpdatedAt = internalUtils.StringRef(value.String())
		}

		variants = append(variants, &variant)
	}
	logger.Debug("After media", zap.Any("medias", medias))
	response.Medias = medias
	response.Variants = variants
	return &response
}

func (p *ProductService) GetProductByID(ctx context.Context, payload *GetProductByIDPayload) (*Product, error) {
	logger := zap.L()

	method := zap.String("method", "ProductService#GetProductID")
	productIdLog := zap.Int("productId", payload.ProductID)
	logger.Info("Looks for a product",
		method,
		productIdLog,
	)

	dbProduct, err := p.client.Product.FindFirst(db.Product.ID.Equals(payload.ProductID)).With(
		Connections...,
	).Exec(ctx)
	if err != nil {
		logger.Error("Error trying to getProductId", method, productIdLog, zap.Any("error", err))
		return nil, err
	}
	logger.Debug("Response for getting product", method, productIdLog, zap.Any("response", dbProduct))

	return MapFromProductDbToOut(dbProduct), nil
}

func (p *ProductService) count(ctx context.Context) (int, error) {
	var resp []struct {
		Count db.BigInt `json:"count"`
	}
	err := p.client.Prisma.QueryRaw("SELECT count(*) FROM project.products").Exec(ctx, &resp)
	if err != nil {
		return 0, err
	}
	if len(resp) == 0 {
		return 0, nil
	}
	count := int(resp[0].Count)
	return count, nil
}

func (p *ProductService) ListProduct(ctx context.Context, payload *ListProductPayload) (*ProductsList, error) {
	data, err := p.client.Product.FindMany().Take(payload.PageSize).Skip(payload.After).With(
		Connections...,
	).Exec(ctx)

	if err != nil {
		return nil, err
	}

	productList := []*Product{}

	for _, p := range data {
		productList = append(productList, MapFromProductDbToOut(&p))
	}

	count, err := p.count(ctx)
	if err != nil {
		return nil, err
	}
	nextPageCursor := internalUtils.MinInt(count, payload.After+payload.PageSize)

	pageInfo := &PageInfo{
		StartCursor:   payload.After,
		EndCursor:     nextPageCursor,
		TotalResource: count,
		HasMore:       nextPageCursor < count,
	}

	response := &ProductsList{
		Data:     productList,
		PageInfo: pageInfo,
	}

	return response, nil
}

func (p *ProductService) CreateProduct(ctx context.Context, payload *ProductInput) (*Product, error) {
	methodLog := zap.String("method", "ProducServicet#CreateProduct")
	payloadLog := zap.Any("payload", payload)
	p.logger.Info("Create a product", methodLog, payloadLog)

	var handle string
	if payload.Handle == nil {
		handle = uuid.NewString()
	} else {
		handle = *payload.Handle

	}
	dbProduct, err := p.client.Product.CreateOne(
		db.Product.Title.Set(payload.Title),
		db.Product.Description.Set(payload.Description),
		db.Product.Handler.Set(handle),
		db.Product.Status.Set(db.ProductStatus(payload.Status)),
		db.Product.Tags.Set(payload.Tags),
	).With(
		Connections...,
	).Exec(ctx)
	p.logger.Debug("Product is created", methodLog, zap.Any("record", dbProduct))
	if err != nil {
		p.logger.Error("Error trying to createProduct", methodLog, payloadLog, zap.Any("error", err))
		return nil, err
	}
	var txs []db.PrismaTransaction

	if payload.Medias != nil {
		for _, media := range payload.Medias {
			changes := []db.ProductMediaSetParam{}
			if media.Alt != nil {
				changes = append(changes, db.ProductMedia.Alt.Set(*media.Alt))
			}
			changes = append(changes,
				db.ProductMedia.ProductID.Set(dbProduct.ID),
				db.ProductMedia.MediaID.Set(media.MediaID),
			)
			txs = append(txs, p.client.ProductMedia.CreateOne(
				db.ProductMedia.SortNumber.Set(media.SortNumber),
				changes...,
			).Tx())
		}
	}

	if payload.Variants != nil {
		for _, variant := range payload.Variants {
			changes := []db.ProductVariantSetParam{}
			if variant.ColorHex != nil {
				changes = append(changes, db.ProductVariant.ColorHex.Set(*variant.ColorHex))
			}
			changes = append(changes,
				db.ProductVariant.ProductID.Set(dbProduct.ID),
			)
			txs = append(txs, p.client.ProductVariant.CreateOne(
				db.ProductVariant.ColorName.Set(variant.ColorName),
				db.ProductVariant.Price.Set(variant.Price),
				changes...,
			).Tx())
		}
	}

	err = p.client.Prisma.Transaction(txs...).Exec(ctx)
	if err != nil {
		p.client.Product.FindUnique(
			db.Product.ID.Equals(dbProduct.ID),
		).Delete().Exec(ctx)
		p.logger.Error("Error trying to stitching others elements", methodLog, payloadLog, zap.Any("error", err))
		return nil, err
	}

	prod, err := p.GetProductByID(ctx, &GetProductByIDPayload{ProductID: dbProduct.ID})
	if err != nil {
		p.logger.Error("Error trying to getProductId", methodLog, payloadLog, zap.Any("error", err))
		return nil, err
	}
	return prod, nil
}

func (p *ProductService) UpdateProductByID(ctx context.Context, payload *UpdateProductByIDPayload) (*Product, error) {
	methodSign := zap.String("method", "Product#UpdateProductByID")
	payloadLog := zap.Any("payload", payload)
	p.logger.Info("Update product by id got called with", methodSign, payloadLog)

	updatedProduct, err := p.client.Product.UpsertOne(db.Product.ID.Equals(payload.ProductID)).Update(
		db.Product.Title.Set(payload.Payload.Title),
		db.Product.Description.Set(payload.Payload.Description),
		db.Product.Status.Set(db.ProductStatus(payload.Payload.Status)),
		db.Product.Tags.Set(payload.Payload.Tags),
	).Exec(ctx)
	if err != nil {
		p.logger.Error("error on updating product", methodSign, payloadLog, zap.Error(err))
		return nil, err
	}
	p.logger.Error("Response by updating product", methodSign, zap.Any("updatedProduct", updatedProduct))
	return MapFromProductDbToOut(updatedProduct), nil
}

func (p *ProductService) DeleteProductByID(ctx context.Context, payload *DeleteProductByIDPayload) (bool, error) {
	methodLog := zap.String("method", "ProductService#DeleteProductByID")
	_, err := p.client.Product.FindUnique(db.Product.ID.Equals(payload.ProductID)).Delete().Exec(ctx)
	if err != nil {
		p.logger.Error("Error on deleting product", methodLog, zap.Error(err))
		return false, err
	}
	return true, nil
}
func (p *ProductService) UpsertVariant(ctx context.Context, productId int, payload *ProductVariantUpsertInput) (*Product, error) {
	methodLog := zap.String("methodName", "ProductService#AddVariant")
	changes := []db.ProductVariantSetParam{}
	if payload.ColorHex != nil {
		changes = append(changes, db.ProductVariant.ColorHex.Set(*payload.ColorHex))
	}
	p.logger.Info("Upsert product variant got called", methodLog, zap.Any("payload", payload))
	id := payload.ID
	if id == nil {
		variant, err := p.client.ProductVariant.CreateOne(
			db.ProductVariant.ColorName.Set(payload.ColorName),
			db.ProductVariant.Price.Set(payload.Price),
			changes...,
		).Exec(ctx)
		if err != nil {
			p.logger.Error("Error on adding variant", methodLog, zap.Error(err))
			return nil, err
		}
		p.logger.Info("Added variant", methodLog, zap.Any("newVariant", variant))
	} else {

		changes = append(
			changes,
			db.ProductVariant.ColorName.Set(payload.ColorName),
			db.ProductVariant.Price.Set(payload.Price),
		)
		variant, err := p.client.ProductVariant.
			FindUnique(db.ProductVariant.ID.Equals(*id)).
			Update(changes...).Exec(ctx)
		if err != nil {
			p.logger.Error("Error on adding variant", methodLog, zap.Error(err))
			return nil, err
		}
		p.logger.Info("Added variant", methodLog, zap.Any("newVariant", variant))
	}

	return p.GetProductByID(ctx, &GetProductByIDPayload{ProductID: productId})
}

func (p *ProductService) RemoveVariants(ctx context.Context, productId int, ids []int) (*Product, error) {
	methodLog := zap.String("method", "ProductService#RemoveVariant")
	p.logger.Info("Remove variant got called", methodLog, zap.Any("payload", ids))
	_, err := p.client.ProductVariant.FindMany(
		db.ProductVariant.ID.In(ids),
	).Delete().Exec(ctx)
	if err != nil {
		p.logger.Error("Error on remove variant", methodLog, zap.Error(err))
		return nil, err
	}
	p.logger.Info("Removed", methodLog, zap.Any("variantIds", ids))
	return p.GetProductByID(ctx, &GetProductByIDPayload{ProductID: productId})
}

func (p *ProductService) AddMedia(ctx context.Context, productId int, payload *ProductMediaInput) (*Product, error) {
	methodLog := zap.String("method", "ProductService#AddMedia")
	p.logger.Info("Add media got called", methodLog, zap.Any("payload", payload))

	changes := []db.ProductMediaSetParam{
		db.ProductMedia.MediaID.Set(payload.MediaID),
	}

	if payload.Alt != nil {
		changes = append(changes, db.ProductMedia.Alt.Set(*payload.Alt))
	}
	newMedia, err := p.client.ProductMedia.CreateOne(
		db.ProductMedia.SortNumber.Set(payload.SortNumber),
		changes...,
	).Exec(ctx)

	if err != nil {
		p.logger.Error("Error on adding media", methodLog, zap.Error(err))
		return nil, err
	}
	p.logger.Info("Added product media", methodLog, zap.Any("newMedia", newMedia))

	return nil, nil
}

func (p *ProductService) RemoveMedia(ctx context.Context, productId int, payload []int) (*Product, error) {
	methodLog := zap.String("method", "ProductService#RemoveMedia")
	p.logger.Info("Add media got called", methodLog, zap.Any("payload", payload))
	removedMedia, err := p.client.ProductMedia.FindMany(
		db.ProductMedia.ID.In(payload),
	).Delete().Exec(ctx)
	if err != nil {
		p.logger.Error("Error on removing media", methodLog, zap.Error(err))
		return nil, err
	}
	p.logger.Info("Remove product media", methodLog, zap.Any("removedMedia", removedMedia))
	return p.GetProductByID(ctx, &GetProductByIDPayload{ProductID: productId})
}

func NewProductService(client *db.PrismaClient) *ProductService {
	logger := zap.L()
	return &ProductService{client, logger}
}

func MountProductSVC(mux goahttp.Muxer, svc *ProductService) {
	endpoints := productGen.NewEndpoints(svc)
	req := goahttp.RequestDecoder
	res := goahttp.ResponseEncoder

	handler := productGenhttp.New(endpoints, mux, req, res, nil, nil)
	productGenhttp.Mount(mux, handler)

	go func() {
		for _, mount := range handler.Mounts {
			zap.L().Info(fmt.Sprintf("%q mounted on %s %s\n", mount.Method, mount.Verb, mount.Pattern))
		}
	}()
}
