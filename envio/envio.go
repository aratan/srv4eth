package envio

import (
	"encoding/json"
	"fmt"
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
		"method":"eth_sendTransaction",
		"params":[{"from":"0xAE93e90bEE174DaFA37AdCe3084BD0b4Aea97a71", 
		"to":"0x802c74CCE0F501106F634eCED227e443675A61Da", 
		"gas":"0x15f90",
		"gasPrice":"0x430e23400",
		"value":"0x9b6e64a8ec60000"
		}],
		"data": "0xd46e8dd67c5d32be8d46e8dd67c5d32be8058bb8eb970870f072445675058bb8eb970870f072445675",
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
	b := fmt.Sprintf("Transferencia Nº: \n%s",body[34:100])
	return b
}
