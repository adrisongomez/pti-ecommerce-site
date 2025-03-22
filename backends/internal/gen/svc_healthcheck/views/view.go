// Code generated by goa v3.20.0, DO NOT EDIT.
//
// svc-healthcheck views
//
// Command:
// $ goa gen github.com/adrisongomez/pti-ecommerce-site/backends/design -o
// ./internal

package views

import (
	goa "goa.design/goa/v3/pkg"
)

// HealthcheckResponse is the viewed result type that is projected based on a
// view.
type HealthcheckResponse struct {
	// Type to project
	Projected *HealthcheckResponseView
	// View to render
	View string
}

// HealthcheckResponseView is a type that runs validations on a projected type.
type HealthcheckResponseView struct {
	Status *string
}

var (
	// HealthcheckResponseMap is a map indexing the attribute names of
	// HealthcheckResponse by view name.
	HealthcheckResponseMap = map[string][]string{
		"default": {
			"status",
		},
	}
)

// ValidateHealthcheckResponse runs the validations defined on the viewed
// result type HealthcheckResponse.
func ValidateHealthcheckResponse(result *HealthcheckResponse) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateHealthcheckResponseView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []any{"default"})
	}
	return
}

// ValidateHealthcheckResponseView runs the validations defined on
// HealthcheckResponseView using the "default" view.
func ValidateHealthcheckResponseView(result *HealthcheckResponseView) (err error) {

	return
}
