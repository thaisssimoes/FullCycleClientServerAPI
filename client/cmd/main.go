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

	_cotacaoTimeout = 3000 * time.Millisecond
)

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), _cotacaoTimeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, _localURL+_cotacaoURL, nil)
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := http.DefaultClient.Do(req)

	cotacao, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	cotacaoFloat, err := strconv.ParseFloat(strings.Trim(string(cotacao), "\""), 64)
	if err != nil {
		log.Fatalln("o valor não pode ser convertido para float64. err = v", err)
	}

	escreverArquivo("./files/cotacao.txt", cotacaoFloat)

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

}

func escreverArquivo(path string, cotacao float64) {
	f, err := os.Create(path)
	if err != nil {
		log.Fatalf("erro ao criar o arquivo. err = %v", err)
	}

	defer f.Close()

	mensagem := fmt.Sprintf("Cotação atual é: %f", cotacao)
	mensagemByte := []byte(mensagem)

	_, err = f.Write(mensagemByte)
	if err != nil {
		log.Fatalf("erro ao escrever no arquivo. err = %v", err)
	}

	log.Println("arquivo preenchido com sucesso")
}
