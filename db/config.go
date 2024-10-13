package db

import (
	"fmt"
	"github.com/ducconit/gobase/config"
	"github.com/ducconit/gobase/db/errors"
	mysql2 "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Config struct {
	gorm.Dialector

	Driver string
	cfg    config.Store
}

func (c *Config) fromConfig(cfg config.Store) error {
	c.Driver = cfg.GetString("db.driver")
	c.cfg = cfg

	switch c.Driver {
	case "postgres":
		c.Dialector = c.newPostgres()
		return nil
	case "mysql":
		c.Dialector = c.newMySQL()
		return nil
	case "sqlite":
		c.Dialector = c.newSQLite()
		return nil
	}
	return errors.DriverNotSupported
}

func (c *Config) newPostgres() gorm.Dialector {
	dsn := c.cfg.GetString("db.dsn")
	if dsn != "" {
		return postgres.Open(dsn)
	}

	return postgres.Open(fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		c.getHost(), c.getPort(), c.getUser(), c.getPass(), c.dbName()))
}

func (c *Config) newMySQL() gorm.Dialector {
	dsn := c.getBaseOrInfoDriver("dsn")
	if dsn != "" {
		return mysql.Open(dsn)
	}
	return mysql.New(mysql.Config{
		DSNConfig: &mysql2.Config{
			User:      c.getUser(),
			Passwd:    c.getPass(),
			Addr:      c.getAddr(),
			DBName:    c.dbName(),
			Params:    nil,
			Collation: c.getCollation(),
		},
	})
}

func (c *Config) newSQLite() gorm.Dialector {
	dsn := c.getBaseOrInfoDriver("dsn")
	if dsn != "" {
		return sqlite.Open(dsn)
	}
	return sqlite.Open(c.getBaseOrInfoDriver("database"))
}

func (c *Config) getBaseOrInfoDriver(key string) string {
	base := c.cfg.GetString("db." + key)
	if base != "" {
		return base
	}
	return c.cfg.GetString(fmt.Sprintf("db.%s.%s", key, c.Driver))
}

func (c *Config) getUser() string {
	return c.getBaseOrInfoDriver("user")
}

func (c *Config) getPass() string {
	return c.getBaseOrInfoDriver("pass")
}

func (c *Config) getAddr() string {
	addr := c.getBaseOrInfoDriver("addr")
	if addr != "" {
		return addr
	}

	host := c.getBaseOrInfoDriver("host")
	port := c.getBaseOrInfoDriver("port")
	return host + ":" + port
}

func (c *Config) getHost() string {
	return c.getBaseOrInfoDriver("host")
}

func (c *Config) getPort() string {
	return c.getBaseOrInfoDriver("port")
}

func (c *Config) dbName() string {
	return c.getBaseOrInfoDriver("database")
}

func (c *Config) getCollation() string {
	return c.getBaseOrInfoDriver("collation")
}
