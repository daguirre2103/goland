package main

import (
	"net/http"
)

func main() {

	/* Definimos la ruta que llamará la función home */
	http.HandleFunc("/", home)

	http.ListenAndServe(":8080", nil)

}

/* Definimos la función home */
func home(w http.ResponseWriter, r *http.Request) {
	html := "<html>"
	html += "<body>"
	html += "<h1>Hola Mundo</h1>"
	html += "</body>"
	html += "</html>"
	w.Write([]byte(html))
}
