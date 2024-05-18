package main

import (
	"net/http"

	"github.com/Orololuwa/collect_am-api/src/config"
	"github.com/Orololuwa/collect_am-api/src/driver"
	"github.com/Orololuwa/collect_am-api/src/handlers"
	middleware "github.com/Orololuwa/collect_am-api/src/middleware"
	"github.com/go-chi/chi/v5"
	middlewareChi "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func routes(a *config.AppConfig, conn *driver.DB) http.Handler {
	// Initialize internal middlewares
	md := middleware.New(a, conn)	

	// 
	mux := chi.NewRouter()

	// middlewares
	mux.Use(middlewareChi.Logger)
	mux.Use(middlewareChi.Recoverer)

	corsMiddleware := cors.New(cors.Options{
        AllowedOrigins:   []string{"*"},
        AllowedMethods:   []string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"},
        AllowedHeaders:   []string{"*"},
        AllowCredentials: true,
        Debug:            true,
    })
	mux.Use(corsMiddleware.Handler)

	// 
	mux.Get("/health", handlers.Repo.Health)

	// auth
	mux.Post("/auth/signup", handlers.Repo.SignUp)
	mux.Post("/auth/login", handlers.Repo.LoginUser)

	// Authenticated Routes
	mux.With(md.Authorization).Group(func(r chi.Router) {
		//business
		r.Post("/business", handlers.Repo.AddBusiness)
		r.Get("/business", handlers.Repo.GetBusiness)
		r.Patch("/business", handlers.Repo.UpdateBusiness)

		// misc
		r.Get("/protected-route", handlers.Repo.ProtectedRoute)
	})


	return mux;
}