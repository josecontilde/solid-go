package capacidades

import "fmt"

type PrestamoRenovable struct {
	id string
}

func (p PrestamoRenovable) Renovar(prestamoID string) bool {
	fmt.Printf("🔄 El préstamo %s fue renovado\n", prestamoID)
	return true
}
