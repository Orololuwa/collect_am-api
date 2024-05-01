package main

import (
	"net/http"

	"github.com/Orololuwa/collect_am-api/src/config"
	"github.com/Orololuwa/collect_am-api/src/driver"
	"github.com/Orololuwa/collect_am-api/src/dtos"
	"github.com/Orololuwa/collect_am-api/src/handlers"
	middleware "github.com/Orololuwa/collect_am-api/src/middleware"
	"github.com/go-chi/chi/v5"
	middlewareChi "github.com/go-chi/chi/v5/middleware"
)

func routes(a *config.AppConfig, conn *driver.DB) http.Handler {
	// Initialize internal middlewares
	md := middleware.New(a, conn)	

	// 
	mux := chi.NewRouter()

	// middlewares
	mux.Use(middlewareChi.Logger)

	mux.Get("/health", handlers.Repo.Health)

	// auth
	mux.Post("/auth/signup", handlers.Repo.SignUp)
	mux.Post("/login", md.ValidateReqBody(http.HandlerFunc(handlers.Repo.LoginUser), &dtos.UserLoginBody{} ).ServeHTTP)

	// protected route
	mux.Get("/protected-route", md.Authorization(http.HandlerFunc(handlers.Repo.ProtectedRoute)).ServeHTTP)

	return mux;
}