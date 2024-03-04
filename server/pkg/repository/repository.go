package repository

import (
	"context"
	"database/sql"
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

func Cotacao(ctx context.Context) {
	cotacao := ctx.Value("cotacao")
	cotacao = cotacao

}
