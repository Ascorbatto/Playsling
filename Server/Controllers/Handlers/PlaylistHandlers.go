package handlers

import (
	controllers "Conversify/Server/Controllers"
	spotify "Conversify/Server/Controllers/Api/Spotify"
	youtube "Conversify/Server/Controllers/Api/Youtube"
	services "Conversify/Server/Services"
	spotifyservices "Conversify/Server/Services/Spotify"
	youtubeservices "Conversify/Server/Services/Youtube"
	"net/http"
	"strings"
)

var playlistControllers map[string]*controllers.PlaylistController

func InitPlaylistController(platform string, client *http.Client) {
	playlistControllers = make(map[string]*controllers.PlaylistController)
	playlistConfig := services.NewPlaylistConfig(client)

	switch platform {
	case "spotify":
		playlistService := spotifyservices.NewSpotifyPlaylistService(playlistConfig)
		playlistControllers["spotify"] = controllers.NewPlaylistController(playlistService)
	case "youtube":
		playlistService := youtubeservices.NewYoutubePlaylistService(playlistConfig)
		playlistControllers["youtube"] = controllers.NewPlaylistController(playlistService)

	}
}

func PlaylistInfoHandler(w http.ResponseWriter, r *http.Request) {
	platform := strings.TrimPrefix(r.URL.Path, "/playlist_info/")

	playlistController, ok := playlistControllers[platform]
	if !ok {
		http.Error(w, "Servicio no válido", http.StatusNotFound)
		return
	}

	playlistController.GetPlaylistInfo(w, r)
}

func UserPlaylistsHandler(w http.ResponseWriter, r *http.Request) {
	platform := strings.TrimPrefix(r.URL.Path, "/user_playlists/")

	switch platform {
	case "youtube":
		youtube.HandleLogin(w, r)
	case "spotify":
		spotify.GetCurrentUserPlaylists(w, r)
	default:
		http.Error(w, "Servicio no válido", http.StatusNotFound)
	}
}

func PlaylistHandler(w http.ResponseWriter, r *http.Request) {
	platform := strings.TrimPrefix(r.URL.Path, "/playlist/")

	switch platform {
	case "youtube":
		youtube.GetPlaylistItemsInfo(w, r)
	case "spotify":
		spotify.GetPlaylistItemsInfo(w, r)
	default:
		http.Error(w, "Servicio no válido", http.StatusNotFound)
	}
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
