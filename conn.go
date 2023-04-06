package oracle

import (
	"database/sql"
	"time"

	_ "github.com/mattn/go-oci8"
)

type Config struct {
	DataSourceName  string
	MaxOpenConn     int
	MaxIdleConn     int
	ConnMaxLifeTime int
	ConnMaxIdleTime int
}

type Client struct {
	cfg *Config
}

func NewClient(cfg *Config) *Client {
	c := new(Client)
	c.cfg = cfg
	return c
}

func (c *Client) Connect() (*sql.DB, error) {
	db, err := sql.Open("oci8", c.cfg.DataSourceName)
	if err != nil {
		return nil, err
	}
	db.SetConnMaxLifetime(time.Duration(c.cfg.ConnMaxLifeTime) * time.Millisecond)
	db.SetConnMaxIdleTime(time.Duration(c.cfg.ConnMaxIdleTime) * time.Millisecond)
	db.SetMaxOpenConns(c.cfg.MaxOpenConn)
	db.SetMaxIdleConns(c.cfg.MaxIdleConn)
	return db, err
}
