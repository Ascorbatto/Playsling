package handlers

import (
	controllers "Conversify/Server/Controllers"
	services "Conversify/Server/Services"
	spotifyservices "Conversify/Server/Services/Spotify"
	youtubeservices "Conversify/Server/Services/Youtube"
	"net/http"
	"strings"

	"golang.org/x/oauth2"
)

var authControllers map[string]*controllers.AuthController

func InitAuthControllers(oauthConfigs map[string]*oauth2.Config) {
	authControllers = make(map[string]*controllers.AuthController)

	youtubeAuthConfig := services.NewAuthConfig(oauthConfigs["youtube"])
	youtubeAuthService := youtubeservices.NewYoutubeAuthService(youtubeAuthConfig)
	authControllers["youtube"] = controllers.NewAuthController(youtubeAuthService)

	spotifyAuthConfig := services.NewAuthConfig(oauthConfigs["spotify"])
	spotifyAuthService := spotifyservices.NewSpotifyAuthService(spotifyAuthConfig)
	authControllers["spotify"] = controllers.NewAuthController(spotifyAuthService)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	platform := strings.TrimPrefix(r.URL.Path, "/login/")
	authController, ok := authControllers[platform]
	if !ok {
		http.Error(w, "Servicio no válido", http.StatusNotFound)
		return
	}

	authController.HandleLogin(w, r)
}

func CallbackHandler(w http.ResponseWriter, r *http.Request) {
	platform := strings.TrimPrefix(r.URL.Path, "/callback/")
	authController, ok := authControllers[platform]
	if !ok {
		http.Error(w, "Servicio no válido", http.StatusNotFound)
		return
	}

	client := *authController.HandleCallback(w, r)
	InitControllers(platform, &client)
}

func InitControllers(platform string, client *http.Client) {
	InitPlaylistController(platform, client)
}
