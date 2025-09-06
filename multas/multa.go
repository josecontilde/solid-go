package multas

type CalculadoraMulta interface {
	Calcular(diasRetraso int) int
}
