package tools

import (
	"github.com/spf13/cast"
	"net/url"
)

const (
	Limit   = "limit"
	Page    = "page"
	OrderBy = "order_by"
	Sort    = "sort"

	DefaultLimit   = 10
	DefaultPage    = 1
	DefaultOrderBy = "id"
	DefaultSort    = "desc"
)

// Filter struct for filtering result
type Filter struct {
	Limit   int    `json:"limit"`
	Page    int    `json:"page"`
	Offset  int    `json:"-"`
	OrderBy string `json:"order_by,omitempty"`
	Sort    string `json:"sort,omitempty"`
}

// CalculateOffset method for calculate offset for pagination
func (f *Filter) CalculateOffset() {
	f.Offset = (f.Page - 1) * f.Limit
}

func ParseQueryParamsToFilter(urlValues url.Values) Filter {
	var filter = Filter{}

	// set default filter
	filter.Limit = DefaultLimit
	filter.Page = DefaultPage
	filter.OrderBy = DefaultOrderBy
	filter.Sort = DefaultSort

	// loop each url values
	for key, val := range urlValues {
		switch key {
		case Limit:
			if val[0] != "" {
				filter.Limit = cast.ToInt(val[0])
			}
		case Page:
			if val[0] != "" {
				filter.Page = cast.ToInt(val[0])
			}
		case OrderBy:
			if val[0] != "" {
				filter.OrderBy = val[0]
			}
		case Sort:
			if val[0] != "" {
				filter.Sort = val[0]
			}
		}
	}

	return filter
}
