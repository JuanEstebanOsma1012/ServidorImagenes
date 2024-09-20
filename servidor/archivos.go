package main

import (
	"encoding/base64"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
)

type Imagen struct {
	Nombre    string
	Contenido string
	Indice    int
	Extension string
}

type ImagenContenedor struct {
	Imagenes []Imagen
}

var extensionesPermitidas = []string{".jpg", ".jpeg", ".png"}

const rutaImagenes = "../estaticos/imagenes"

func obtenerImagenesPorTema(tema string, cantidad int) ImagenContenedor {

	archivos, err := os.ReadDir(fmt.Sprintf("%s/%s", rutaImagenes, tema))
	if err != nil {
		panic(err)
	}

	imagenes := filter(archivos, func(archivo os.DirEntry) bool {
		return esImagen(archivo)
	})

	if len(imagenes) < cantidad {
		panic(fmt.Sprintf("No hay suficientes imagenes en el tema %s", tema))
	}

	var imagenesElegidas []Imagen

	for i := 0; i < cantidad; {

		// obtener un numero aleatorio entre 0 y la cantidad de archivos
		indiceAleatorio := rand.Intn(len(imagenes))

		// verificar si el archivo es una imagen
		if !tieneIndicesRepetidos(imagenesElegidas, indiceAleatorio) {

			imagen := imagenes[indiceAleatorio]

			// convertir la imagen a base64
			imagenEncriptada := convertirABase64(fmt.Sprintf("%s/%s/%s", rutaImagenes, tema, imagen.Name()))
			// agregar la imagen al slice de imagenes
			imagenesElegidas = append(imagenesElegidas, Imagen{
				Nombre:    imagen.Name(),
				Contenido: imagenEncriptada,
				Indice:    indiceAleatorio,
				Extension: obtenerExtension(imagen),
			})

			i++
		}
	}

	// retornar un slice de ImageContainer con las imagenes
	imagenContenedor := ImagenContenedor{Imagenes: imagenesElegidas}

	return imagenContenedor
}

func tieneIndicesRepetidos(imagenes []Imagen, indice int) bool {
	for _, imagen := range imagenes {
		if imagen.Indice == indice {
			return true
		}
	}
	return false
}

func esImagen(archivo os.DirEntry) bool {
	if archivo.IsDir() {
		return false
	}

	for _, extension := range extensionesPermitidas {
		if extension == obtenerExtension(archivo) {
			return true
		}
	}
	return false
}

func obtenerExtension(archivo os.DirEntry) string {
	return filepath.Ext(archivo.Name())
}

func convertirABase64(rutaCompleta string) string {
	imagen, err := os.Open(rutaCompleta)
	if err != nil {
		panic(err)
	}
	defer imagen.Close()

	fileInfo, _ := imagen.Stat()
	size := fileInfo.Size()
	buffer := make([]byte, size)

	imagen.Read(buffer)
	imagenBase64 := base64.StdEncoding.EncodeToString(buffer)

	return imagenBase64
}

func filter(ss []os.DirEntry, test func(os.DirEntry) bool) (ret []os.DirEntry) {
	for _, s := range ss {
		if test(s) {
			ret = append(ret, s)
		}
	}
	return
}
