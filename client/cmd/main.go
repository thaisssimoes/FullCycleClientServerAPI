package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	_localURL   = "http://localhost:8080"
	_cotacaoURL = "/cotacao"

	_cotacaoTimeout = 300 * time.Millisecond
)

func main() {
	var cotacao []byte

	ctx, cancel := context.WithTimeout(context.Background(), _cotacaoTimeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, _localURL+_cotacaoURL, nil)
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		select {
		case <-ctx.Done():
			log.Fatalf("tempo de contexto do client excedido. err = %v", err)
		default:
			log.Fatalf("err= %v", err)
		}
	}

	cotacao, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	cotacaoFloat, err := strconv.ParseFloat(strings.Trim(string(cotacao), "\""), 64)
	if err != nil {
		log.Fatalln("\no valor não pode ser convertido para float64. err = v", err)
	}

	escreverCotacaoArquivo("./files/cotacao.txt", cotacaoFloat)

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

}

func escreverCotacaoArquivo(path string, cotacao float64) {
	f, err := os.Create(path)
	if err != nil {
		log.Fatalf("erro ao criar o arquivo. err = %v", err)
	}

	defer f.Close()

	mensagem := []byte(fmt.Sprintf("Dólar: %f", cotacao))

	_, err = f.Write(mensagem)
	if err != nil {
		log.Printf("erro ao escrever no arquivo. err = %v", err)
	}

	log.Println("arquivo preenchido com sucesso")
}
