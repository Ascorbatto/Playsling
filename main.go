package main

import (
	youtube "Conversify/Server/Controllers/Api/Youtube"
	handlers "Conversify/Server/Controllers/Handlers"
	"net/http"
)

func main() {

	http.HandleFunc("/login/", handlers.LoginHandler)
	http.HandleFunc("/callback/", handlers.CallbackHandler)
	http.HandleFunc("/user_playlists/", handlers.UserPlaylistsHandler)
	http.HandleFunc("/playlist/", handlers.PlaylistHandler)
	http.HandleFunc("/create/", handlers.CreatePlaylist)

	http.HandleFunc("/post_yt", youtube.TestPostPlaylistItem)
	http.HandleFunc("/convert_yt_to_s", youtube.ToSpotify)
	http.ListenAndServe(":8888", nil)
}
