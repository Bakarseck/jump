package main

import (
	"fmt"
	"jump/cmd/cli"
	"jump/internals/models"
	"jump/internals/utils"
	"log"
	"os"

	"github.com/spf13/cobra"
)

func main() {

	homeDir, err := os.UserHomeDir()
	models.HomeDir = homeDir

	if err != nil {
		log.Fatalf("Erreur lors de la récupération du répertoire personnel : %v", err)
	}

	models.PathJson = homeDir + "/jump.json"

	var rootCmd = &cobra.Command{
		Use:   "jump",
		Short: "Un gestionnaire de répertoires flexible et rapide",
		Long: "\n" + `
		Jump est un gestionnaire de répertoires construit en Go.
		Il permet d'ajouter et de naviguer facilement entre différents répertoires.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Cette logique s'exécute avant chaque commande
			if username != "" && email != "" {
				utils.UpdateEnvFile(username, email)
				utils.LoadEnv(models.HomeDir + "/.env")
				cli.ConfigureGit()
			}
		},
	}

	rootCmd.Flags().StringVarP(&email, "email", "e", "", "Votre adresse email")
	rootCmd.Flags().StringVarP(&username, "username", "u", "", "Votre nom d'utilisateur")

	// Définit une commande pour ajouter un répertoire
	var cmdAdd = &cobra.Command{
		Use:   "add",
		Short: "Ajouter un répertoire",
		Long: "\n" + `
		Ajoute un nouveau répertoire à la liste de gestion.
		Vous pouvez spécifier un nom et un chemin, ou utiliser '.' pour ajouter le répertoire courant.
		Exemple: 'jump add projet /chemin/vers/projet' ou 'jump add .'
		Pour pouvoir utiliser le 'jump add .', il faut naviguer jusqu'au répértoire que tu veux ajouter.`,

		Run: cli.Add,
	}

	// Définit une commande pour changer de répertoire
	var cmdJump = &cobra.Command{
		Use:   "to",
		Short: "Jump vers un répertoire",
		Long: "\n" + `
		Change le répertoire courant en celui spécifié.
		Utilise le nom du répertoire tel qu'ajouté avec la commande 'add'.
		Exemple: 'jump to projet' pour aller au répertoire nommé 'projet'.`,
		Run: cli.To,
	}

	var cmdClone = &cobra.Command{
		Use:   "clone [url]",
		Short: "Clone un dépôt et configure .env",
		Long: `Clone un dépôt Git et ajoute/actualise les variables d'utilisateur dans un fichier .env à la racine.
		Exemple: 'jump clone https://example.com/repo.git -u monNom -e monEmail@example.com'`,
		Run: cli.CloneRepo,
	}

	// Attache les commandes à l'application principale
	rootCmd.AddCommand(cmdAdd)
	rootCmd.AddCommand(cmdJump)
	rootCmd.AddCommand(cmdClone)

	// Exécute l'application
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var (
	username string
	email    string
)
