package web

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func Show() {
	fmt.Println("80")
	http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hola index")
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
}

func handler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hola mundo")
	fmt.Println("Hola mundo")

}
