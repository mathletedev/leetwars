package server

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/markbates/goth/gothic"
	"github.com/mathletedev/leetwars/internal/config"
)

func (s *Server) RegisterRoutes(allowedOrigins []string) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   allowedOrigins,
		AllowCredentials: true,
	}))

	r.Get("/api/hello", s.HandleHello)
	r.Get("/api/me", s.HandleMe)
	r.Get("/auth/{provider}", s.HandleAuth)
	r.Get("/auth/{provider}/callback", s.HandleAuthCallback)
	r.Get("/signout/{provider}", s.HandleSignout)

	return r
}

func (s *Server) HandleHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}

func (s *Server) HandleMe(w http.ResponseWriter, r *http.Request) {
	id, err := gothic.GetFromSession("user", r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	user, err := s.db.ReadUser(id)
	if err != nil {
		if err.Error() == "no username" {
			// https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/418
			w.WriteHeader(http.StatusTeapot)
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (s *Server) HandleAuth(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")
	r = r.WithContext(context.WithValue(context.Background(), "provider", provider))

	gothic.BeginAuthHandler(w, r)
}

func (s *Server) HandleAuthCallback(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")
	r = r.WithContext(context.WithValue(context.Background(), "provider", provider))

	user, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		log.Println(err)
		fmt.Fprintln(w, r)
		return
	}

	rows, err := s.db.Query(context.Background(), "SELECT id FROM users WHERE email=$1;", user.Email)
	if err != nil {
		log.Println(err)
		fmt.Fprintln(w, r)
		return
	}

	defer rows.Close()

	var id string
	if rows.Next() {
		rows.Scan(&id)
	} else {
		// create user if one doesn't exist
		id, err = s.db.CreateUser(user.Email)
		if err != nil {
			log.Println(err)
			fmt.Fprintln(w, r)
			return
		}
	}

	// set user id in session
	gothic.StoreInSession("user", id, r, w)

	http.Redirect(w, r, config.WebUrl, http.StatusFound)
}

func (s *Server) HandleSignout(w http.ResponseWriter, r *http.Request) {
	gothic.Logout(w, r)
	w.Header().Set("Location", config.WebUrl)
	w.WriteHeader(http.StatusTemporaryRedirect)
}
