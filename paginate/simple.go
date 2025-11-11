package paginate

import (
	"golang.org/x/exp/constraints"
)

// DefaultPageSize is the default page size when not specified or invalid.
const DefaultPageSize = 10

// SimplePagination represents simple offset-limit pagination metadata.
type SimplePagination struct {
	// Total is the total number of items available.
	Total int64 `json:"total"`

	// Page is the current page number (1-indexed).
	Page int `json:"page"`

	// PageSize is the number of items per page.
	PageSize int `json:"page_size"`

	// LastPage is the last page number.
	LastPage int64 `json:"last_page"`
}

// NewSimplePagination creates pagination metadata from total items, page number, and page size.
func NewSimplePagination(total int64, page int, pageSize int) *SimplePagination {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = DefaultPageSize
	}

	ps := int64(pageSize)
	lastPage := int64(1)
	if total > ps {
		lastPage = (total + ps - 1) / ps
	}

	return &SimplePagination{
		Total:    total,
		Page:     page,
		PageSize: pageSize,
		LastPage: lastPage,
	}
}

// GetOffset calculates the offset for the given page and page size.
func GetOffset[T constraints.Integer](page T, pageSize T) T {
	return (page - 1) * pageSize
}

// GetLimit returns the page size (limit), or DefaultPageSize if invalid.
func GetLimit[T constraints.Integer](pageSize T) T {
	if pageSize <= 0 {
		return T(DefaultPageSize)
	}
	return pageSize
}
