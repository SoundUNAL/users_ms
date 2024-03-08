package handler

import (
	"database/sql"
	"fmt"
	"net/http"
	"users/model"
)

type Action struct {
}

func (a *Action) Create(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Create Action ")
	model.CreateUserHandler(db, w, r)
	fmt.Fprintf(w, "Usuario creado correctamente11")
}

func (a *Action) Login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Login Action")
}

func (a *Action) Logout(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Logout Action")
}

func (a *Action) Update(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Update Action")
}

func (a *Action) Delete(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Delete Action")
}
