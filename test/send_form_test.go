package test

import (
	"fmt"
	"net/http"
)

func enviarFormulario(w http.ResponseWriter, r *http.Request) {
	// Obtener los valores de los campos del formulario
	nombre := r.FormValue("nombre")
	email := r.FormValue("email")
	mensaje := r.FormValue("mensaje")

	// Hacer algo con los valores (por ejemplo, enviar un correo electrónico)

	fmt.Println("DATOS ENVIADOS: ", nombre, email, mensaje)

	// Redirigir al usuario a una página de confirmación
	// http.Redirect(w, r, "/confirmacion.html", http.StatusSeeOther)
}
