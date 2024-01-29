package services

import (
	models "Conversify/Server/Models"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/oauth2"
)

type AuthConfig struct {
	Config *oauth2.Config
	Token  models.AuthToken
}

type AuthService interface {
	HandleLogin(w http.ResponseWriter, r *http.Request)
	HandleCallback(w http.ResponseWriter, r *http.Request)
}

func (a *AuthConfig) HandleLogin(w http.ResponseWriter, r *http.Request) {
	url := a.Config.AuthCodeURL("random")
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func (a *AuthConfig) HandleCallback(w http.ResponseWriter, r *http.Request) {
	if r.FormValue("state") != "random" {
		log.Println("invalid oauth2 state")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	token, err := a.Config.Exchange(r.Context(), r.FormValue("code"))
	if err != nil {
		log.Fatalf("code exchange wrong: %s", err.Error())
	}

	a.Token.AccessToken = token.AccessToken
	fmt.Println("Auth Token: " + a.Token.AccessToken)
}
