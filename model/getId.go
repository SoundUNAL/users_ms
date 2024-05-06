package model

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"github.com/go-chi/chi/v5"
)


type UserId struct {
	Id      string `json:"id"`
}

func  GetUserIDHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	// Obtener el ID del usuario de la URL
	username := chi.URLParam(r, "username")

	// Preparar la consulta SQL para obtener la información del usuario por su ID
	query := "SELECT  idUser FROM User WHERE username = ?"

	var user UserId


	// Ejecutar la consulta SQL y obtener las filas resultantes

	err := db.QueryRow(query, username).Scan(&user.Id)
	if err != nil {
		http.Error(w, "Error al obtener la información del usuario de la base de datos", http.StatusInternalServerError)
		println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}


	// Devolver la información del usuario como respuesta
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}
