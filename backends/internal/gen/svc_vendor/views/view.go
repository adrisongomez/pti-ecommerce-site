// Code generated by goa v3.20.0, DO NOT EDIT.
//
// svc-vendor views
//
// Command:
// $ goa gen github.com/adrisongomez/pti-ecommerce-site/backends/design -o
// ./internal

package views

import (
	goa "goa.design/goa/v3/pkg"
)

// VendorList is the viewed result type that is projected based on a view.
type VendorList struct {
	// Type to project
	Projected *VendorListView
	// View to render
	View string
}

// Vendor is the viewed result type that is projected based on a view.
type Vendor struct {
	// Type to project
	Projected *VendorView
	// View to render
	View string
}

// VendorListView is a type that runs validations on a projected type.
type VendorListView struct {
	// Data
	Data VendorCollectionView
	// Pagination information
	PageInfo *PageInfoView
}

// VendorCollectionView is a type that runs validations on a projected type.
type VendorCollectionView []*VendorView

// VendorView is a type that runs validations on a projected type.
type VendorView struct {
	// Key ID
	ID   *int
	Name *string
}

// PageInfoView is a type that runs validations on a projected type.
type PageInfoView struct {
	// The starting cursor for pagination
	StartCursor *string
	// The ending cursor for pagination
	EndCursor *string
	// Indicates if there are more results available
	HasMore *bool
	// Total number of resources available
	TotalResource *int
}

var (
	// VendorListMap is a map indexing the attribute names of VendorList by view
	// name.
	VendorListMap = map[string][]string{
		"default": {
			"data",
			"pageInfo",
		},
	}
	// VendorMap is a map indexing the attribute names of Vendor by view name.
	VendorMap = map[string][]string{
		"default": {
			"id",
			"name",
		},
	}
	// VendorCollectionMap is a map indexing the attribute names of
	// VendorCollection by view name.
	VendorCollectionMap = map[string][]string{
		"default": {
			"id",
			"name",
		},
	}
	// PageInfoMap is a map indexing the attribute names of PageInfo by view name.
	PageInfoMap = map[string][]string{
		"default": {
			"startCursor",
			"endCursor",
			"hasMore",
			"totalResource",
		},
	}
)

// ValidateVendorList runs the validations defined on the viewed result type
// VendorList.
func ValidateVendorList(result *VendorList) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateVendorListView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []any{"default"})
	}
	return
}

// ValidateVendor runs the validations defined on the viewed result type Vendor.
func ValidateVendor(result *Vendor) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateVendorView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []any{"default"})
	}
	return
}

// ValidateVendorListView runs the validations defined on VendorListView using
// the "default" view.
func ValidateVendorListView(result *VendorListView) (err error) {

	if result.Data != nil {
		if err2 := ValidateVendorCollectionView(result.Data); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if result.PageInfo != nil {
		if err2 := ValidatePageInfoView(result.PageInfo); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateVendorCollectionView runs the validations defined on
// VendorCollectionView using the "default" view.
func ValidateVendorCollectionView(result VendorCollectionView) (err error) {
	for _, item := range result {
		if err2 := ValidateVendorView(item); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateVendorView runs the validations defined on VendorView using the
// "default" view.
func ValidateVendorView(result *VendorView) (err error) {
	if result.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "result"))
	}
	if result.ID != nil {
		if *result.ID < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("result.id", *result.ID, 1, true))
		}
	}
	return
}

// ValidatePageInfoView runs the validations defined on PageInfoView using the
// "default" view.
func ValidatePageInfoView(result *PageInfoView) (err error) {
	if result.StartCursor == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("startCursor", "result"))
	}
	if result.EndCursor == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("endCursor", "result"))
	}
	if result.HasMore == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("hasMore", "result"))
	}
	if result.TotalResource == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("totalResource", "result"))
	}
	return
}
