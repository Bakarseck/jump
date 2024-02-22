package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Bakarseck/jump/cmd/cli"
	"github.com/Bakarseck/jump/internals/models"
	"github.com/Bakarseck/jump/internals/utils"

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

	cmdClone.Flags().StringVarP(&cli.Collaborateur, "collaborateur", "c", "", "nom du collaborateur")

	var commitCmd = &cobra.Command{
		Use:   "commit",
		Short: "Effectue un commit Git",
		Run:   cli.CommitRepo,
	}

	// Définition des flags
	commitCmd.Flags().StringSliceVarP(&cli.Files, "files", "f", []string{}, "Fichiers à inclure dans le commit")
	commitCmd.Flags().StringVarP(&cli.Message, "message", "m", "Commit automatique", "Message de commit")

	// Ajout de commitCmd à rootCmd si vous avez un objet rootCmd représentant la commande racine de votre application
	rootCmd.AddCommand(commitCmd)

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
