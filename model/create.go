package model

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	//"users/application"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Name      string    `json:"name"`
	LastName  string    `json:"lastname"`
	Username  string    `json:"username"`
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

	// db, err := sql.Open("mysql", "root:madlies@tcp(127.0.0.1:3306)/soundunal_users_db")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// Verificar la conexión a MySQL
	// err = db.Ping()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//fmt.Fprintf(w, "Usuario creado correctamente")
	//fmt.Fprintf(w, "Conexión a la base de datos exitosa")

	// Establecer la fecha y hora actual
	user.CreatedAt = time.Now()

	// Insertar el nuevo usuario en la base de datos
	_, err = db.Exec("INSERT INTO User (name, lastname, username, email, phone, birthday, idRol, lastconnection) VALUES ( ?, ?, ?, ?, ?, ?, ? , ?)",
		user.Name, user.LastName, user.Username, user.Email, user.Phone, user.Birthday, user.Role, user.CreatedAt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respuesta exitosa
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Usuario creado correctamente")
}
