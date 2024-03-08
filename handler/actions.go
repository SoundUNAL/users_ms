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
	fmt.Fprintf(w, "Create User\n")
	model.CreateUserHandler(db, w, r)
}

func (a *Action) Login(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Login Action\n")
	model.LoginHandler(db, w, r)
}

func (a *Action) Logout(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Logout Action")
	model.LogoutHandler(w, r)
}

func (a *Action) Update(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Update Action")
}

func (a *Action) Delete(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Delete Action")
}
