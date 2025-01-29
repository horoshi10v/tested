package app

import (
	"01-server/internal/repository"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	httpSwagger "github.com/swaggo/http-swagger"

	"01-server/internal/config"
	"01-server/internal/handlers"
	"01-server/internal/middleware"
)

func RunServer(cfg *config.Config) error {
	db, err := sql.Open("postgres", cfg.DSN())
	if err != nil {
		return err
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		return err
	}
	log.Println("DB connected")

	userRepo := repository.NewUserRepo(db)
	sellerRepo := repository.NewSellerRepo(db)

	authHandler := handlers.NewAuthHandler(userRepo, cfg.JWTSecret)
	sellerHandler := handlers.NewSellerHandler(sellerRepo)

	mux := http.NewServeMux()

	// Swagger
	mux.Handle("/swagger/", httpSwagger.WrapHandler)

	// Auth
	mux.HandleFunc("/register", authHandler.Register)
	mux.HandleFunc("/login", authHandler.Login)
	mux.HandleFunc("/logout", authHandler.Logout)

	// Sellers
	mux.Handle("/sellers", middleware.CheckAuth(sellerHandler, cfg.JWTSecret))
	mux.Handle("/sellers/", middleware.CheckAuth(sellerHandler, cfg.JWTSecret))

	addr := fmt.Sprintf(":%s", cfg.ServerPort)
	log.Printf("Server listening on %s", addr)
	return http.ListenAndServe(addr, mux)
}
