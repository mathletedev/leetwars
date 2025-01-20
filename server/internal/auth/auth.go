package auth

import (
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
)

const MaxAge = 60 * 60 * 24 * 7

func NewAuth(prod *bool) {
	sessionKey := os.Getenv("SESSION_KEY")

	googleClientId := os.Getenv("GOOGLE_CLIENT_ID")
	googleClientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")

	store := sessions.NewCookieStore([]byte(sessionKey))
	store.Options.Path = "/"
	store.Options.HttpOnly = true
	store.Options.Secure = *prod
	store.Options.SameSite = http.SameSiteLaxMode
	store.Options.MaxAge = MaxAge

	gothic.Store = store

	goth.UseProviders(
		google.New(
			googleClientId,
			googleClientSecret,
			"http://localhost:8080/auth/google/callback",
		),
	)
}
