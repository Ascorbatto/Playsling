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

func (pc *PlaylistController) GetCurrentUserPlaylists(w http.ResponseWriter, r *http.Request) {
	pc.PlaylistService.GetCurrentUserPlaylists(w, r)
}

func (pc *PlaylistController) GetPlaylistItemsInfo(w http.ResponseWriter, r *http.Request) {
	playlistId := r.URL.Query().Get("playlist")
	pc.PlaylistService.GetPlaylistItemsInfo(w, r, playlistId)
}
