package controllers

import (
	services "Conversify/Server/Services"
	"fmt"
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
	PlaylistInfo := pc.PlaylistService.GetPlaylistInfo(w, r, playlistId)
	fmt.Println(PlaylistInfo)
}
