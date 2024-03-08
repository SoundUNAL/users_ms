package model

import (
	"net/http"
	"time"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	// Eliminar el token JWT almacenado en el cliente
	// Por ejemplo, puedes eliminar una cookie que almacena el token
	http.SetCookie(w, &http.Cookie{
		Name:    "jwt",
		Value:   "",
		Expires: time.Unix(0, 0),
	})

	// Responder con un mensaje de éxito
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Logout exitoso"))
}
