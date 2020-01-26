package web

import (
	"github.com/aratan/eth/cuentas"
	"fmt"
	"io"
	"log"
	"net/http"
	//"strconv"
	"time"
)

func Show() string {
	fmt.Println("80")
	http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hola")

	})

	http.HandleFunc("/", handler)
	// Se cambia http.ListenAndServe(":80",nil) por lo siguiente:
	// Configuracion servicio web
	server := &http.Server{
		Addr:           ":80",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
		// 1 mega
	}
	//activa servidor web
	log.Println("Escuchando 80")
	log.Fatal(server.ListenAndServe())
	return "aaa"
}

func handler(w http.ResponseWriter, r *http.Request) {
	var cuenta string
	cuenta = cuentas.Show() + "Mi cuenta"
	//cuenta := cuentas.Show()
	
	io.WriteString(w, cuenta)
	
}
