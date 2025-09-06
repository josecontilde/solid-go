package domain

import "fmt"

type Libro struct {
	titulo     string
	autor      string
	isbn       string
	disponible bool
}

func NewLibro(titulo, autor, isbn string) *Libro {
	return &Libro{
		titulo:     titulo,
		autor:      autor,
		isbn:       isbn,
		disponible: true,
	}
}

func (l *Libro) Titulo() string   { return l.titulo }
func (l *Libro) Autor() string    { return l.autor }
func (l *Libro) ISBN() string     { return l.isbn }
func (l *Libro) Disponible() bool { return l.disponible }

func (l *Libro) SetTitulo(t string)   { l.titulo = t }
func (l *Libro) SetAutor(a string)    { l.autor = a }
func (l *Libro) SetISBN(i string)     { l.isbn = i }
func (l *Libro) SetDisponible(d bool) { l.disponible = d }

func (l *Libro) String() string {
	return fmt.Sprintf("Libro: %s, Autor: %s, ISBN: %s, Disponible: %t",
		l.titulo, l.autor, l.isbn, l.disponible)
}
