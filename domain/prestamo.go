package domain

import (
	"fmt"
	"time"
)

type Prestamo struct {
	libro           *Libro
	usuario         *Usuario
	fechaPrestamo   time.Time
	fechaDevolucion time.Time
	devuelto        bool
}

func NewPrestamo(libro *Libro, usuario *Usuario) *Prestamo {
	return &Prestamo{
		libro:           libro,
		usuario:         usuario,
		fechaPrestamo:   time.Now(),
		fechaDevolucion: time.Now().Add(14 * 24 * time.Hour),
		devuelto:        false,
	}
}

func (p *Prestamo) Libro() *Libro              { return p.libro }
func (p *Prestamo) Usuario() *Usuario          { return p.usuario }
func (p *Prestamo) FechaPrestamo() time.Time   { return p.fechaPrestamo }
func (p *Prestamo) FechaDevolucion() time.Time { return p.fechaDevolucion }
func (p *Prestamo) Devuelto() bool             { return p.devuelto }

func (p *Prestamo) SetLibro(l *Libro)              { p.libro = l }
func (p *Prestamo) SetUsuario(u *Usuario)          { p.usuario = u }
func (p *Prestamo) SetFechaPrestamo(t time.Time)   { p.fechaPrestamo = t }
func (p *Prestamo) SetFechaDevolucion(t time.Time) { p.fechaDevolucion = t }
func (p *Prestamo) SetDevuelto(d bool)             { p.devuelto = d }

func (p *Prestamo) String() string {
	return fmt.Sprintf("Prestamo -> Libro: [%s], Usuario: [%s], Fecha Prestamo: %s, Fecha Devolucion: %s, Devuelto: %t",
		p.libro.String(),
		p.usuario.String(),
		p.fechaPrestamo.Format("2006-01-02"),
		p.fechaDevolucion.Format("2006-01-02"),
		p.devuelto,
	)
}
