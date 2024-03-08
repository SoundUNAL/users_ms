package application

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	//log"
	"net/http"
	"time"

	//"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type App struct {
	router http.Handler
	db     *sql.DB
}

func New() *App {

	db, err := sql.Open("mysql", "root:madlies@tcp(127.0.0.1:3306)/soundunal_users_db")
	if err != nil {

		log.Fatal(err)
	}

	// Verificar la conexión a MySQL
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	//Cerrar la conexión al salir de la función

	app := &App{
		router: loadRoutes(db),
		db:     db,
	}

	return app
}

func (a *App) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:    ":8080",
		Handler: a.router,
	}

	defer func() {
		if err := a.db.Close(); err != nil {
			fmt.Printf("Error al cerrar la conexión a MySQL: %v\n", err)
		}
		print("Conexión cerrada")
	}()

	fmt.Println("Starting server on :8080")

	channel := make(chan error, 1)

	go func() {
		err := server.ListenAndServe()

		if err != nil {
			channel <- fmt.Errorf("server error: %w", err)
		}
		close(channel)
	}()

	select {
	case err := <-channel:
		return err
	case <-ctx.Done():
		timeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		return server.Shutdown(timeout)
	}

	return nil
}
