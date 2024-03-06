package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type Config struct {
	File string
}

func NewDB(file string) *Config {
	return &Config{
		File: file,
	}
}

func (c Config) Connect() (*sqlx.DB, error) {
	db, err := sqlx.Open("sqlite3", c.File)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func InsertCotacao(ctx context.Context, db *sqlx.DB) error {
	cotacao := ctx.Value("cotacao")

	c, ok := cotacao.(CotacaoAtual)
	if !ok {
	}

	stmt, err := db.Prepare("INSERT INTO cotacao_dolar (code, code_in, name, high, low, var_bid, pct_change, bid, " +
		"ask, timestamp, create_date) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)")
	if err != nil {
		log.Fatalln(err)
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, c.CotacaoDolarReal.Code, c.CotacaoDolarReal.CodeIn, c.CotacaoDolarReal.Name, c.CotacaoDolarReal.High,
		c.CotacaoDolarReal.Low, c.CotacaoDolarReal.VarBid, c.CotacaoDolarReal.PctChange, c.CotacaoDolarReal.Bid,
		c.CotacaoDolarReal.Ask, c.CotacaoDolarReal.Timestamp, c.CotacaoDolarReal.CreateDate)
	if err != nil {
		log.Fatalln(err)
	}
	return nil
}
