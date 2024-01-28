package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

// routes generates our routes and attaches them to handlers, using the chi router
// note that we return type http.Handler, and not *chi.Mux; since chi.Mux satisfies
// the interface requirements for http.Handler, it makes sense to return the type
// that is part of the standard library.
func (app *application) routes() http.Handler {
	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	// LOGIN REQUESTS
	mux.Get("/", app.serveLogin)
	mux.Post("/", app.serveLogin)
	mux.Get("/api/userlogin", app.UserloginAPI)
	mux.Post("/api/userlogin", app.UserloginAPI)
	mux.Post("/api/adminlogin", app.AdminloginAPI)

	// USER REGISTRATION REQUESTS
	// mux.Get("/api/user-registration", app.serveUserRegistrationPage)
	// mux.Post("/api/user-registration", app.registerUser)

	//OTP-Working
	mux.Get("/api/otp-verfication", app.OTP_verfication)
	mux.Post("/api/otp-verfication", app.OTP_verfication)

	mux.Get("/api/verify-otp", app.VerifyOTP)
	mux.Post("/api/verify-otp", app.VerifyOTP)
	return mux
}
