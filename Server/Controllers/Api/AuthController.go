package api

import (
	services "Conversify/Server/Services"
	"net/http"
)

/*func init() {
	configDir, _ := os.Getwd()
	fmt.Println(configDir)

	dotenvPath := filepath.Join(configDir, "Server", "Config", ".env")

	err := godotenv.Load(dotenvPath)
	utils.ErrorManager(utils.LoadDotenvError, err)
	fmt.Println(os.Getenv("SPOTIFY_CLIENT_ID") + " " + os.Getenv("SPOTIFY_CLIENT_SECRET") + " " + os.Getenv("SPOTIFY_REDIRECTURI"))
}*/

type AuthController struct {
	AuthService services.AuthService
}

func (ac *AuthController) HandleLogin(w http.ResponseWriter, r *http.Request) {
	ac.AuthService.HandleLogin(w, r)
}

func (ac *AuthController) HandleCallback(w http.ResponseWriter, r *http.Request) {
	ac.AuthService.HandleCallback(w, r)
}
