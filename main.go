package goarea

import "math"

// Pi é uma proporcao numerica definida pela relação entre
// o perimetro de uma circuferencia e o seu diametro
const Pi = 3.1416

// Circ é responsavel por calcular a area da circuferencia
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

// Rect eh responsavel por calcular a area de uma retangulo
func Rect(base, altura float64) float64 {
	return base * altura
}

// Não eh visivel!
func _TrianguloEq(base, altura float64) float64 {
	return (base * altura) / 2
}
