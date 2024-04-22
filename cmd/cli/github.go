package cli

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/Bakarseck/jump/internals/models"
	"github.com/Bakarseck/jump/internals/utils"
	"github.com/spf13/cobra"
)

var (
	Private = false
)

type Repository struct {
	Name        string `json:"name"`
	Private     bool   `json:"private"`
	Description string `json:"description,omitempty"`
}

type CollaboratorOptions struct {
	Permissions string `json:"permissions"`
}

func CreateRepo(cmd *cobra.Command, args []string) {
	utils.LoadEnv(models.HomeDir + "/.env")

	token, err := utils.DecryptString(os.Getenv("GITHUB_TOKEN"), os.Getenv("SECRET_KEY"))
	if err != nil {
		fmt.Println("Le token GitHub n'est pas défini dans le fichier .env")
		os.Exit(1)
	}

	repo := Repository{
		Name:    args[0],
		Private: !Private,
	}

	repoBytes, err := json.Marshal(repo)
	if err != nil {
		fmt.Println("Erreur lors de la sérialisation du dépôt:", err)
		os.Exit(1)
	}

	req, err := http.NewRequest("POST", "https://api.github.com/user/repos", bytes.NewBuffer(repoBytes))
	if err != nil {
		fmt.Println("Erreur lors de la création de la requête:", err)
		os.Exit(1)
	}

	req.Header.Set("Authorization", "token "+token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Erreur lors de l'envoi de la requête:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		fmt.Println("Erreur lors de la création du dépôt:", resp.Status)
		os.Exit(1)
	}

	fmt.Println("Dépôt créé avec succès !")
}

// AddCollab adds a collaborator to a GitHub repository.
func AddCollab(cmd *cobra.Command, args []string) {
	utils.LoadEnv(models.HomeDir + "/.env")

	token, err := utils.DecryptString(os.Getenv("GITHUB_TOKEN"), os.Getenv("SECRET_KEY"))
	if err != nil {
		fmt.Println("Error decrypting GitHub token:", err)
		os.Exit(1)
	}

	repoName, collaborator := args[0], args[1]
	permissions := CollaboratorOptions{
		Permissions: "push",
	}

	data, err := json.Marshal(permissions)
	if err != nil {
		fmt.Println("Error marshalling permissions:", err)
		os.Exit(1)
	}

	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/collaborators/%s", os.Getenv("USERNAME_GITHUB"), repoName, collaborator)
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(data))
	if err != nil {
		fmt.Println("Error creating request:", err)
		os.Exit(1)
	}

	req.Header.Set("Authorization", "token "+token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		fmt.Println("Failed to add collaborator:", resp.Status)
		os.Exit(1)
	}

	fmt.Println("Collaborator added successfully!")
}

// ChangeVisibility changes the visibility of a GitHub repository.
func ChangeVisibility(cmd *cobra.Command, args []string) {
	utils.LoadEnv(models.HomeDir + "/.env")

	token, err := utils.DecryptString(os.Getenv("GITHUB_TOKEN"), os.Getenv("SECRET_KEY"))
	if err != nil {
		fmt.Println("Error decrypting GitHub token:", err)
		os.Exit(1)
	}

	repoName := args[0]
	username := os.Getenv("USERNAME_GITHUB")
	if username == "" {
		fmt.Println("GitHub username is not set.")
		os.Exit(1)
	}

	// Step 1: Get current repository details
	getURL := fmt.Sprintf("https://api.github.com/repos/%s/%s", username, repoName)
	getRequest, err := http.NewRequest("GET", getURL, nil)
	if err != nil {
		fmt.Println("Error creating request for getting repo details:", err)
		os.Exit(1)
	}
	getRequest.Header.Set("Authorization", "token "+token)

	getClient := &http.Client{}
	getResp, err := getClient.Do(getRequest)
	if err != nil {
		fmt.Println("Error sending request for getting repo details:", err)
		os.Exit(1)
	}
	defer getResp.Body.Close()

	if getResp.StatusCode != http.StatusOK {
		fmt.Println("Failed to get repository details:", getResp.Status)
		os.Exit(1)
	}

	var repoData Repository
	if err := json.NewDecoder(getResp.Body).Decode(&repoData); err != nil {
		fmt.Println("Error decoding repository data:", err)
		os.Exit(1)
	}

	// Step 2: Change the visibility
	newVisibility := !repoData.Private // Toggle visibility

	repoData.Private = newVisibility

	patchData, err := json.Marshal(repoData)
	if err != nil {
		fmt.Println("Error marshalling repository data for update:", err)
		os.Exit(1)
	}

	patchURL := fmt.Sprintf("https://api.github.com/repos/%s/%s", username, repoName)
	patchRequest, err := http.NewRequest("PATCH", patchURL, bytes.NewBuffer(patchData))
	if err != nil {
		fmt.Println("Error creating patch request:", err)
		os.Exit(1)
	}

	patchRequest.Header.Set("Authorization", "token "+token)
	patchRequest.Header.Set("Content-Type", "application/json")

	patchResp, err := http.DefaultClient.Do(patchRequest)
	if err != nil {
		fmt.Println("Error sending patch request:", err)
		os.Exit(1)
	}
	defer patchResp.Body.Close()

	if patchResp.StatusCode != http.StatusOK {
		fmt.Println("Failed to change repository visibility:", patchResp.Status)
		os.Exit(1)
	}

	fmt.Println("Repository visibility changed successfully!")
}

// DeleteRepo deletes a GitHub repository.
func DeleteRepo(cmd *cobra.Command, args []string) {
	utils.LoadEnv(models.HomeDir + "/.env")

	token, err := utils.DecryptString(os.Getenv("GITHUB_TOKEN"), os.Getenv("SECRET_KEY"))
	if err != nil {
		fmt.Println("Error decrypting GitHub token:", err)
		os.Exit(1)
	}

	repoName := args[0]
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s", os.Getenv("USERNAME_GITHUB"), repoName)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		os.Exit(1)
	}

	req.Header.Set("Authorization", "token "+token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		fmt.Println("Failed to delete repository:", resp.Status)
		os.Exit(1)
	}

	fmt.Println("Repository deleted successfully!")
}
