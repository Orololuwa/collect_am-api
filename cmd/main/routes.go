package main

import (
	"net/http"

	"github.com/Orololuwa/collect_am-api/src/config"
	v1 "github.com/Orololuwa/collect_am-api/src/controllers/v1"
	"github.com/Orololuwa/collect_am-api/src/driver"
	"github.com/Orololuwa/collect_am-api/src/handlers"
	middleware "github.com/Orololuwa/collect_am-api/src/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func routes(a *config.AppConfig, h handlers.HandlerFunc, conn *driver.DB) http.Handler {
	// Initialize internal middlewares
	md := middleware.New(a, conn)
	v1Routes := v1.NewController(a, h)

	// 
	mux := chi.NewRouter()

	// middlewares
	// mux.Use(middlewareChi.Logger)
	// mux.Use(middlewareChi.Recoverer)

	corsMiddleware := cors.New(cors.Options{
        AllowedOrigins:   []string{"*"},
        AllowedMethods:   []string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"},
        AllowedHeaders:   []string{"*"},
        AllowCredentials: true,
        Debug:            false,
    })

	// 
	mux.Get("/health", v1Routes.Health)

	mux.Route("/api/v1", func(v1Router chi.Router) {
		v1Router.Use(corsMiddleware.Handler)

		// auth
		v1Router.Post("/auth/signup", v1Routes.SignUp)
		v1Router.Post("/auth/login", v1Routes.LoginUser)

		// Authenticated Routes
		v1Router.With(md.Authorization).Group(func(r chi.Router) {
			//business
			r.Post("/business", v1Routes.AddBusiness)
			r.Get("/business/{id}", v1Routes.GetBusiness)
			r.Patch("/business/{id}", v1Routes.UpdateBusiness)
		})

	})



	return mux;
}