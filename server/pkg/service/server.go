package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/thaisssimoes/FullCycleClientServerAPI/server/pkg/repository"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	_cotacaoDolarUSDAPITimeout = 300 * time.Millisecond

	_cotacaoUSDBRLURL = "https://economia.awesomeapi.com.br/json/last/USD-BRL"
)

func App() {
	s := gin.Default()
	rotas(s)
	log.Fatalln(s.Run(":8080"))
}

func rotas(s *gin.Engine) {
	s.GET("/cotacao", Cotacao)
}

func Cotacao(c *gin.Context) {
	var resp *http.Response
	var cotacaoDolarReal repository.CotacaoAtual

	newDB := repository.NewDB("./db/fullcycle.db")
	db, err := newDB.Connect()
	if err != nil {
		fmt.Printf("%v", err)
	}

	ctx, cancel := context.WithTimeout(c, _cotacaoDolarUSDAPITimeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, _cotacaoUSDBRLURL, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err,
		})
	}

	resp, err = http.DefaultClient.Do(req)
	if err != nil || resp == nil {
		select {
		case <-ctx.Done():
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "tempo de contexto do servidor excedido",
				"err":     err,
			})
			log.Printf("tempo de contexto do servidor excedido")

		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"err": err,
			})
			log.Printf("err = %v", err)
		}
	}

	err = json.NewDecoder(resp.Body).Decode(&cotacaoDolarReal)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		log.Printf("err = %v", err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	c.Set("cotacao", cotacaoDolarReal)

	err = repository.InsertCotacao(c, db)
	if err != nil {
		log.Printf("erro ao tentar salvar valor no banco de dados. err = %v", err)
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
	}

	c.IndentedJSON(http.StatusOK, string(cotacaoDolarReal.CotacaoDolarReal.Bid))

}
