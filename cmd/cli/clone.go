package cli

import (
	"fmt"
	"jump/internals/models"
	"jump/internals/utils"
	"log"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

func ConfigureGit() {
	utils.LoadEnv(models.HomeDir + "/.env")
	username := os.Getenv("USERNAME")
	email := os.Getenv("EMAIL")

	if username == "" || email == "" {
		log.Println("Le nom d'utilisateur et l'email doivent être définis dans le fichier .env")
		return
	}

	// Configurer le nom d'utilisateur et l'email pour Git
	utils.ExecCommand("git", "config", "--global", "user.name", username)
	utils.ExecCommand("git", "config", "--global", "user.email", email)

	fmt.Println("Configuration Git mise à jour avec succès.")
}

func CloneRepo(cmd *cobra.Command, args []string) {
	utils.LoadEnv(models.HomeDir + "/.env")

	name := args[0]

	username := os.Getenv("USERNAME")

	url := fmt.Sprintf("https://learn.zone01dakar.sn/git/%v/%v", username, name)

	if err := GitClone(url); err != nil {
		log.Fatalf("Erreur lors du clonage du dépôt : %v", err)
	}
	fmt.Println("Dépôt cloné avec succès !")
}

func GitClone(url string) error {
	cmd := exec.Command("git", "clone", url)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
