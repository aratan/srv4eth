package main

import(
	"fmt"
	"github.com/aratan/eth/cuentas"
	"github.com/aratan/eth/saldo"
	"github.com/aratan/eth/envio"
	//"github.com/aratan/eth/conversion"
)

func main() {
	fmt.Println("Cuenta:..........")
	fmt.Println(cuentas.Show())
	fmt.Println("Saldo:...........")
	fmt.Println(saldo.Show())
	fmt.Println("Envio:...........")
	fmt.Println(envio.Show())

	fmt.Println("Saldo:...........")
	fmt.Println(saldo.Show())
	//fmt.Println(conversion.Show())
}