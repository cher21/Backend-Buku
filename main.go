package main

import (
	"crud/router"
	"fmt"
	"log"
	"net/http"

	"github.com/rs/cors"
)

func main() {
	r := router.Router()

	// Konfigurasi CORS
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"}, // Atur sesuai dengan asal permintaan Anda
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type"},
		AllowCredentials: true,
	})

	// Pasang handler CORS pada router Anda
	handler := c.Handler(r)

	fmt.Println("Server dijalankan pada port 8080...")

	// Mulai server Anda
	log.Fatal(http.ListenAndServe(":8080", handler))
}
