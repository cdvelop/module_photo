package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		method, object := getMethodAndObjectFromPath(r.URL.Path)

		switch method {

		case "create":
			fmt.Fprintf(w, "Estás en la página de creación del objeto %s", object)

		case "delete":
			fmt.Fprintf(w, "Estás en la página de eliminación del objeto %s", object)

		case "update":
			fmt.Fprintf(w, "Estás en la página de actualización del objeto %s", object)

		case "read":
			id := r.URL.Query().Get("id")
			if id != "" {
				fmt.Fprintf(w, "Estás en la página de lectura del objeto %s con ID %s", object, id)
			} else {
				fmt.Fprintf(w, "Estás en la página de lectura del objeto %s, pero no se proporcionó ningún ID.", object)
			}

		default:
			if r.URL.Path == "/" {
				fmt.Fprint(w, "¡Hola! Esta es la página principal.")
			} else {
				http.NotFound(w, r)
			}
		}
	})

	fmt.Println("Servidor escuchando en http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getMethodAndObjectFromPath(path string) (string, string) {
	parts := strings.Split(path, "/")
	if len(parts) < 3 || parts[1] == "" {
		return "", ""
	}
	return parts[1], parts[2]
}
