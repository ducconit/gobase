package httputil

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ducconit/gobase/paginate"
	"github.com/gin-gonic/gin"
)

type TestItem struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func TestSimplePagination(t *testing.T) {
	tests := []struct {
		name      string
		items     []TestItem
		total     int64
		page      int
		pageSize  int
		message   string
		wantCode  string
		wantTotal int64
		wantPage  int
	}{
		{
			name: "first page",
			items: []TestItem{
				{ID: "1", Name: "Item 1"},
				{ID: "2", Name: "Item 2"},
			},
			total:     20,
			page:      1,
			pageSize:  10,
			message:   "Items fetched",
			wantCode:  "0",
			wantTotal: 20,
			wantPage:  1,
		},
		{
			name: "middle page",
			items: []TestItem{
				{ID: "11", Name: "Item 11"},
				{ID: "12", Name: "Item 12"},
			},
			total:     30,
			page:      2,
			pageSize:  10,
			message:   "Items fetched",
			wantCode:  "0",
			wantTotal: 30,
			wantPage:  2,
		},
		{
			name:      "empty items",
			items:     []TestItem{},
			total:     0,
			page:      1,
			pageSize:  10,
			message:   "No items",
			wantCode:  "0",
			wantTotal: 0,
			wantPage:  1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request = httptest.NewRequest("GET", "/", nil)

			SimplePagination(c, tt.items, tt.total, tt.page, tt.pageSize, tt.message)

			if w.Code != http.StatusOK {
				t.Errorf("SimplePagination() status = %d, want %d", w.Code, http.StatusOK)
			}

			var resp JsonResponse[[]TestItem, *paginate.SimplePagination]
			if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
				t.Errorf("SimplePagination() failed to unmarshal response: %v", err)
			}

			if resp.Code != tt.wantCode {
				t.Errorf("SimplePagination() code = %s, want %s", resp.Code, tt.wantCode)
			}

			if resp.Extra == nil {
				t.Errorf("SimplePagination() extra is nil")
				return
			}

			if resp.Extra.Total != tt.wantTotal {
				t.Errorf("SimplePagination() total = %d, want %d", resp.Extra.Total, tt.wantTotal)
			}

			if resp.Extra.Page != tt.wantPage {
				t.Errorf("SimplePagination() page = %d, want %d", resp.Extra.Page, tt.wantPage)
			}

			if len(resp.Data) != len(tt.items) {
				t.Errorf("SimplePagination() data length = %d, want %d", len(resp.Data), len(tt.items))
			}
		})
	}
}

func TestCursorPagination(t *testing.T) {
	tests := []struct {
		name       string
		items      []TestItem
		cursor     string
		nextCursor string
		hasMore    bool
		message    string
		wantCode   string
		wantCursor string
	}{
		{
			name: "has more items",
			items: []TestItem{
				{ID: "1", Name: "Item 1"},
				{ID: "2", Name: "Item 2"},
			},
			cursor:     "cursor1",
			nextCursor: "cursor2",
			hasMore:    true,
			message:    "Items fetched",
			wantCode:   "0",
			wantCursor: "cursor1",
		},
		{
			name: "last page",
			items: []TestItem{
				{ID: "9", Name: "Item 9"},
				{ID: "10", Name: "Item 10"},
			},
			cursor:     "cursor8",
			nextCursor: "",
			hasMore:    false,
			message:    "Items fetched",
			wantCode:   "0",
			wantCursor: "cursor8",
		},
		{
			name:       "empty items",
			items:      []TestItem{},
			cursor:     "",
			nextCursor: "",
			hasMore:    false,
			message:    "No items",
			wantCode:   "0",
			wantCursor: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request = httptest.NewRequest("GET", "/", nil)

			CursorPagination(c, tt.items, tt.cursor, tt.nextCursor, tt.hasMore, tt.message)

			if w.Code != http.StatusOK {
				t.Errorf("CursorPagination() status = %d, want %d", w.Code, http.StatusOK)
			}

			var resp JsonResponse[[]TestItem, *paginate.CursorPagination]
			if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
				t.Errorf("CursorPagination() failed to unmarshal response: %v", err)
			}

			if resp.Code != tt.wantCode {
				t.Errorf("CursorPagination() code = %s, want %s", resp.Code, tt.wantCode)
			}

			if resp.Extra == nil {
				t.Errorf("CursorPagination() extra is nil")
				return
			}

			if resp.Extra.Cursor != tt.wantCursor {
				t.Errorf("CursorPagination() cursor = %s, want %s", resp.Extra.Cursor, tt.wantCursor)
			}

			if resp.Extra.HasMore != tt.hasMore {
				t.Errorf("CursorPagination() hasMore = %v, want %v", resp.Extra.HasMore, tt.hasMore)
			}

			if len(resp.Data) != len(tt.items) {
				t.Errorf("CursorPagination() data length = %d, want %d", len(resp.Data), len(tt.items))
			}
		})
	}
}
