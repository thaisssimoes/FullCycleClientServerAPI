package service

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/thaisssimoes/FullCycleClientServerAPI/server/pkg/repository"
	"log"
	"net/http"
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

	r, err := http.Get(_cotacaoUSDBRLURL)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err,
		})
	}

	defer r.Body.Close()

	err = json.NewDecoder(r.Body).Decode(&cotacaoDolarReal)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err,
		})
	}

	c.Set("cotacao", cotacaoDolarReal)
	repository.Cotacao(c)

	c.IndentedJSON(http.StatusOK, cotacaoDolarReal)

}
