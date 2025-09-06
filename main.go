package main

import (
	"fmt"
	"golang/solid-go/app"
	"golang/solid-go/domain"
	"golang/solid-go/multas"
	"golang/solid-go/noti"
	"time"
)

func main() {
	fmt.Println("🏛️  SISTEMA DE BIBLIOTECA - PRINCIPIOS SOLID")
	fmt.Println("==================================================")

	// Crear libros
	libro1 := domain.NewLibro("1984", "George Orwell", "978-0-452-28423-4")
	libro2 := domain.NewLibro("Cien años de soledad", "Gabriel García Márquez", "978-84-376-0494-7")

	// Crear usuarios
	usuario1 := domain.NewUsuario("Ana García", "U001")
	usuario2 := domain.NewUsuario("Carlos López", "U002")

	// Crear diferentes gestores con distintas configuraciones
	fmt.Println("\n📋 GESTOR PARA USUARIOS REGULARES:")
	gestorRegular := app.NewGestorPrestamos(
		multas.MultaEstandar{},
		noti.NotificadorEmail{},
	)

	fmt.Println("\n📋 GESTOR PARA ESTUDIANTES:")
	gestorEstudiante := app.NewGestorPrestamos(
		multas.MultaEstudiante{},
		noti.NotificadorSMS{},
	)

	// Realizar préstamos
	fmt.Println("\n🔄 REALIZANDO PRÉSTAMOS:")
	prestamo1 := gestorRegular.RealizarPrestamo(libro1, usuario1)
	prestamo2 := gestorEstudiante.RealizarPrestamo(libro2, usuario2)

	// Simular devolución tardía
	fmt.Println("\n📅 SIMULANDO DEVOLUCIÓN TARDÍA:")
	if prestamo1 != nil {
		// forzamos fecha de devolución al pasado
		prestamo1.SetFechaDevolucion(time.Now().Add(-3 * 24 * time.Hour))
		gestorRegular.DevolverLibro(prestamo1)
	}

	// Devolución a tiempo
	fmt.Println("\n📅 DEVOLUCIÓN A TIEMPO:")
	if prestamo2 != nil {
		gestorEstudiante.DevolverLibro(prestamo2)
	}
}
