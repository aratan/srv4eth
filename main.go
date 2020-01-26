package main

import(
	"fmt"
	"github.com/aratan/srv4eth/cuentas"
	"github.com/aratan/srv4eth/saldo"
	"github.com/aratan/srv4eth/envio"
	//"github.com/aratan/srv4eth/conversion"
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