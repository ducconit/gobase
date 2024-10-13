package db

import (
	"github.com/ducconit/gobase/config"
	"gorm.io/gorm"
	"log"
)

type DB struct {
	*gorm.DB
}

func New(cfg *Config, opts ...gorm.Option) (*DB, error) {
	db, err := gorm.Open(cfg, opts...)
	if err != nil {
		return nil, err
	}

	return &DB{db}, nil
}

func NewFromConfig(cfg config.Store) (*DB, error) {
	c := new(Config)

	if err := c.fromConfig(cfg); err != nil {
		return nil, err
	}
	return New(c)
}

func (d *DB) Setup() error {
	return nil
}

func (d *DB) Close() error {
	db, err := d.DB.DB()
	if err != nil {
		return err
	}

	log.Println("Closing db connection")

	return db.Close()
}
