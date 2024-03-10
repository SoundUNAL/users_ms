package model

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"
	"users/utils"

	"github.com/go-chi/chi/v5"
	_ "github.com/go-sql-driver/mysql"
)

// User representa la estructura de un usuario
type UserData struct {
	Name     string    `json:"name"`
	Lastname string    `json:"lastname"`
	Username string    `json:"username"`
	Password string    `json:"password"`
	Email    string    `json:"email"`
	Phone    string    `json:"phone"`
	Birthday time.Time `json:"birthday"`
}

// updateUser actualiza la información de un usuario en la base de datos
func UpdateUserHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	// Obtener el ID del usuario de los parámetros de la solicitud
	userID := chi.URLParam(r, "id")

	var updatedUser UserData
	// Decodificar el cuerpo de la solicitud JSON en una estructura User
	if err := json.NewDecoder(r.Body).Decode(&updatedUser); err != nil {
		http.Error(w, "Error al decodificar el cuerpo de la solicitud\n", http.StatusBadRequest)
		return
	}

	query := "UPDATE User SET"
	var params []interface{}
	if updatedUser.Name != "" {
		query += " name=?,"
		params = append(params, updatedUser.Name)
	}
	if updatedUser.Lastname != "" {
		query += " lastname=?,"
		params = append(params, updatedUser.Lastname)
	}
	if updatedUser.Password != "" {
		query += " password=?,"
		updatedUser.Password = utils.HashPassword(updatedUser.Password)
		params = append(params, updatedUser.Password)
	}
	if updatedUser.Email != "" {

		query += " email=?,"
		params = append(params, updatedUser.Email)
	}
	if updatedUser.Phone != "" {
		query += " phone=?,"
		params = append(params, updatedUser.Phone)
	}
	if !updatedUser.Birthday.IsZero() {
		query += " birthday=?,"
		params = append(params, updatedUser.Birthday)
	}
	// Eliminar la coma adicional al final de la consulta
	query = query[:len(query)-1]

	// Agregar el condicional WHERE para actualizar el usuario específico
	query += " WHERE idUser=?"
	params = append(params, userID)

	// Ejecutar la consulta
	_, err := db.Exec(query, params...)
	if err != nil {
		print(err.Error())
		http.Error(w, "Error al actualizar el usuario en la base de datos\n", http.StatusInternalServerError)
		return
	}

	// Responder con un mensaje de éxito
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Información del usuario actualizada correctamente\n"))
}
