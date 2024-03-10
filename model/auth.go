package model

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"
	"users/utils"

	"golang.org/x/crypto/bcrypt"
)

type UserAuth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func authenticateUser(db *sql.DB, username, password string) (*User, error) {

	// Consultar la base de datos para obtener la contraseña almacenada del usuario

	var storedPassword string
	err := db.QueryRow("SELECT password FROM User WHERE username = ?", username).Scan(&storedPassword)
	if err != nil {
		return nil, err
	}

	// Verificar la contraseña almacenada en la base de datos con la contraseña proporcionada
	//println("Stored Password")
	//println(storedPassword)

	password = utils.HashPassword(password)
	//println(password)
	if password != storedPassword {
		//Puedes darle a la variable err un valor personalizado
		err = bcrypt.ErrMismatchedHashAndPassword
		return nil, err
	}

	// Devolver solo el nombre de usuario si la autenticación es exitosa
	return &User{Username: username, Password: password}, nil
}

func LoginHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	// Decodificar la información del usuario desde el cuerpo de la solicitud
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Verificar las credenciales del usuario
	authenticatedUser, err := authenticateUser(db, user.Username, user.Password)
	if err != nil {
		http.Error(w, "Credenciales inválidas", http.StatusUnauthorized)
		return
	}

	query := "UPDATE User SET  lastconnection = ? WHERE Username = ?"
	_, err = db.Exec(query, time.Now(), user.Username)

	if err != nil {
		http.Error(w, "Error al actualizar el campo del usuario en la base de datos", http.StatusInternalServerError)
		return
	}
	// Generar un token JWT si las credenciales son válidas
	tokenString, err := utils.GenerateJWT(authenticatedUser.Username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Devolver el token JWT al cliente
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})

	http.SetCookie(w, &http.Cookie{
		Name:     "jwt",
		Value:    tokenString,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
	})

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Login exitoso"))

}
