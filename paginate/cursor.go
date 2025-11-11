package paginate

// CursorPagination represents cursor-based pagination metadata for "load more" pattern.
// It stores the current cursor, next cursor, and whether more items are available.
type CursorPagination struct {
	// Cursor is the cursor pointing to the current position in the dataset.
	Cursor string `json:"cursor,omitempty"`

	// NextCursor is the cursor for fetching the next batch of items.
	// Empty string indicates no more items available.
	NextCursor string `json:"next_cursor,omitempty"`

	// HasMore indicates whether there are more items available after the current batch.
	HasMore bool `json:"has_more"`
}

// NewCursorPagination creates a new CursorPagination instance.
// hasMore indicates whether there are more items available.
func NewCursorPagination(cursor string, nextCursor string, hasMore bool) *CursorPagination {
	return &CursorPagination{
		Cursor:     cursor,
		NextCursor: nextCursor,
		HasMore:    hasMore,
	}
}

// IsLastPage checks if this is the last page (no more items available).
func (cp *CursorPagination) IsLastPage() bool {
	return !cp.HasMore
}
