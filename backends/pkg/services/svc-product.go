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

	dbVendor := model.Vendor()
	if dbVendor != nil {
		response.Vendor = (*productGen.Vendor)(MapVendorToVendorResponse(*dbVendor))
	}

	variants := []*ProductVariant{}
	medias := []*ProductMedia{}

	// logger.Info("datas", zap.Any("variants", len(dbVariants)), zap.Any("medias", len(dbProductMedias)))

	// if len(dbProductMedias) == 0 {
	// 	return &response
	// }

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
			updatedAtStr := value.String()
			media.UpdatedAt = &updatedAtStr
		}
		logger.Info("data", zap.Any("data", dbMedia))

		if value, ok := productMedia.Alt(); ok {
			strAlt := string(value)
			media.Alt = &strAlt
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
		}

		if value, ok := dbVariant.ColorHex(); ok {
			variant.ColorHex = &value
		}

		if value, ok := dbVariant.FeatureMediaID(); ok {
			variant.FeatureMediaID = &value
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
		db.Product.Vendor.Fetch(),
		db.Product.Medias.Fetch().With(db.ProductMedia.Media.Fetch()),
		db.Product.Variants.Fetch(),
		db.Product.MediasIn.Fetch(),
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
		db.Product.Vendor.Fetch(),
		db.Product.Medias.Fetch().With(db.ProductMedia.Media.Fetch()),
		db.Product.Variants.Fetch(),
		db.Product.MediasIn.Fetch(),
	).Exec(ctx)
	if err != nil {
		return nil, err
	}

	var productList ProductCollection = []*Product{}

	for _, p := range data {
		productList = append(productList, MapFromProductDbToOut(&p))
	}

	count := 0
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
	logger := zap.L()
	logger.Info("Create a product", methodLog, payloadLog)

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
		db.Product.Vendor.Link(db.Vendor.ID.Equals(payload.VendorID)),
		db.Product.Tags.Set(payload.Tags),
	).With().Exec(ctx)
	logger.Debug("Product is created", methodLog, zap.Any("record", dbProduct))
	if err != nil {
		logger.Error("Error trying to createProduct", methodLog, payloadLog, zap.Any("error", err))
		return nil, err
	}
	var txs []db.PrismaTransaction

	if payload.Medias != nil {
		for _, media := range payload.Medias {
			var alt string = ""
			if media.Alt != nil {
				alt = *media.Alt
			}
			txs = append(txs, p.client.ProductMedia.CreateOne(
				db.ProductMedia.SortNumber.Set(media.SortNumber),
				db.ProductMedia.ProductID.Set(dbProduct.ID),
				db.ProductMedia.MediaID.Set(media.MediaID),
				db.ProductMedia.Alt.Set(alt),
			).Tx())
		}
	}

	if payload.Variants != nil {
		for _, variant := range payload.Variants {
			colorHex := ""

			if variant.ColorHex != nil {
				colorHex = *variant.ColorHex
			}
			txs = append(txs, p.client.ProductVariant.CreateOne(
				db.ProductVariant.ColorName.Set(variant.ColorName),
				db.ProductVariant.Price.Set(variant.Price),
				db.ProductVariant.ProductID.Set(dbProduct.ID),
				db.ProductVariant.ColorHex.Set(colorHex),
			).Tx())
		}
	}

	err = p.client.Prisma.Transaction(txs...).Exec(ctx)
	if err != nil {
		p.client.Product.FindUnique(
			db.Product.ID.Equals(dbProduct.ID),
		).Delete().Exec(ctx)
		logger.Error("Error trying to stitching others elements", methodLog, payloadLog, zap.Any("error", err))
		return nil, err
	}

	prod, err := p.GetProductByID(ctx, &GetProductByIDPayload{ProductID: dbProduct.ID})
	if err != nil {
		logger.Error("Error trying to getProductId", methodLog, payloadLog, zap.Any("error", err))
		return nil, err
	}
	return prod, nil
}

func (p *ProductService) UpdateProductByID(ctx context.Context, payload *UpdateProductByIDPayload) (*Product, error) {
	logger := zap.L()
	methodSign := zap.String("method", "Product#UpdateProductByID")
	payloadLog := zap.Any("payload", payload)
	logger.Info("Update product by id got called with", methodSign, payloadLog)

	updatedProduct, err := p.client.Product.UpsertOne(db.Product.ID.Equals(payload.ProductID)).Update(
		db.Product.Title.Set(payload.Payload.Title),
		db.Product.Description.Set(payload.Payload.Description),
		db.Product.Status.Set(db.ProductStatus(payload.Payload.Status)),
		db.Product.Tags.Set(payload.Payload.Tags),
	).Exec(ctx)
	if err != nil {
		logger.Error("error on updating product", methodSign, payloadLog, zap.Error(err))
		return nil, err
	}
	logger.Error("Response by updating product", methodSign, zap.Any("updatedProduct", updatedProduct))
	return MapFromProductDbToOut(updatedProduct), nil
}

func (p *ProductService) DeleteProductByID(ctx context.Context, payload *DeleteProductByIDPayload) (bool, error) {
	_, err := p.client.Product.FindUnique(db.Product.ID.Equals(payload.ProductID)).Exec(ctx)
	if err != nil {
		return false, err
	}
	return true, nil
}

func NewProductService(client *db.PrismaClient) *ProductService {
	return &ProductService{client}
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
