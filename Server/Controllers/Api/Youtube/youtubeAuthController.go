package youtube

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
	"golang.org/x/oauth2/google"
)

func init() {
	configDir, _ := os.Getwd()

	dotenvPath := filepath.Join(configDir, "Server", "Config", ".env")

	err := godotenv.Load(dotenvPath)
	utils.ErrorManager(utils.LoadDotenvError, err)

	googleOauthConfig = &oauth2.Config{
		RedirectURL:  os.Getenv("YOUTUBE_REDIRECTURI"),
		ClientID:     os.Getenv("YOUTUBE_CLIENT_ID"),
		ClientSecret: os.Getenv("YOUTUBE_CLIENT_SECRET"),
		Scopes:       []string{"https://www.googleapis.com/auth/youtube"},
		Endpoint:     google.Endpoint,
	}
}

var (
	Token             models.AuthToken
	googleOauthConfig *oauth2.Config
)

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	state := googleOauthConfig.AuthCodeURL("random")
	http.Redirect(w, r, state, http.StatusTemporaryRedirect)
}

func RetrieveToken(w http.ResponseWriter, r *http.Request) {

	if r.FormValue("state") != "random" {
		log.Println("invalid oauth2 state")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	token, err := googleOauthConfig.Exchange(r.Context(), r.FormValue("code"))
	if err != nil {
		log.Fatalf("code exchange wrong: %s", err.Error())
	}

	Token.AccessToken = token.AccessToken
	fmt.Println("Youtube Auth: " + Token.AccessToken)
}
