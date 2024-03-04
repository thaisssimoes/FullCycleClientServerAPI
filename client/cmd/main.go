package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

const (
	_localURL   = "http://localhost:8080"
	_cotacaoURL = "/cotacao"

	_cotacaoTimeout = 300 * time.Millisecond
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

	fmt.Print(string(cotacao))

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

}
