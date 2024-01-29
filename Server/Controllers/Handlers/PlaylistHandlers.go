package handlers

import (
	spotify "Conversify/Server/Controllers/Api/Spotify"
	youtube "Conversify/Server/Controllers/Api/Youtube"
	"net/http"
	"strings"
)

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
