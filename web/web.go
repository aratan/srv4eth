package web

import (
	"github.com/aratan/srv4eth/cuentas"
	//"fmt"
	"io"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	//"html/template"
	//"strconv"
	"time"
)




func Gethandler(w http.ResponseWriter, r *http.Request) {
	var cuenta string
	cuenta = "Mi cuenta: " + cuentas.Show()
	//cuenta := cuentas.Show()
	io.WriteString(w, cuenta)
}


func Show() string {
	// Ruteador
	r := mux.NewRouter().StrictSlash(false)
	// Rutas
	r.HandleFunc("/", Gethandler).Methods("GET")


	// Configuracion servicio web
	server := &http.Server{
		Addr:           ":80",
		Handler:		r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,  // 1 mega
	}

	//activa servidor web
	log.Println("Escuchando 80")
	log.Fatal(server.ListenAndServe())
	return ""
}
