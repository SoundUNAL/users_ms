package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("superSecretKey")

func GenerateJWT(user string) (string, error) {
	// Crear el payload del token
	claims := jwt.MapClaims{
		"username":  user,
		"exp":       time.Now().Add(time.Hour * 24).Unix(), // Expira en 24 horas
		"issued_at": time.Now().Unix(),
	}

	// Crear el token JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Firmar el token con la clave secreta
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func HashPassword(password string) string {
	// Crear un nuevo hash SHA-256
	hash := sha256.New()

	// Escribir la contraseña en el hash
	hash.Write([]byte(password))

	// Obtener la suma de comprobación del hash
	hashed := hash.Sum(nil)

	// Convertir la suma de comprobación en una cadena hexadecimal
	hashedPassword := hex.EncodeToString(hashed)

	return hashedPassword
}
