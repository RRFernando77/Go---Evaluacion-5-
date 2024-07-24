package main

// Cada programa en Go está formado por paquetes.
// 'package main' indica que este es el paquete principal, el que se ejecuta.

import (
	"fmt"  // fmt (format) nos permite imprimir y formatear texto.
	"math" // math nos da acceso a funciones matemáticas como potencias, raíces, etc.
)

// Función Objetivo:  Nuestra Parábola

// Define una función llamada 'objectiveFunction' que recibe un número
// 'x' de tipo 'float64' (números con decimales) y devuelve otro 'float64'.
func objectiveFunction(x float64) float64 {
	return math.Pow(x-3, 2) // Calcula (x-3) elevado al cuadrado. ¡Esta es la ecuación de nuestra parábola!
}

// Derivada:  La Pendiente de la Parábola

// Define la función 'derivative' que calcula la derivada de la función objetivo.
// La derivada nos indica la pendiente de la parábola en un punto 'x'.
func derivative(x float64) float64 {
	return 2 * (x - 3) // La derivada de (x-3)^2 es 2(x-3)
}

// Descenso de Gradiente:  Encontrando el Punto Más Bajo

// Función principal del algoritmo, busca el mínimo de la función.
func gradientDescent(x0, alpha float64, maxIter int) (float64, float64) {
	// x0: Punto inicial de búsqueda
	// alpha: Tamaño del paso en cada iteración (controla la velocidad)
	// maxIter: Número máximo de iteraciones (para evitar bucles infinitos)

	x := x0 // Empezamos en el punto x0

	// Bucle principal del algoritmo. Se ejecutará 'maxIter' veces como máximo.
	for i := 0; i < maxIter; i++ {
		gradient := derivative(x) // 1. Calcula la pendiente (gradiente) en el punto 'x'.

		// 2. Actualiza 'x' moviéndonos en la dirección opuesta a la pendiente.
		//    Multiplicamos por 'alpha' para controlar la velocidad del descenso.
		x = x - alpha*gradient

		// Imprimimos información sobre la iteración actual (opcional).
		fmt.Printf("Iteración %d: x = %.4f, f(x) = %.4f\n", i+1, x, objectiveFunction(x))
	}

	// Devuelve el 'x' donde encontramos el mínimo (aproximado) y el valor de la función en ese punto.
	return x, objectiveFunction(x)
}

// Función Principal:  Donde Todo Comienza

func main() {
	// Parámetros iniciales para el algoritmo
	x0 := 0.0      // Punto inicial
	alpha := 0.1   // Tamaño del paso
	maxIter := 100 // Máximo de iteraciones

	// Llamamos al algoritmo de descenso de gradiente con nuestros parámetros.
	// Guarda los resultados en 'xMin' (el 'x' del mínimo) y 'fMin' (el valor mínimo).
	xMin, fMin := gradientDescent(x0, alpha, maxIter)

	// Imprime los resultados finales.
	fmt.Println("\nResultados:")
	fmt.Printf("Valor mínimo de x: %.4f\n", xMin)
	fmt.Printf("Valor mínimo de la función: %.4f\n", fMin)
}
