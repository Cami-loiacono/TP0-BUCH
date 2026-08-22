package ejercicios

const (
	VALOR_INVALIDO = -1
	VALOR_NULO     = 0
)

// Swap intercambia dos valores enteros.
func Swap(x *int, y *int) {
	*x, *y = *y, *x
}

// Maximo devuelve la posición del mayor elemento del arreglo, o -1 si el el arreglo es de largo 0. Si el máximo
// elemento aparece más de una vez, se debe devolver la primera posición en que ocurre.
func Maximo(vector []int) int {
	PosicionMaximoElemento := VALOR_INVALIDO
	if len(vector) != VALOR_NULO {
		for i := VALOR_NULO; i < len(vector); i++ {
			if i == 0 {
				PosicionMaximoElemento = i
			} else if vector[i] > vector[PosicionMaximoElemento] && vector[i] != vector[PosicionMaximoElemento] {
				PosicionMaximoElemento = i
			}
		}
		return PosicionMaximoElemento
	}
	return -1
}

// Comparar compara dos arreglos de longitud especificada.
// Devuelve -1 si el primer arreglo es menor que el segundo; 0 si son iguales; o 1 si el primero es el mayor.
// Un arreglo es menor a otro cuando al compararlos elemento a elemento, el primer elemento en el que difieren
// no existe o es menor.
func Comparar(vector1 []int, vector2 []int) int {

	return 0
}

// Seleccion ordena el arreglo recibido mediante el algoritmo de selección.
func Seleccion(vector []int) {
	for i := len(vector) - 1; i > VALOR_NULO; i-- {
		VectorSiguiente := vector[:i+1]
		PosicionMaximoElemento := Maximo(VectorSiguiente)
		Swap(&vector[i], &vector[PosicionMaximoElemento])
	}
}

// Suma devuelve la suma de los elementos de un arreglo. En caso de no tener elementos, debe devolver 0.
// Esta función debe implementarse de forma RECURSIVA. Se puede usar una función auxiliar (que sea
// la recursiva).
func Suma(vector []int) int {
	if len(vector)!=VALOR_NULO {
		NumeroActual := vector[0]
		SliceSiguiente := vector[1:]
		return NumeroActual + Suma(SliceSiguiente)
	}
	return 0
}

// EsCadenaCapicua devuelve si la cadena es un palíndromo. Es decir, si se lee igual al derecho que al revés.
// Esta función debe implementarse de forma RECURSIVA. Se puede usar una función auxiliar (que sea
// la recursiva).
func EsCadenaCapicua(cadena string) bool {
	return false
}