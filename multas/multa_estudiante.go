package multas

type MultaEstudiante struct{}

func (m MultaEstudiante) Calcular(diasRetraso int) int {
	return diasRetraso * 5
}
