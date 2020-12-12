package main

type person struct {
	nombre           string
	edad             int
	cedula           int
	cuidadRecidencia string
}

func (p *person) getNombre() string {
	return p.nombre
}
