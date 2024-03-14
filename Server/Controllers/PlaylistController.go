package controllers

import (
	services "Conversify/Server/Services"
	"net/http"
)

type PlaylistController struct {
	PlaylistService services.PlaylistService
}

func NewPlaylistController(playlistService services.PlaylistService) *PlaylistController {
	return &PlaylistController{PlaylistService: playlistService}
}

// Retrieves the playlist details (name, author, description...).
func (pc *PlaylistController) GetPlaylistInfo(w http.ResponseWriter, r *http.Request) {
	playlistId := r.URL.Query().Get("playlist")
	pc.PlaylistService.GetPlaylistInfo(w, r, playlistId)
}

// Get a list of the playlists owned or followed by the current Spotify user.
func (pc *PlaylistController) GetCurrentUserPlaylists(w http.ResponseWriter, r *http.Request) {
	pc.PlaylistService.GetCurrentUserPlaylists(w, r)
}

// Get a list of playlist items.
func (pc *PlaylistController) GetPlaylistItemsInfo(w http.ResponseWriter, r *http.Request) {
	playlistId := r.URL.Query().Get("playlist")
	pc.PlaylistService.GetPlaylistItemsInfo(w, r, playlistId)
}

// Creates a new playlist with the Name specified in the request.
// If already created, will return the items of said playlist.
func (pc *PlaylistController) CreatePlaylist(w http.ResponseWriter, r *http.Request) {
	playlistName := r.URL.Query().Get("name")
	pc.PlaylistService.CreatePlaylist(w, r, playlistName)
}
