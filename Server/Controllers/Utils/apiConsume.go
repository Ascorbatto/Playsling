package utils

const (
	//Spotify Endpoints
	SpotifyCurrentUser           = "https://api.spotify.com/v1/me"
	SpotifyCurrentUserPlaylists  = "https://api.spotify.com/v1/me/playlists"
	SpotifyPlaylistItemsInfo     = "https://api.spotify.com/v1/playlists/"
	SpotifySearch                = "https://api.spotify.com/v1/search?type=track&limit=50&q="
	SpotifyCreatePlaylist        = "https://api.spotify.com/v1/users/"
	SpotifyConvertPlaylistFromYT = "http://localhost:8888/convert_yt_to_s?destination="

	//Youtube Endpoints
	YoutubePlaylist         = "https://www.googleapis.com/youtube/v3/playlists?part=snippet&id="
	YoutubePlaylistItems    = "https://www.googleapis.com/youtube/v3/playlistItems?part=snippet,contentDetails&maxResults=50&playlistId="
	YoutubeVideoInfo        = "https://youtube.googleapis.com/youtube/v3/videos?part=snippet,contentDetails&id="
	YoutubeSearch           = "https://www.googleapis.com/youtube/v3/search?part=snippet&type=video&videoCategoryId=10&q="
	YoutubePostPlaylistItem = "https://www.googleapis.com/youtube/v3/playlistItems"

	//Host
	Host = "http://localhost:8888"
)
