package model

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"users/utils"
	//"users/application"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Name      string    `json:"name"`
	LastName  string    `json:"lastname"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Birthday  time.Time `json:"birthday"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}

func CreateUserHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	// Decodificar la información del usuario desde el cuerpo de la solicitud
	var user User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Establecer la fecha y hora actual
	user.CreatedAt = time.Now()

	//Cifrar la contraseña
	user.Password = utils.HashPassword(user.Password)

	if user.Role != "2" && user.Role != "3" {
		fmt.Fprintf(w, "Invalid Data")
		return
	}
	// Insertar el nuevo usuario en la base de datos
	_, err = db.Exec("INSERT INTO User (name, lastname, username, password,  email, phone, birthday, idRol, lastconnection) VALUES ( ?, ?, ?, ?,?, ?, ?, ? , ?)",
		user.Name, user.LastName, user.Username, user.Password, user.Email, user.Phone, user.Birthday, user.Role, user.CreatedAt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respuesta exitosa
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Usuario creado correctamente")
}
