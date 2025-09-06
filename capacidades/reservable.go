package capacidades

import "fmt"

type Reservable interface {
	Reservar(usuario string) bool
}

type LibroReservable struct {
	titulo string
}

func (l LibroReservable) Reservar(usuario string) bool {
	fmt.Printf("📖 El libro '%s' fue reservado por %s\n", l.titulo, usuario)
	return true
}
