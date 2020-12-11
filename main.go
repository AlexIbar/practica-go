package main

import (
	"fmt"
	"net/http"
)

func main() {

	var c = Person{
		nombre:           "Juan",
		edad:             22,
		cedula:           1036620,
		cuidadRecidencia: "Bogota",
	}
	fmt.Println(c)

	fileServe := http.FileServer(http.Dir("public"))
	http.Handle("/", http.StripPrefix("/", fileServe))

	http.ListenAndServe("8000", nil)
}

//Para poder ejecutar el go run y que en ejecución importe los archivos del mismo paquete debemos hacerlo así: go run *.go
