package main

import (
	"fmt"
	"net/http"

	"github.com/JDoornink/go-rest-api-course-2/internal/database"
	transportHTTP "github.com/JDoornink/go-rest-api-course-2/internal/transport/http"
)

// App - the struct which contains things like
// pointers to database connections
type App struct{}

// Run -  sets up our application
func (app *App) Run() error {
	fmt.Println("Setting Up Our APP")

	var err error
	_, err = database.NewDatabase()
	if err != nil {
		return err
	}

	handler := transportHTTP.NewHandler()
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8081", handler.Router); err != nil {
		fmt.Println("Failed to set up server")
		return err
	}
	return nil
}

func main() {
	fmt.Println("Go REST API Course")
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error starting up your REST API")
		fmt.Println(err)
	}
}
