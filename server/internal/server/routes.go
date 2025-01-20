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
	user, err := gothic.GetFromSession("user", r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	w.Write([]byte(user))
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

	encodedUser, _ := json.Marshal(user)
	gothic.StoreInSession("user", string(encodedUser), r, w)

	// create user if one doesn't exist
	rows, _ := s.db.Query(context.Background(), "SELECT FROM users WHERE email=$1;", user.Email)
	defer rows.Close()
	if !rows.Next() {
		s.db.CreateUser(user.Email)
	}

	http.Redirect(w, r, config.WebUrl, http.StatusFound)
}

func (s *Server) HandleSignout(w http.ResponseWriter, r *http.Request) {
	gothic.Logout(w, r)
	w.Header().Set("Location", config.WebUrl)
	w.WriteHeader(http.StatusTemporaryRedirect)
}
