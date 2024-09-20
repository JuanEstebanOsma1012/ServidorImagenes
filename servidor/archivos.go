package main

import (
	"encoding/base64"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
)

type Image struct {
	Name    string
	Content string
	Index   int
}

var allowedExtensions = []string{".jpg", ".jpeg", ".png"}

const (
	imagesPath = "../estaticos/imagenes"
)

func main() {
	// obtener las imagenes de un tema
	images := getImagesByTopic("animales", 2)

	fmt.Println("Imagenes obtenidas:")

	for _, image := range images {
		fmt.Println(image)
	}
}

func getImagesByTopic(topic string, number int) []Image {

	files, err := os.ReadDir(fmt.Sprintf("%s/%s", imagesPath, topic))
	if err != nil {
		panic(err)
	}

	filesFilter := filter(files, func(file os.DirEntry) bool {
		return isImage(file)
	})

	if len(filesFilter) < number {
		panic(fmt.Sprintf("No hay suficientes imagenes en el tema %s", topic))
	}

	var images []Image

	for i := 0; i < number; {

		// obtener un numero aleatorio entre 0 y la cantidad de archivos
		randomIndex := rand.Intn(len(filesFilter))

		// verificar si el archivo es una imagen
		if !isIndexRepeated(images, randomIndex) {

			file := filesFilter[randomIndex]

			if isImage(file) {

				// convertir la imagen a base64
				image := convertImageToBase64(fmt.Sprintf("%s/%s/%s", imagesPath, topic, file.Name()))
				// agregar la imagen al slice de imagenes
				images = append(images, Image{
					Name:    file.Name(),
					Content: image,
					Index:   randomIndex,
				})

				i++

			}
		}
	}

	return images
}

func isIndexRepeated(images []Image, index int) bool {
	for _, image := range images {
		if image.Index == index {
			return true
		}
	}
	return false
}

func isImage(file os.DirEntry) bool {
	if file.IsDir() {
		return false
	}

	for _, extension := range allowedExtensions {
		if extension == filepath.Ext(file.Name()) {
			return true
		}
	}
	return false
}

func convertImageToBase64(completePath string) string {
	image, err := os.Open(completePath)
	if err != nil {
		panic(err)
	}
	defer image.Close()

	fileInfo, _ := image.Stat()
	size := fileInfo.Size()
	buffer := make([]byte, size)

	image.Read(buffer)
	imageBase64 := base64.StdEncoding.EncodeToString(buffer)

	return imageBase64
}

func filter(ss []os.DirEntry, test func(os.DirEntry) bool) (ret []os.DirEntry) {
	for _, s := range ss {
		if test(s) {
			ret = append(ret, s)
		}
	}
	return
}
