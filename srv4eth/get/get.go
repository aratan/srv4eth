package get

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

func Show() int {
	datos := strings.NewReader(`{
		"jsonrpc":"2.0",
		"method":"eth_call",
		"params":[{"to":"0x6bf1dd22cb476e7244ca0b70c985b8ebdc4b8225",
		"data":"0xb703ec59"},
		"latest"],
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
	//fmt.Println("Transferencia: \n", body )
	b,_ := fmt.Printf("Leer Smart Contract: \n%s", cuentas.Result[:])
	return b
}
