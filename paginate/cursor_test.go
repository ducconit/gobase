package paginate

import (
	"testing"
)

func TestNewCursorPagination(t *testing.T) {
	tests := []struct {
		name        string
		cursor      string
		nextCursor  string
		hasMore     bool
		wantHasMore bool
	}{
		{
			name:        "has more items",
			cursor:      "cursor1",
			nextCursor:  "cursor2",
			hasMore:     true,
			wantHasMore: true,
		},
		{
			name:        "no more items",
			cursor:      "cursor1",
			nextCursor:  "",
			hasMore:     false,
			wantHasMore: false,
		},
		{
			name:        "empty cursor",
			cursor:      "",
			nextCursor:  "cursor2",
			hasMore:     true,
			wantHasMore: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewCursorPagination(tt.cursor, tt.nextCursor, tt.hasMore)
			if got.Cursor != tt.cursor || got.NextCursor != tt.nextCursor || got.HasMore != tt.wantHasMore {
				t.Errorf("NewCursorPagination() = {Cursor: %s, NextCursor: %s, HasMore: %v}, want {Cursor: %s, NextCursor: %s, HasMore: %v}",
					got.Cursor, got.NextCursor, got.HasMore, tt.cursor, tt.nextCursor, tt.wantHasMore)
			}
		})
	}
}

func TestIsLastPage(t *testing.T) {
	tests := []struct {
		name string
		cp   *CursorPagination
		want bool
	}{
		{
			name: "has more items",
			cp: &CursorPagination{
				HasMore: true,
			},
			want: false,
		},
		{
			name: "last page",
			cp: &CursorPagination{
				HasMore: false,
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.cp.IsLastPage()
			if got != tt.want {
				t.Errorf("IsLastPage() = %v, want %v", got, tt.want)
			}
		})
	}
}

