package main

type Person struct {
	nombre           string
	edad             int
	cedula           int
	cuidadRecidencia string
}

func (p *Person) getNombre() string {
	return p.nombre
}
