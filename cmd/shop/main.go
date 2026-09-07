package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"shop/internal/catalog"
)

func main() {
	shutdownCtx, shutdownCancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer shutdownCancel()

	mux := InitRoutes(nil)

	go func() {
		log.Println("Starting server on :8080")
		if err := http.ListenAndServe(":8080", mux); err != nil {
			fmt.Println("server listen error:", err)
			os.Exit(1)
		}
	}()

	<-shutdownCtx.Done()
	log.Println("Received shutdown signal, shutting down.")
	log.Println("Server shutdown gracefully.")
}

func InitRoutes(catalogService *catalog.Service) *http.ServeMux {
	mux := http.NewServeMux()

	// MAIN HANDLER
	mux.HandleFunc("/health",
		func(w http.ResponseWriter, r *http.Request) {
			log.Println("GET /health")
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte("main healthy"))
		},
	)

	// CATALOG HANDLER
	catalogHandler := catalog.New(nil)
	mux.HandleFunc("/catalog/create", catalogHandler.CreateProducts)
	mux.HandleFunc("/catalog/product/{id}", catalogHandler.GetProduct)
	mux.HandleFunc("/catalog/products", catalogHandler.GetProducts)
	mux.HandleFunc("/catalog/health", catalogHandler.Health)

	return mux
}
