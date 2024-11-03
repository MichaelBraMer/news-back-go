package main

import (
	"context"
	"log"
	"net/http"
	"news-back-go/internal/app/services"
	"news-back-go/internal/infrastructure/db"
	"news-back-go/internal/infrastructure/httpHandler"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func disableCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Manejar preflight requests (opciones)
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		h.ServeHTTP(w, r)
	})
}

func main() {
	r := mux.NewRouter()

	log.Println("Iniciando servidor en el puerto 8080...")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Conexion BD establecida")
	}

	repo := db.NewMongoNewsRepository(client, "newsDB", "news")
	service := services.NewNewsService(repo)
	handler := httpHandler.NewNewsHandler(service)
	r.HandleFunc("/news", handler.CreateNews).Methods("POST")
	r.HandleFunc("/news/{id}", handler.GetByIdNews).Methods("GET")
	r.HandleFunc("/news", handler.GetAllNews).Methods("GET")
	r.HandleFunc("/news/{id}", handler.DeleteNews).Methods("DELETE")
	r.HandleFunc("/news/{id}", handler.UpdateNews).Methods("PUT")

	err = http.ListenAndServe("localhost:8080", disableCORS(r))

	if err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
