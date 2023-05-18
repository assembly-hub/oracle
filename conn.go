package oracle

import (
	"database/sql"
	"time"

	"github.com/assembly-hub/db"
	"github.com/assembly-hub/impl-db-sql"
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

func (c *Client) Connect() (db.Executor, error) {
	conn, err := sql.Open("oci8", c.cfg.DataSourceName)
	if err != nil {
		return nil, err
	}
	conn.SetConnMaxLifetime(time.Duration(c.cfg.ConnMaxLifeTime) * time.Millisecond)
	conn.SetConnMaxIdleTime(time.Duration(c.cfg.ConnMaxIdleTime) * time.Millisecond)
	conn.SetMaxOpenConns(c.cfg.MaxOpenConn)
	conn.SetMaxIdleConns(c.cfg.MaxIdleConn)
	return impl.NewDB(conn), err
}
