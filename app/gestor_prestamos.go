package app

import (
	"fmt"
	"golang/solid-go/domain"
	"golang/solid-go/multas"
	"golang/solid-go/noti"
	"time"
)

type GestorPrestamos struct {
	calculadora multas.CalculadoraMulta
	notificador noti.Notificador
	prestamos   []*domain.Prestamo
}

func NewGestorPrestamos(calc multas.CalculadoraMulta, not noti.Notificador) *GestorPrestamos {
	return &GestorPrestamos{
		calculadora: calc,
		notificador: not,
		prestamos:   []*domain.Prestamo{},
	}
}

func (g *GestorPrestamos) RealizarPrestamo(libro *domain.Libro, usuario *domain.Usuario) *domain.Prestamo {
	if !libro.Disponible() {
		fmt.Printf("❌ El libro '%s' no está disponible\n", libro.Titulo())
		return nil
	}

	libro.SetDisponible(false)
	prestamo := domain.NewPrestamo(libro, usuario)
	g.prestamos = append(g.prestamos, prestamo)

	mensaje := fmt.Sprintf("Has prestado '%s'. Fecha de devolución: %s",
		libro.Titulo(), prestamo.FechaDevolucion().Format("2006-01-02"))
	g.notificador.Enviar(mensaje, usuario.Nombre())

	fmt.Printf("✅ Préstamo realizado: '%s' para %s\n", libro.Titulo(), usuario.Nombre())
	return prestamo
}

func (g *GestorPrestamos) DevolverLibro(prestamo *domain.Prestamo) {
	if prestamo.Devuelto() {
		fmt.Println("❌ Este libro ya fue devuelto")
		return
	}

	fechaActual := time.Now()
	if fechaActual.After(prestamo.FechaDevolucion()) {
		diasRetraso := int(fechaActual.Sub(prestamo.FechaDevolucion()).Hours() / 24)
		multa := g.calculadora.Calcular(diasRetraso)
		fmt.Printf("⚠️  Retraso de %d días. Multa: $%d\n", diasRetraso, multa)
	} else {
		fmt.Println("✅ Libro devuelto a tiempo")
	}

	prestamo.SetDevuelto(true)
	prestamo.Libro().SetDisponible(true)
	fmt.Printf("📚 '%s' devuelto por %s\n", prestamo.Libro().Titulo(), prestamo.Usuario().Nombre())
}
