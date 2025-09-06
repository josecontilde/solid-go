package capacidades

import "fmt"

type LibroPrestable struct {
	titulo string
}

func (l LibroPrestable) Reservar(usuario string) bool {
	fmt.Printf("📖 El libro '%s' fue reservado por %s\n", l.titulo, usuario)
	return true
}

func (l LibroPrestable) Prestar(usuario string) bool {
	fmt.Printf("📚 El libro '%s' fue prestado a %s\n", l.titulo, usuario)
	return true
}
