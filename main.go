package main

import (
	"codogenerator/handlers"
	ogenspec "codogenerator/spec"
	"codogenerator/storage"
	"context"
	"log"
	"net/http"
	"os"
)

// It sets up the logging, connects to the database, creates the service instance,
// and starts the HTTP server.
func main() {
	// Set up logging
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		log.SetOutput(file)
	} else {
		log.Println("Не удалось открыть файл логов, используется стандартный stderr")
	}
	defer file.Close()

	// Connect to the database
	// The URL for the database connection is read from the environment variable "MYURL".
	myURL := "MYURL"
	ctx := context.Background()
	conn, err := storage.ConnectToDB(ctx, myURL)
	if err != nil {
		log.Fatalf("error connecting to the database: %s", err.Error())
	}
	defer conn.Close(ctx)

	// Create the repository instance
	// The repository is responsible for interacting with the database.
	repos := storage.NewExpenseRepo(conn)

	// Create the handler instance
	// The handler is responsible for handling the API requests.
	handler := handlers.NewHandler(repos)

	// Create the service instance
	// The service is responsible for serving the API endpoints.
	srv, err := ogenspec.NewServer(handler, &handlers.SecurityServer{})
	if err != nil {
		log.Fatal("error creating the server")
	}

	// Start the HTTP server
	// The server listens on port 8080 and serves the API endpoints.
	if err := http.ListenAndServe(":8080", srv); err != nil {
		log.Fatal(err)
	}
}
