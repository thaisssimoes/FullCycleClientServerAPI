package repository

import (
	"context"
	"database/sql"
	"github.com/gin-gonic/gin"
)

type DB struct {
	db *sql.DB
}

func NewDatabase(db *sql.DB) DB {
	return DB{db: db}
}

func (d DB) Conn(ctx context.Context) {
	d.db.Conn(ctx)
}

func (d DB) CloseConn() {
	d.db.Close()
}

func Cotacao(c *gin.Context) {
	cotacao, _ := c.Get("cotacao")

}
