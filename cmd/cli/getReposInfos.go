package cli

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/Bakarseck/jump/internals/models"
	"github.com/Bakarseck/jump/internals/utils"
)

func GetRepoInfo(repo string) error {
	utils.LoadEnv(models.HomeDir + "/.env")
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s", os.Getenv("USERNAME_GITHUB"), repo)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	token, err := utils.DecryptString(os.Getenv("GITHUB_TOKEN"), os.Getenv("SECRET_KEY"))
	if err != nil {
		fmt.Println("Le token GitHub n'est pas défini dans le fichier .env")
		os.Exit(1)
	}
	req.Header.Set("Authorization", "token "+token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("HTTP error: %s", resp.Status)
	}

	var repoInfo Repository
	if err := json.NewDecoder(resp.Body).Decode(&repoInfo); err != nil {
		return err
	}

	fmt.Printf("Repository Info: %+v\n", repoInfo)
	return nil
}
