package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"tp0/ejercicios"
)

const (
	rutaArchivo1 = "archivo1.in"
	rutaArchivo2 = "archivo2.in"
)

func LeerArchivo(ruta string) ([]int, error) {
	archivo, err := os.Open(ruta)
	if err != nil {
		return nil, err
	}
	defer archivo.Close()

	var arrayArchivo []int
	s := bufio.NewScanner(archivo)
	for s.Scan() {
		elemento, err := strconv.Atoi(s.Text())
		if err != nil {
			return nil, err
		}
		arrayArchivo = append(arrayArchivo, elemento)
	}
	err = s.Err()
	if err != nil {
		return nil, err
	}

	return arrayArchivo, nil
}

func ImprimirArreglo(arreglo []int) {
	for i := 0; i < len(arreglo); i++ {
		fmt.Println(arreglo[i])
	}
}

func LeerArchivos(arreglo1 *[]int, arreglo2 *[]int) bool {
	exito := true
	var err error
	*arreglo1, err = LeerArchivo(rutaArchivo1)

	if err != nil {
		fmt.Println("Error al leer", rutaArchivo1, "por: ", err)
		exito = false
	}

	*arreglo2, err = LeerArchivo(rutaArchivo2)
	if err != nil {
		fmt.Println("Error al leer", rutaArchivo2, " por: ", err)
		exito = false
	}
	return exito
}

func main() {
	var arregloArchivo1, arregloArchivo2 []int
	if !LeerArchivos(&arregloArchivo1, &arregloArchivo2) {
		return
	}
	arregloMayor := ejercicios.Comparar(arregloArchivo1, arregloArchivo2)
	switch arregloMayor {
	case 1:
		fmt.Println("El arreglo mayor es el del archivo 1")
		ejercicios.Seleccion(arregloArchivo1)
		ImprimirArreglo(arregloArchivo1)
	case -1:
		fmt.Println("El arreglo mayor es el del archivo 2")
		ejercicios.Seleccion(arregloArchivo2)
		ImprimirArreglo(arregloArchivo2)
	case 0:
		fmt.Println("Los arreglos son iguales")
		ejercicios.Seleccion(arregloArchivo1)
		ImprimirArreglo(arregloArchivo1)
	}
}
