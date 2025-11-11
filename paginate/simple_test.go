package paginate

import (
	"testing"
)

func TestNewSimplePagination(t *testing.T) {
	tests := []struct {
		name     string
		total    int64
		page     int
		pageSize int
		want     *SimplePagination
	}{
		{
			name:     "normal case",
			total:    100,
			page:     1,
			pageSize: 10,
			want: &SimplePagination{
				Total:    100,
				Page:     1,
				PageSize: 10,
				LastPage: 10,
			},
		},
		{
			name:     "page 5 of 10",
			total:    100,
			page:     5,
			pageSize: 10,
			want: &SimplePagination{
				Total:    100,
				Page:     5,
				PageSize: 10,
				LastPage: 10,
			},
		},
		{
			name:     "invalid page < 1",
			total:    100,
			page:     0,
			pageSize: 10,
			want: &SimplePagination{
				Total:    100,
				Page:     1,
				PageSize: 10,
				LastPage: 10,
			},
		},
		{
			name:     "invalid pageSize < 1",
			total:    100,
			page:     1,
			pageSize: 0,
			want: &SimplePagination{
				Total:    100,
				Page:     1,
				PageSize: 10,
				LastPage: 10,
			},
		},
		{
			name:     "empty result",
			total:    0,
			page:     1,
			pageSize: 10,
			want: &SimplePagination{
				Total:    0,
				Page:     1,
				PageSize: 10,
				LastPage: 1,
			},
		},
		{
			name:     "total less than pageSize",
			total:    5,
			page:     1,
			pageSize: 10,
			want: &SimplePagination{
				Total:    5,
				Page:     1,
				PageSize: 10,
				LastPage: 1,
			},
		},
		{
			name:     "total equals pageSize",
			total:    10,
			page:     1,
			pageSize: 10,
			want: &SimplePagination{
				Total:    10,
				Page:     1,
				PageSize: 10,
				LastPage: 1,
			},
		},
		{
			name:     "uneven pages",
			total:    95,
			page:     1,
			pageSize: 10,
			want: &SimplePagination{
				Total:    95,
				Page:     1,
				PageSize: 10,
				LastPage: 10,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewSimplePagination(tt.total, tt.page, tt.pageSize)
			if got.Total != tt.want.Total || got.Page != tt.want.Page ||
				got.PageSize != tt.want.PageSize || got.LastPage != tt.want.LastPage {
				t.Errorf("NewSimplePagination() = %+v, want %+v", got, tt.want)
			}
		})
	}
}

func TestGetOffset(t *testing.T) {
	tests := []struct {
		name     string
		page     int
		pageSize int
		want     int64
	}{
		{
			name:     "page 1",
			page:     1,
			pageSize: 10,
			want:     0,
		},
		{
			name:     "page 2",
			page:     2,
			pageSize: 10,
			want:     10,
		},
		{
			name:     "page 5",
			page:     5,
			pageSize: 10,
			want:     40,
		},
		{
			name:     "page 1 with pageSize 20",
			page:     1,
			pageSize: 20,
			want:     0,
		},
		{
			name:     "page 3 with pageSize 20",
			page:     3,
			pageSize: 20,
			want:     40,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetOffset(tt.page, tt.pageSize)
			if got != int(tt.want) {
				t.Errorf("GetOffset() = %d, want %d", got, tt.want)
			}
		})
	}
}

func TestGetLimit(t *testing.T) {
	tests := []struct {
		name     string
		pageSize int
		want     int
	}{
		{
			name:     "pageSize 10",
			pageSize: 10,
			want:     10,
		},
		{
			name:     "pageSize 20",
			pageSize: 20,
			want:     20,
		},
		{
			name:     "pageSize 50",
			pageSize: 50,
			want:     50,
		},
		{
			name:     "pageSize 0",
			pageSize: 0,
			want:     DefaultPageSize,
		},
		{
			name:     "pageSize -1",
			pageSize: -1,
			want:     DefaultPageSize,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetLimit(tt.pageSize)
			if got != tt.want {
				t.Errorf("GetLimit() = %d, want %d", got, tt.want)
			}
		})
	}
}
