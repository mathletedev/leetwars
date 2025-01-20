package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/sessions"
	"github.com/mathletedev/leetwars/internal/auth"
	"github.com/mathletedev/leetwars/internal/config"
	"github.com/mathletedev/leetwars/internal/db"
)

type Server struct {
	port  int
	db    *db.Database
	store *sessions.CookieStore
}

func NewServer(prod *bool) *http.Server {
	envPort, present := os.LookupEnv("PORT")
	if !present {
		envPort = "8080"
	}

	port, err := strconv.Atoi(envPort)
	if err != nil {
		log.Println("invalid port, using default")
	}

	allowedOrigins := []string{"*"}
	if !*prod {
		allowedOrigins = []string{config.WebUrl}
	}

	store := sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))
	store.Options.Path = "/"
	store.Options.HttpOnly = true
	store.Options.Secure = *prod
	store.Options.SameSite = http.SameSiteLaxMode
	store.Options.MaxAge = auth.MaxAge

	s := &Server{
		port:  port,
		db:    db.NewDatabase(),
		store: store,
	}

	return &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: s.RegisterRoutes(allowedOrigins),
	}
}
