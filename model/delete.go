package model

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func DeleteUserHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "id")

	_, err := db.Exec("DELETE FROM User WHERE idUser=?", userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Usuario eliminado correctamente"))

}
