package cuentas

import (
	"encoding/json"
	//"fmt"//
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Cuentas struct {
	ID      int      `json:"id"`
	Jsonrpc string   `json:"jsonrpc"`
	Result  []string `json:"result"`
}

var cuentas Cuentas

func Show() string {
	datos := strings.NewReader(`{
		"jsonrpc":"2.0",
		"method":"eth_accounts",
		"params":[],
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

	err = json.Unmarshal(body, &cuentas)
	//fmt.Println(cuentas.Result)//
	return cuentas.Result[0]
}
