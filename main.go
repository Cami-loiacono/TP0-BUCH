package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"tp0/ejercicios"
)

const (
	Ruta_Archivo1 = "archivo1.in"
	Ruta_Archivo2 = "archivo2.in"
)

func LeerArchivo(ruta string) []int {
	var arrayArchivo []int
	archivo, err := os.Open(ruta)
	if err != nil {
		fmt.Println("Error al abrir el archivo:", err)
		return nil
	}
	defer archivo.Close()
	s := bufio.NewScanner(archivo)
	for s.Scan() {
		elemento, err := strconv.Atoi(s.Text())
		if err != nil {
			fmt.Println("Error al convertir el elemento:", err)
			return nil
		}
		arrayArchivo = append(arrayArchivo, elemento)
	}
	return arrayArchivo
}

func LeerArchivos(arreglo1 *[]int, arreglo2 *[]int) {
	*arreglo1 = LeerArchivo(Ruta_Archivo1)
	*arreglo2 = LeerArchivo(Ruta_Archivo2)
}

func ImprimirArreglo(arreglo []int) {
	for i := 0; i < len(arreglo); i++ {
		fmt.Println(arreglo[i])
	}
}

func ImprimirOrdenarArreglo(arreglo []int) {
	ejercicios.Seleccion(arreglo)
	ImprimirArreglo(arreglo)
}

func main() {
	var arregloArchivo1, arregloArchivo2 []int
	LeerArchivos(&arregloArchivo1, &arregloArchivo2)
	if arregloArchivo1 == nil || arregloArchivo2 == nil {
		fmt.Println("Error al leer alguno de los dos archivitos")
		return
	}
	arregloMayor := ejercicios.Comparar(arregloArchivo1, arregloArchivo2)
	switch arregloMayor {
	case ejercicios.PRIMER_MAYOR:
		ImprimirOrdenarArreglo(arregloArchivo1)
	case ejercicios.PRIMER_MENOR:
		ImprimirOrdenarArreglo(arregloArchivo2)
	case ejercicios.ARREGLOS_IGUALES:
		ImprimirOrdenarArreglo(arregloArchivo1)
	}
}
