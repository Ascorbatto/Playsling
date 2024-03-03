package services

import (
	"log"
	"net/http"

	"golang.org/x/oauth2"
)

type AuthConfig struct {
	Config *oauth2.Config
}

type AuthService interface {
	HandleLogin(w http.ResponseWriter, r *http.Request)
	HandleCallback(w http.ResponseWriter, r *http.Request) *http.Client
}

func NewAuthConfig(config *oauth2.Config) *AuthConfig {
	return &AuthConfig{Config: config}
}

func (as *AuthConfig) HandleLogin(w http.ResponseWriter, r *http.Request) {
	url := as.Config.AuthCodeURL("random")
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func (as *AuthConfig) HandleCallback(w http.ResponseWriter, r *http.Request) *http.Client {
	if r.FormValue("state") != "random" {
		log.Println("invalid oauth2 state")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return nil
	}

	token, err := as.Config.Exchange(r.Context(), r.FormValue("code"))
	if err != nil {
		log.Fatalf("code exchange wrong: %s", err.Error())
	}

	return as.Config.Client(r.Context(), token)
}
