package multimedia

import "fmt"

type ContenidoWeb struct {
	Contenido []Multimedia
}

func (c *ContenidoWeb) Mostrar() {
	fmt.Println()
	for _, v := range c.Contenido {
		v.Mostrar()
	}
}

type Multimedia interface {
	Mostrar()
}

type Imagen struct {
	Titulo string
	Formato string
	Canales string
}

func (i *Imagen) Mostrar() {
	fmt.Println("IMAGEN")
	fmt.Println("Titulo: ", i.Titulo)
	fmt.Println("Formato: ", i.Formato)
	fmt.Println("Canales: ", i.Canales)
	fmt.Println("--------------------------------")
}

type Audio struct {
	Titulo string
	Formato string
	Duracion int
}

func (a *Audio) Mostrar() {
	fmt.Println("AUDIO")
	fmt.Println("Titulo: ", a.Titulo)
	fmt.Println("Formato: ", a.Formato)
	fmt.Println("Duracion (seg): ", a.Duracion)
	fmt.Println("--------------------------------")
}

type Video struct {
	Titulo string
	Formato string
	Frames string
}

func (v *Video) Mostrar() {
	fmt.Println("VIDEO")
	fmt.Println("Titulo: ", v.Titulo)
	fmt.Println("Formato: ", v.Formato)
	fmt.Println("Frames: ", v.Frames)
	fmt.Println("--------------------------------")
}