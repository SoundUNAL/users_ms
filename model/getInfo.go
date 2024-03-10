package model

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"users/utils"

	"github.com/go-chi/chi/v5"
)

type UserInfo struct {
	Name      string `json:"name"`
	LastName  string `json:"lastname"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Birthday  string `json:"birthday"`
	Role      string `json:"role"`
	CreatedAt string `json:"created_at"`
}

func GetUserByID(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	// Obtener el ID del usuario de la URL
	userID := chi.URLParam(r, "id")

	// Preparar la consulta SQL para obtener la información del usuario por su ID
	query := "SELECT  name, lastname, username, email, phone, birthday, lastconnection, idRol  FROM User WHERE idUser = ?"

	var user UserInfo
	//var username string
	//query := "SELECT  name FROM User WHERE idUser = ?"
	// Ejecutar la consulta SQL y obtener las filas resultantes

	var birthdayHelp []uint8
	var lastLog []uint8
	err := db.QueryRow(query, userID).Scan(&user.Name, &user.LastName, &user.Username, &user.Email, &user.Phone, &birthdayHelp, &lastLog, &user.Role)
	if err != nil {
		http.Error(w, "Error al obtener la información del usuario de la base de datos", http.StatusInternalServerError)
		println(err.Error())
		return
	}

	user.Birthday = utils.SinceUINTtoText(birthdayHelp)
	user.CreatedAt = utils.SinceUINTtoText(lastLog)

	// Devolver la información del usuario como respuesta
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}
