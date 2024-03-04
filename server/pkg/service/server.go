package service

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	_cotacaoDolarUSBTimeout = 300 * time.Millisecond
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
	var cotacaoDolarReal CotacaoAtual

	ctx, cancel := context.WithTimeout(c, _cotacaoDolarUSBTimeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, _cotacaoUSDBRLURL, nil)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err,
		})
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err,
		})
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	err = json.NewDecoder(resp.Body).Decode(&cotacaoDolarReal)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err,
		})
	}

	// repository.Cotacao(c)

	c.IndentedJSON(http.StatusOK, cotacaoDolarReal)

}
