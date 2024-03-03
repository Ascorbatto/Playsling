package handlers

import (
	controllers "Conversify/Server/Controllers"
	spotify "Conversify/Server/Controllers/Api/Spotify"
	services "Conversify/Server/Services"
	spotifyservices "Conversify/Server/Services/Spotify"
	youtubeservices "Conversify/Server/Services/Youtube"
	"net/http"
	"strings"
)

var playlistControllers = make(map[string]*controllers.PlaylistController)

func InitPlaylistController(platform string, client *http.Client) {
	playlistConfig := services.NewPlaylistConfig(client)
	var playlistService services.PlaylistService

	switch platform {
	case "spotify":
		playlistService = spotifyservices.NewSpotifyPlaylistService(playlistConfig)
	case "youtube":
		playlistService = youtubeservices.NewYoutubePlaylistService(playlistConfig)
	}

	playlistControllers[platform] = controllers.NewPlaylistController(playlistService)
}

func PlaylistInfoHandler(w http.ResponseWriter, r *http.Request) {
	platform := strings.TrimPrefix(r.URL.Path, "/playlist-info/")

	playlistController, ok := playlistControllers[platform]
	if !ok {
		http.Error(w, "Servicio no válido", http.StatusNotFound)
		return
	}

	playlistController.GetPlaylistInfo(w, r)
}

func CurrentUserPlaylistsHandler(w http.ResponseWriter, r *http.Request) {
	platform := strings.TrimPrefix(r.URL.Path, "/user-playlists/")

	playlistController, ok := playlistControllers[platform]
	if !ok {
		http.Error(w, "Servicio no válido", http.StatusNotFound)
		return
	}

	playlistController.GetCurrentUserPlaylists(w, r)
}

func PlaylistItemsHandler(w http.ResponseWriter, r *http.Request) {
	platform := strings.TrimPrefix(r.URL.Path, "/playlist/")

	playlistController, ok := playlistControllers[platform]
	if !ok {
		http.Error(w, "Service not available", http.StatusNotFound)
		return
	}

	playlistController.GetPlaylistItemsInfo(w, r)
}

func CreatePlaylist(w http.ResponseWriter, r *http.Request) {
	platform := strings.TrimPrefix(r.URL.Path, "/create/")

	switch platform {
	case "youtube":

	case "spotify":
		spotify.CreatePlaylist(w, r)
	default:
		http.Error(w, "Servicio no válido", http.StatusNotFound)
	}
}
