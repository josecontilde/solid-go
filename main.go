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
	title("🏛️  SISTEMA DE BIBLIOTECA - PRINCIPIOS SOLID")

	c := setupContainer()

	seed := seedData()

	p1, p2 := escenarioPrestamos(c, seed)

	escenarioDevoluciones(c, p1, p2)
}

type container struct {
	gestorRegular    *app.GestorPrestamos
	gestorEstudiante *app.GestorPrestamos
}

func setupContainer() *container {
	gestorRegular := app.NewGestorPrestamos(multas.MultaEstandar{}, noti.NotificadorEmail{})
	gestorEstudiante := app.NewGestorPrestamos(multas.MultaEstudiante{}, noti.NotificadorSMS{})
	return &container{gestorRegular, gestorEstudiante}
}

type seedSet struct {
	libro1   *domain.Libro
	libro2   *domain.Libro
	usuario1 *domain.Usuario
	usuario2 *domain.Usuario
}

func seedData() *seedSet {
	section("Datos de prueba")
	libro1 := domain.NewLibro("1984", "George Orwell", "978-0-452-28423-4")
	libro2 := domain.NewLibro("Cien años de soledad", "G. García Márquez", "978-84-376-0494-7")
	usuario1 := domain.NewUsuario("Ana García", "U001")
	usuario2 := domain.NewUsuario("Carlos López", "U002")
	fmt.Println("• Libros y usuarios creados")
	return &seedSet{libro1, libro2, usuario1, usuario2}
}

// -------------------- Escenarios de negocio --------------------

func escenarioPrestamos(c *container, s *seedSet) (p1, p2 *domain.Prestamo) {
	section("Escenario 1 · Realizando préstamos")
	fmt.Println("📋 Gestor REGULAR")
	p1 = c.gestorRegular.RealizarPrestamo(s.libro1, s.usuario1)

	fmt.Println("\n📋 Gestor ESTUDIANTE")
	p2 = c.gestorEstudiante.RealizarPrestamo(s.libro2, s.usuario2)
	return
}

func escenarioDevoluciones(c *container, p1, p2 *domain.Prestamo) {
	section("Escenario 2 · Devoluciones")

	fmt.Println("📅 Devolución tardía (forzada 3 días)")
	if p1 != nil {
		p1.SetFechaDevolucion(time.Now().Add(-3 * 24 * time.Hour)) // simulamos atraso
		c.gestorRegular.DevolverLibro(p1)
	}

	fmt.Println("\n📅 Devolución a tiempo")
	if p2 != nil {
		c.gestorEstudiante.DevolverLibro(p2)
	}
}

// -------------------- helpers estéticos --------------------

func title(t string) {
	fmt.Println(t)
	fmt.Println("==================================================")
}
func section(t string) { fmt.Printf("\n%s\n", t) }
