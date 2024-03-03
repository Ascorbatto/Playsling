package handlers

import (
	controllers "Conversify/Server/Controllers"
	services "Conversify/Server/Services"
	spotifyservices "Conversify/Server/Services/Spotify"
	youtubeservices "Conversify/Server/Services/Youtube"
	"net/http"
	"strings"
)

var userControllers = make(map[string]*controllers.UserController)

func InitUserController(platform string, client *http.Client) {
	userConfig := services.NewUserConfig(client)
	var userService services.UserService

	switch platform {
	case "spotify":
		userService = spotifyservices.NewSpotifyUserService(userConfig)
	case "youtube":
		userService = youtubeservices.NewYoutubeUserService(userConfig)
	}

	userControllers[platform] = controllers.NewUserController(userService)
}

func CurrentUserHandler(w http.ResponseWriter, r *http.Request) {
	platform := strings.TrimPrefix(r.URL.Path, "/me/")

	userController, ok := userControllers[platform]
	if !ok {
		http.Error(w, "Service not available", http.StatusNotFound)
		return
	}

	userController.GetCurrentUser(w, r)
}
