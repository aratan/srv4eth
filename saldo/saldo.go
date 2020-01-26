package saldo

import (
	"encoding/json"
	"io/ioutil"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type Saldo struct {
	ID      int      `json:"id"`
	Jsonrpc string   `json:"jsonrpc"`
	Result  []string `json:"result"`
}

var saldo Saldo

func Show() string {
	datos := strings.NewReader(`{
		"jsonrpc":"2.0",
		"method":"eth_getBalance",
		"params":["0xae93e90bee174dafa37adce3084bd0b4aea97a71","latest"],
		"id":1
	}`)


	res, err := http.Post("http://127.0.0.1:8545", "application/json", datos)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		panic(err.Error())
	}

	err = json.Unmarshal(body, &saldo)

	b := fmt.Sprintf("%s",body[34:53])
	return b
}
