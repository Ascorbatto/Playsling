package main

import (
	youtube "Conversify/Server/Controllers/Api/Youtube"
	handlers "Conversify/Server/Controllers/Handlers"
	utils "Conversify/Server/Controllers/Utils"
	"fmt"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/spotify"
)

func main() {
	utils.LoadDotEnv()

	oauthConfigs := make(map[string]*oauth2.Config)

	youtubeOauthConfig := &oauth2.Config{
		RedirectURL:  os.Getenv("YOUTUBE_REDIRECTURI"),
		ClientID:     os.Getenv("YOUTUBE_CLIENT_ID"),
		ClientSecret: os.Getenv("YOUTUBE_CLIENT_SECRET"),
		Scopes:       []string{"https://www.googleapis.com/auth/youtube"},
		Endpoint:     google.Endpoint,
	}
	oauthConfigs["youtube"] = youtubeOauthConfig

	spotifyOauthConfig := &oauth2.Config{
		RedirectURL:  os.Getenv("SPOTIFY_REDIRECTURI"),
		ClientID:     os.Getenv("SPOTIFY_CLIENT_ID"),
		ClientSecret: os.Getenv("SPOTIFY_CLIENT_SECRET"),
		Scopes:       []string{"user-read-private user-read-email playlist-modify-private playlist-modify-public user-library-read user-library-modify"},
		Endpoint:     spotify.Endpoint,
	}
	oauthConfigs["spotify"] = spotifyOauthConfig

	handlers.InitAuthControllers(oauthConfigs)

	http.HandleFunc("/login/", handlers.LoginHandler)
	http.HandleFunc("/callback/", handlers.CallbackHandler)
	http.HandleFunc("/playlist-info/", handlers.PlaylistInfoHandler)
	http.HandleFunc("/user-playlists/", handlers.CurrentUserPlaylistsHandler)
	http.HandleFunc("/playlist/", handlers.PlaylistItemsHandler)
	http.HandleFunc("/create-playlist/", handlers.CreatePlaylist)
	http.HandleFunc("/me/", handlers.CurrentUserHandler)

	http.HandleFunc("/post_yt", youtube.TestPostPlaylistItem)
	http.HandleFunc("/convert_yt_to_s", youtube.ToSpotify)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello")
	})
	http.ListenAndServe(":8888", nil)
}
