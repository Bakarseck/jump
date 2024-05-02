package cli

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Bakarseck/jump/internals/models"
	"github.com/Bakarseck/jump/internals/utils"
)

func FollowUser(username string) error {
	utils.LoadEnv(models.HomeDir + "/.env")
	url := fmt.Sprintf("https://api.github.com/user/following/%s", username)
	req, err := http.NewRequest("PUT", url, nil)
	if err != nil {
		return err
	}

	token, err := utils.DecryptString(os.Getenv("GITHUB_TOKEN"), os.Getenv("SECRET_KEY"))
	if err != nil {
		fmt.Println("Le token GitHub n'est pas défini dans le fichier .env")
		os.Exit(1)
	}
	req.Header.Set("Authorization", "token "+token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 204 {
		return fmt.Errorf("HTTP error: %s", resp.Status)
	}

	return nil
}
