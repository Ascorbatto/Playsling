package handlers

import (
	api "Conversify/Server/Controllers/Api"
	utils "Conversify/Server/Controllers/Utils"
	models "Conversify/Server/Models"
	services "Conversify/Server/Services"
	spotifyservices "Conversify/Server/Services/Spotify"
	youtubeservices "Conversify/Server/Services/Youtube"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	youtubeAuthController *api.AuthController
	spotifyAuthController *api.AuthController
)

func init() {
	configDir, _ := os.Getwd()

	dotenvPath := filepath.Join(configDir, "Server", "Config", ".env")

	err := godotenv.Load(dotenvPath)
	utils.ErrorManager(utils.LoadDotenvError, err)

	googleOauthConfig := &services.AuthConfig{
		Config: &oauth2.Config{
			RedirectURL:  os.Getenv("YOUTUBE_REDIRECTURI"),
			ClientID:     os.Getenv("YOUTUBE_CLIENT_ID"),
			ClientSecret: os.Getenv("YOUTUBE_CLIENT_SECRET"),
			Scopes:       []string{"https://www.googleapis.com/auth/youtube"},
			Endpoint:     google.Endpoint,
		},
		Token: models.AuthToken{},
	}
	youtubeAuthService := &youtubeservices.YoutubeAuthService{AuthConfig: *googleOauthConfig}
	youtubeAuthController = &api.AuthController{AuthService: youtubeAuthService}

	spotifyOauthConfig := &services.AuthConfig{
		Config: &oauth2.Config{
			RedirectURL:  os.Getenv("SPOTIFY_REDIRECTURI"),
			ClientID:     os.Getenv("SPOTIFY_CLIENT_ID"),
			ClientSecret: os.Getenv("SPOTIFY_CLIENT_SECRET"),
			Scopes:       []string{"user-read-private user-read-email playlist-modify-private playlist-modify-public user-library-read user-library-modify"},
			Endpoint: oauth2.Endpoint{
				AuthURL:  utils.SpotifyAuth,
				TokenURL: utils.SpotifyToken,
			},
		},
		Token: models.AuthToken{},
	}
	spotifyAuthService := &spotifyservices.SpotifyAuthService{AuthConfig: *spotifyOauthConfig}
	spotifyAuthController = &api.AuthController{AuthService: spotifyAuthService}
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	platform := strings.TrimPrefix(r.URL.Path, "/login/")

	switch platform {
	case "youtube":
		youtubeAuthController.HandleLogin(w, r)
	case "spotify":
		spotifyAuthController.HandleLogin(w, r)
	default:
		http.Error(w, "Servicio no válido", http.StatusNotFound)
	}
}
func CallbackHandler(w http.ResponseWriter, r *http.Request) {
	platform := strings.TrimPrefix(r.URL.Path, "/callback/")
	switch platform {
	case "youtube":
		youtubeAuthController.HandleCallback(w, r)
	case "spotify":
		spotifyAuthController.HandleCallback(w, r)
	default:
		http.Error(w, "Servicio no válido", http.StatusNotFound)
	}
}
