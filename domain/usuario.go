package domain

import "fmt"

type Usuario struct {
	nombre    string
	idUsuario string
}

func NewUsuario(nombre, id string) *Usuario {
	return &Usuario{nombre: nombre, idUsuario: id}
}

func (u *Usuario) Nombre() string    { return u.nombre }
func (u *Usuario) IDUsuario() string { return u.idUsuario }

func (u *Usuario) SetNombre(n string)     { u.nombre = n }
func (u *Usuario) SetIDUsuario(id string) { u.idUsuario = id }

func (u *Usuario) String() string {
	return fmt.Sprintf("Usuario: %s, ID: %s", u.nombre, u.idUsuario)
}
