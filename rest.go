package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

//"Peticones rest para el usuario..."
func Peticiones() {
	http.HandleFunc("/selectAllUser", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			w.Write([]byte("Pedido get"))
		}
	})
	http.HandleFunc("/selectUser", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			/* localhost:8000?nombre="Edwin"&apellido="Ibarra"
			Muestra
			ma[nombre["Edwin"]]
			fmt.Println(r.URL.Query()) */
			fmt.Println(r.URL.RawQuery)
			fmt.Println(r.URL.Query().Get("id"))
			for m, v := range r.URL.Query() {
				fmt.Printf("%s resultado %s\n", m, v)
			}
		}
	})
	http.HandleFunc("/agregarUser", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			w.Write([]byte("Agregar user"))
			fmt.Println(r.Body)
			datosDelBody, err := ioutil.ReadAll(r.Body)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(string(datosDelBody))
				fmt.Fprintf(w, " el dato es %s", datosDelBody)
			}
		}
	})
	http.HandleFunc("/editUser", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "PUT" {
			w.Write([]byte("Editar usuario"))
		}
	})
	http.HandleFunc("/eliminarUser", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "DELETE" {
			w.Write([]byte("Eliminar usuario"))
		}
	})
}
