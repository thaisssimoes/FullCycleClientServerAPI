package client

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Cotacao(c *gin.Context) {
	var cotacaoDolarReal CotacaoAtual

	r, err := http.Get(_cotacaoUSDBRLURL)
	if err != nil {
		panic(err)
	}

	err = json.NewDecoder(r.Body).Decode(&cotacaoDolarReal)
	if err != nil {
		panic(err)
	}

	c.IndentedJSON(http.StatusOK, cotacaoDolarReal)
}
