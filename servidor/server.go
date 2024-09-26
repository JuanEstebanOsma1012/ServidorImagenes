package main

import (
	"fmt"
	"html/template"
	"math/rand"
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

	randomNumber := rand.Intn(2)

	tmpl, _ := template.ParseFiles(fmt.Sprintf("../estaticos/index%d.html", randomNumber+1))
	tmpl.Execute(w, obtenerImagenesPorTema(tema, cantidadImagenes))

}
