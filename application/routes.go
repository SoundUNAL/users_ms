package application

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"users/handler"
)

func loadRoutes(db *sql.DB) *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Happy Hacking!!!"))
	})

	router.Route("/api/users/", func(router chi.Router) { loadActionRoutes(db, router) })

	return router
}

func dbInjector(db *sql.DB, h func(*sql.DB, http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) { h(db, w, r) }
}

func loadActionRoutes(db *sql.DB, router chi.Router) {
	actionHandler := &handler.Action{}

	router.Post("/signup", dbInjector(db, actionHandler.Create))
	router.Post("/login", dbInjector(db, actionHandler.Login))
	router.Get("/logout", actionHandler.Logout)
	router.Put("/{id}", dbInjector(db, actionHandler.Update))
	router.Delete("/{id}", actionHandler.Delete)
}
