package spotify

import (
	utils "Conversify/Server/Controllers/Utils"
	models "Conversify/Server/Models"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
)

func init() {
	configDir, _ := os.Getwd()

	dotenvPath := filepath.Join(configDir, "Server", "Config", ".env")

	err := godotenv.Load(dotenvPath)
	utils.ErrorManager(utils.LoadDotenvError, err)

	spotifyOauthConfig = &oauth2.Config{
		RedirectURL:  os.Getenv("SPOTIFY_REDIRECTURI"),
		ClientID:     os.Getenv("SPOTIFY_CLIENT_ID"),
		ClientSecret: os.Getenv("SPOTIFY_CLIENT_SECRET"),
		Scopes:       []string{"user-read-private user-read-email playlist-modify-private playlist-modify-public user-library-read user-library-modify"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  utils.SpotifyAuth,
			TokenURL: utils.SpotifyToken,
		},
	}
}

var (
	Token              models.AuthToken
	spotifyOauthConfig *oauth2.Config
)

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	state := spotifyOauthConfig.AuthCodeURL("random")

	http.Redirect(w, r, state, http.StatusTemporaryRedirect)
}

func RetrieveToken(w http.ResponseWriter, r *http.Request) {

	if r.FormValue("state") != "random" {
		log.Println("invalid oauth2 state")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	token, err := spotifyOauthConfig.Exchange(r.Context(), r.FormValue("code"))
	if err != nil {
		log.Fatalf("code exchange wrong: %s", err.Error())
	}

	Token.AccessToken = token.AccessToken
	fmt.Println("Spotify Auth: " + Token.AccessToken)
}
