package mathutil

func Add(x, y int) int {
	return x + y
}

func Resta(x, y int) int {
	return x - y
}

const GravedadMarte = 3.71
const GravedadTierra = 9.81

func CalcularPeso(masa float64, gravedad float64) float64 {
	return masa * gravedad
}