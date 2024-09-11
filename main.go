package main

import (
	"Sneaker_Inventory/handlers"
	"Sneaker_Inventory/models"
	"Sneaker_Inventory/repository"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"

	// "gorm.io/driver/sqlite"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {
	var err error
	// Connect to the database using gorm. The database file is gorm.db
	db, err = gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Migrate the database schema
	db.AutoMigrate(&models.Sneaker{}, &models.SoldSneaker{})

	// Create a new chi router
	r := chi.NewRouter()
	// Use the middleware to log requests
	r.Use(middleware.Logger)
	// Use the middleware for CORS
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	// Create a new instance of the sneaker repository
	sneakerRepo := repository.NewSneakerRepository(db)
	// Create a new instance of the sneaker handler
	sneakerHandler := handlers.NewHandler(sneakerRepo)

	// Define the routes
	r.Get("/api/sneakers", sneakerHandler.GetSneakers)
	r.Get("/api/sneakers/{id}", sneakerHandler.GetSneaker)
	r.Post("/api/sneakers", sneakerHandler.AddSneaker)
	r.Put("/api/sneakers/{id}", sneakerHandler.UpdateSneaker)
	r.Delete("/api/sneakers/{id}", sneakerHandler.DeleteSneaker)

	r.Post("/api/sneakers/{id}/sell", sneakerHandler.SellSneaker)
	r.Get("/api/sold-sneakers", sneakerHandler.SoldSneakers)

	// SPA from ReactJS
	r.Handle("/*", http.FileServer(http.Dir("frontend/dist")))

	// Start the server
	log.Println("Starting server on :http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
