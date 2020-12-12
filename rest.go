package main

import (
	"net/http"
)

//"Peticones rest para el usuario..."
func Peticiones() {
	http.HandleFunc("/selectAllUser", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			w.Write([]byte("Pedido get"))
		}
	})
	http.HandleFunc("/agregarUser", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			w.Write([]byte("Agregar user"))
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
