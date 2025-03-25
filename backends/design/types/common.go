package types

import (
	. "goa.design/goa/v3/dsl"
	"goa.design/goa/v3/expr"
)

var PageInfo = ResultType("application/vnd.page-info+.json", func() {
	Description("Pagination information")
	Attributes(func() {
		Attribute("startCursor", Int, "The starting cursor for pagination")
		Attribute("endCursor", Int, "The ending cursor for pagination")
		Attribute("hasMore", Boolean, "Indicates if there are more results available")
		Attribute("totalResource", Int, "Total number of resources available")
	})
	Required("startCursor", "endCursor", "hasMore", "totalResource")
})

var fieldID = func() {
	Description("Key ID")
	Minimum(1)
	Example(10)
}

var fieldDatetime = func() {
	Description("Datetime")
	Format("date-time")
}

var TypeFooter = Type("TypeFooter", func() {
	Attribute("createdAt", String, "Date of creation", fieldDatetime)
	Attribute("updatedAt", String, "Last update date", fieldDatetime)

	Required("createdAt")
})

func PaginatedResult(name string, returnType *expr.ResultTypeExpr) *expr.ResultTypeExpr {
	return ResultType("application/vnd."+name+"+json", func() {
		Description("Paginated results")
		Attributes(func() {
			Attribute("data", ArrayOf(returnType), "Data")
			Attribute("pageInfo", PageInfo, "Pagination information")
			Required("data", "pageInfo")
		})
	})
}
