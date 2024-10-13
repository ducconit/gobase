package model

import (
	"github.com/oklog/ulid/v2"
	"gorm.io/gorm"
)

type Model struct {
	*gorm.Model

	ID ulid.ULID `json:"id"`
}
