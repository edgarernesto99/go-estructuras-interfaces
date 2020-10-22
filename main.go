package main

import (
	"bufio"
	"fmt"
	"os"

	"./multimedia"
)

func main() {
	var op string
	scaner := bufio.NewScanner(os.Stdin)
	contenido := multimedia.ContenidoWeb{}

	for true {
		fmt.Println(" ============================ ")
		fmt.Println("|            MENU            |")
		fmt.Println("|============================|")
		fmt.Println("| 1. Capturar imagen         |")
		fmt.Println("| 2. Capturar audio          |")
		fmt.Println("| 3. Capturar video          |")
		fmt.Println("| 4. Mostrar contenido       |")
		fmt.Println("| 0. Salir                   |")
		fmt.Println(" ============================ ")
		fmt.Print("Opción: ")
		scaner.Scan()
		op = scaner.Text()

		if op == "1" {
			capturarImagen(&contenido)
		} else if op == "2" {
			capturarAudio(&contenido)
		} else if op == "3" {
			capturarVideo(&contenido)
		} else if op == "4" {
			contenido.Mostrar()
		} else if op == "0" {
			break
		} else {
			fmt.Println("Opcion no valida")
		}
	}
}

func capturarImagen(contenido *multimedia.ContenidoWeb) {
	scaner := bufio.NewScanner(os.Stdin)
	var titulo string
	var formato string
	var canales string

	fmt.Print("Titulo: ")
	scaner.Scan()
	titulo = scaner.Text()

	fmt.Print("Formato: ")
	scaner.Scan()
	formato = scaner.Text()

	fmt.Print("Canales: ")
	scaner.Scan()
	canales = scaner.Text()
	img := multimedia.Imagen{titulo, formato, canales}
	contenido.Contenido = append(contenido.Contenido, &img)
}

func capturarAudio(contenido *multimedia.ContenidoWeb) {
	scaner := bufio.NewScanner(os.Stdin)
	var titulo string
	var formato string
	var duracion int

	fmt.Print("Titulo: ")
	scaner.Scan()
	titulo = scaner.Text()

	fmt.Print("Formato: ")
	scaner.Scan()
	formato = scaner.Text() 

	fmt.Print("Duracion: ")
	fmt.Scanf("%d",&duracion)
	scaner.Scan()				//Lo puse como si fuera el cin.ignore() de c++
	audio := multimedia.Audio{titulo, formato, duracion}
	contenido.Contenido = append(contenido.Contenido, &audio)
}

func capturarVideo(contenido *multimedia.ContenidoWeb) {
	scaner := bufio.NewScanner(os.Stdin)
	var titulo string
	var formato string
	var frames string

	fmt.Print("Titulo: ")
	scaner.Scan()
	titulo = scaner.Text()

	fmt.Print("Formato: ")
	scaner.Scan()
	formato = scaner.Text()

	fmt.Print("Canales: ")
	scaner.Scan()
	frames = scaner.Text()
	video := multimedia.Video{titulo, formato, frames}
	contenido.Contenido = append(contenido.Contenido, &video)
}