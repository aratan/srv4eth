package web

import (
	"github.com/aratan/srv4eth/cuentas"
	"github.com/aratan/srv4eth/saldo"
	//"fmt"
	//"io/ioutil"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"html/template"
	//"os"
	//"strconv"
	"time"
)

type Comandos struct {
    Cuentas string
    Saldo   string
}

func Gethandler(w http.ResponseWriter, r *http.Request) {

	p := Comandos{Cuentas: 	cuentas.Show(), Saldo:	saldo.Show() }
	t, _ := template.ParseFiles("index.html")
	t.Execute(w, p)	
}


func Show() string {
	// Ruteador
	r := mux.NewRouter().StrictSlash(false)
	// Rutas
	r.HandleFunc("/", Gethandler).Methods("GET")


	// Configuracion servicio web
	server := &http.Server{
		Addr:           "127.0.0.1:80",
		Handler:		r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,  // 1 mega
	}

	//activa servidor web
	log.Println("Create by Victor Arbiol Martinez\nlistening... 127.0.0.1:80")
	log.Fatal(server.ListenAndServe())
	return ""
}
