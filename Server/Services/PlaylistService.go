package services

import (
	"net/http"
)

type PlaylistConfig struct {
	Client *http.Client
}

type PlaylistService interface {
	GetPlaylistInfo(w http.ResponseWriter, r *http.Request, playlistId string)
	GetPlaylistItemsInfo(w http.ResponseWriter, r *http.Request, playlistId string)
	GetCurrentUserPlaylists(w http.ResponseWriter, r *http.Request)
}

func (ps *PlaylistConfig) SetClient(client *http.Client) {
	ps.Client = client
}

func NewPlaylistConfig(client *http.Client) *PlaylistConfig {
	return &PlaylistConfig{Client: client}
}
