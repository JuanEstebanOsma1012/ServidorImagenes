package main

import (
	"html/template"
	"net/http"
	"os"
	"strconv"
)

func main() {

	if len(os.Args) != 4 {
		panic("Se esperaban tres argumentos")
	}

	puerto := os.Args[1]

	http.HandleFunc("/", getImages)
	http.ListenAndServe(":"+puerto, nil)
}

func getImages(w http.ResponseWriter, r *http.Request) {

	tema := os.Args[2]
	cantidadImagenes, _ := strconv.Atoi(os.Args[3])

	tmpl, _ := template.ParseFiles("../estaticos/index.html")
	tmpl.Execute(w, obtenerImagenesPorTema(tema, cantidadImagenes))

}
