package server

import (
	"github.com/gin-gonic/gin"
	"github.com/thaisssimoes/FullCycleClientServerAPI/pkg/client"
	"log"
)

func App() {
	s := gin.Default()
	rotas(s)
	log.Fatalln(s.Run(":8080"))
}

func rotas(s *gin.Engine) {
	s.GET("/cotacao", client.Cotacao)
}
