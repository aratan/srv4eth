package main

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

func main() {
	datos := strings.NewReader(`{
		"jsonrpc":"2.0",
		"method":"eth_sendTransaction",
		"params":[{"from":"0x8f9ca1de6c62e330b568d93dcd626ef099df8817",
		"to":"0xeffc39e83cb734b666f48ef9990f4857c8ced72e",
		"data":"0xdb0702610000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000000b486f6c6120766963746f72000000000000000000000000000000000000000000",
		"gas":"0x925e",
		"value":"0x0"}],"id":4208
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
	//fmt.Println("Transferencia: \n", cuentas.Result[:])
	fmt.Printf("getSaludo(): \n%s", body)
}
