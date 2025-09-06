package multas

type MultaEstandar struct{}

func (m MultaEstandar) Calcular(diasRetraso int) int {
	return diasRetraso * 10
}
