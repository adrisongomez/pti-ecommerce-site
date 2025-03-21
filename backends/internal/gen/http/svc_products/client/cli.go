// Code generated by goa v3.20.0, DO NOT EDIT.
//
// svc-products HTTP client CLI support package
//
// Command:
// $ goa gen github.com/adrisongomez/pti-ecommerce-site/backends/design -o
// ./internal

package client

import (
	"encoding/json"
	"fmt"
	"strconv"

	svcproducts "github.com/adrisongomez/pti-ecommerce-site/backends/internal/gen/svc_products"
	goa "goa.design/goa/v3/pkg"
)

// BuildListProductPayload builds the payload for the svc-products listProduct
// endpoint from CLI flags.
func BuildListProductPayload(svcProductsListProductPageSize string, svcProductsListProductAfter string) (*svcproducts.ListProductPayload, error) {
	var err error
	var pageSize int
	{
		if svcProductsListProductPageSize != "" {
			var v int64
			v, err = strconv.ParseInt(svcProductsListProductPageSize, 10, strconv.IntSize)
			pageSize = int(v)
			if err != nil {
				return nil, fmt.Errorf("invalid value for pageSize, must be INT")
			}
			if pageSize < 10 {
				err = goa.MergeErrors(err, goa.InvalidRangeError("pageSize", pageSize, 10, true))
			}
			if pageSize > 100 {
				err = goa.MergeErrors(err, goa.InvalidRangeError("pageSize", pageSize, 100, false))
			}
			if err != nil {
				return nil, err
			}
		}
	}
	var after *int
	{
		if svcProductsListProductAfter != "" {
			var v int64
			v, err = strconv.ParseInt(svcProductsListProductAfter, 10, strconv.IntSize)
			val := int(v)
			after = &val
			if err != nil {
				return nil, fmt.Errorf("invalid value for after, must be INT")
			}
		}
	}
	v := &svcproducts.ListProductPayload{}
	v.PageSize = pageSize
	v.After = after

	return v, nil
}

// BuildGetProductByIDPayload builds the payload for the svc-products
// getProductById endpoint from CLI flags.
func BuildGetProductByIDPayload(svcProductsGetProductByIDProductID string) (*svcproducts.GetProductByIDPayload, error) {
	var err error
	var productID int
	{
		var v int64
		v, err = strconv.ParseInt(svcProductsGetProductByIDProductID, 10, strconv.IntSize)
		productID = int(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for productID, must be INT")
		}
	}
	v := &svcproducts.GetProductByIDPayload{}
	v.ProductID = productID

	return v, nil
}

// BuildCreateProductPayload builds the payload for the svc-products
// createProduct endpoint from CLI flags.
func BuildCreateProductPayload(svcProductsCreateProductBody string) (*svcproducts.ProductInput, error) {
	var err error
	var body CreateProductRequestBody
	{
		err = json.Unmarshal([]byte(svcProductsCreateProductBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"description\": \"Consectetur enim numquam iste dolorem possimus voluptatem.\",\n      \"handle\": \"Sunt voluptatem doloribus tempora neque reprehenderit eveniet.\",\n      \"medias\": [\n         {\n            \"alt\": \"Sed quis.\",\n            \"mediaId\": \"Tenetur qui consequuntur.\",\n            \"sortNumber\": 2548002682993232380\n         },\n         {\n            \"alt\": \"Sed quis.\",\n            \"mediaId\": \"Tenetur qui consequuntur.\",\n            \"sortNumber\": 2548002682993232380\n         },\n         {\n            \"alt\": \"Sed quis.\",\n            \"mediaId\": \"Tenetur qui consequuntur.\",\n            \"sortNumber\": 2548002682993232380\n         },\n         {\n            \"alt\": \"Sed quis.\",\n            \"mediaId\": \"Tenetur qui consequuntur.\",\n            \"sortNumber\": 2548002682993232380\n         }\n      ],\n      \"status\": \"DRAFT\",\n      \"tags\": [\n         \"Quas rerum voluptatem quas dolorem et.\",\n         \"Esse nemo voluptatem est.\",\n         \"Nam unde sapiente et inventore.\",\n         \"Neque ut.\"\n      ],\n      \"title\": \"Ut ut in eaque in in veniam.\",\n      \"variants\": [\n         {\n            \"colorHex\": \"Molestiae et soluta animi aliquam eaque.\",\n            \"colorName\": \"Veritatis impedit sequi tenetur numquam ad.\",\n            \"price\": 9013296666552782995\n         },\n         {\n            \"colorHex\": \"Molestiae et soluta animi aliquam eaque.\",\n            \"colorName\": \"Veritatis impedit sequi tenetur numquam ad.\",\n            \"price\": 9013296666552782995\n         },\n         {\n            \"colorHex\": \"Molestiae et soluta animi aliquam eaque.\",\n            \"colorName\": \"Veritatis impedit sequi tenetur numquam ad.\",\n            \"price\": 9013296666552782995\n         },\n         {\n            \"colorHex\": \"Molestiae et soluta animi aliquam eaque.\",\n            \"colorName\": \"Veritatis impedit sequi tenetur numquam ad.\",\n            \"price\": 9013296666552782995\n         }\n      ],\n      \"vendorId\": \"Recusandae in aliquid accusamus occaecati.\"\n   }'")
		}
		if body.Tags == nil {
			err = goa.MergeErrors(err, goa.MissingFieldError("tags", "body"))
		}
		if body.Status != nil {
			if !(*body.Status == "ACTIVE" || *body.Status == "DRAFT") {
				err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.status", *body.Status, []any{"ACTIVE", "DRAFT"}))
			}
		}
		if err != nil {
			return nil, err
		}
	}
	v := &svcproducts.ProductInput{
		Title:       body.Title,
		Description: body.Description,
		Handle:      body.Handle,
		VendorID:    body.VendorID,
	}
	if body.Status != nil {
		status := svcproducts.ProductStatus(*body.Status)
		v.Status = &status
	}
	if body.Tags != nil {
		v.Tags = make([]string, len(body.Tags))
		for i, val := range body.Tags {
			v.Tags[i] = val
		}
	} else {
		v.Tags = []string{}
	}
	if body.Variants != nil {
		v.Variants = make([]*svcproducts.ProductVariantInput, len(body.Variants))
		for i, val := range body.Variants {
			v.Variants[i] = marshalProductVariantInputRequestBodyToSvcproductsProductVariantInput(val)
		}
	}
	if body.Medias != nil {
		v.Medias = make([]*svcproducts.ProductMediaInput, len(body.Medias))
		for i, val := range body.Medias {
			v.Medias[i] = marshalProductMediaInputRequestBodyToSvcproductsProductMediaInput(val)
		}
	}

	return v, nil
}

// BuildUpdateProductByIDPayload builds the payload for the svc-products
// updateProductById endpoint from CLI flags.
func BuildUpdateProductByIDPayload(svcProductsUpdateProductByIDBody string, svcProductsUpdateProductByIDProductID string) (*svcproducts.UpdateProductByIDPayload, error) {
	var err error
	var body UpdateProductByIDRequestBody
	{
		err = json.Unmarshal([]byte(svcProductsUpdateProductByIDBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"description\": \"Beatae dignissimos repellat.\",\n      \"handle\": \"Sint sint.\",\n      \"medias\": [\n         {\n            \"alt\": \"At quaerat non perferendis odit commodi aut.\",\n            \"mediaId\": \"Nostrum distinctio quia aut animi quis quod.\",\n            \"sortNumber\": 9107037954132454170\n         },\n         {\n            \"alt\": \"At quaerat non perferendis odit commodi aut.\",\n            \"mediaId\": \"Nostrum distinctio quia aut animi quis quod.\",\n            \"sortNumber\": 9107037954132454170\n         }\n      ],\n      \"status\": \"DRAFT\",\n      \"tags\": [\n         \"Iste consequatur minima similique nulla tempore.\",\n         \"Perferendis distinctio non.\",\n         \"Ut numquam voluptatem quia molestiae eligendi consequuntur.\",\n         \"Vero aut.\"\n      ],\n      \"title\": \"Qui ut accusantium dolorum sit.\",\n      \"variants\": [\n         {\n            \"colorHex\": \"Ut natus porro eaque eius sint.\",\n            \"colorName\": \"Dolorem molestias nam ad.\",\n            \"price\": 3281393361508820864\n         },\n         {\n            \"colorHex\": \"Ut natus porro eaque eius sint.\",\n            \"colorName\": \"Dolorem molestias nam ad.\",\n            \"price\": 3281393361508820864\n         },\n         {\n            \"colorHex\": \"Ut natus porro eaque eius sint.\",\n            \"colorName\": \"Dolorem molestias nam ad.\",\n            \"price\": 3281393361508820864\n         }\n      ],\n      \"vendorId\": \"Error autem enim.\"\n   }'")
		}
		if body.Tags == nil {
			err = goa.MergeErrors(err, goa.MissingFieldError("tags", "body"))
		}
		if body.Status != nil {
			if !(*body.Status == "ACTIVE" || *body.Status == "DRAFT") {
				err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.status", *body.Status, []any{"ACTIVE", "DRAFT"}))
			}
		}
		if err != nil {
			return nil, err
		}
	}
	var productID int
	{
		var v int64
		v, err = strconv.ParseInt(svcProductsUpdateProductByIDProductID, 10, strconv.IntSize)
		productID = int(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for productID, must be INT")
		}
	}
	v := &svcproducts.ProductInput{
		Title:       body.Title,
		Description: body.Description,
		Handle:      body.Handle,
		VendorID:    body.VendorID,
	}
	if body.Status != nil {
		status := svcproducts.ProductStatus(*body.Status)
		v.Status = &status
	}
	if body.Tags != nil {
		v.Tags = make([]string, len(body.Tags))
		for i, val := range body.Tags {
			v.Tags[i] = val
		}
	} else {
		v.Tags = []string{}
	}
	if body.Variants != nil {
		v.Variants = make([]*svcproducts.ProductVariantInput, len(body.Variants))
		for i, val := range body.Variants {
			v.Variants[i] = marshalProductVariantInputRequestBodyRequestBodyToSvcproductsProductVariantInput(val)
		}
	}
	if body.Medias != nil {
		v.Medias = make([]*svcproducts.ProductMediaInput, len(body.Medias))
		for i, val := range body.Medias {
			v.Medias[i] = marshalProductMediaInputRequestBodyRequestBodyToSvcproductsProductMediaInput(val)
		}
	}
	res := &svcproducts.UpdateProductByIDPayload{
		Payload: v,
	}
	res.ProductID = productID

	return res, nil
}

// BuildDeleteProductByIDPayload builds the payload for the svc-products
// deleteProductById endpoint from CLI flags.
func BuildDeleteProductByIDPayload(svcProductsDeleteProductByIDProductID string) (*svcproducts.DeleteProductByIDPayload, error) {
	var err error
	var productID int
	{
		var v int64
		v, err = strconv.ParseInt(svcProductsDeleteProductByIDProductID, 10, strconv.IntSize)
		productID = int(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for productID, must be INT")
		}
	}
	v := &svcproducts.DeleteProductByIDPayload{}
	v.ProductID = productID

	return v, nil
}
