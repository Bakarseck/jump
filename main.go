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
			if username != "" && email != "" {
				utils.UpdateEnvFile(username, email)
				utils.LoadEnv(models.HomeDir + "/.env")
				cli.ConfigureGit()
			}

			if usernameGithub != "" {
				utils.AddUsernameGithub(homeDir, usernameGithub)
				utils.LoadEnv(models.HomeDir + "/.env")
			}

			if saveCredentials {
				cli.SaveCredentials()
			}

			if GITHUB_TOKEN != "" {
				utils.LoadEnv(models.HomeDir + "/.env")
				utils.AddToken(homeDir, GITHUB_TOKEN)
			}

			if decrypt != "" {
				utils.LoadEnv(models.HomeDir + "/.env")
				cli.GetToken(decrypt)
			}
		},
	}

	rootCmd.Flags().StringVarP(&email, "email", "e", "", "Votre adresse email")
	rootCmd.Flags().StringVarP(&username, "username", "u", "", "Votre nom d'utilisateur")
	rootCmd.Flags().BoolVarP(&saveCredentials, "save", "s", false, "Sauvegarde les credentials dans un fichier")
	rootCmd.Flags().StringVarP(&usernameGithub, "ugithub", "g", "", "Username github")
	rootCmd.Flags().StringVarP(&GITHUB_TOKEN, "token", "t", "", "Save votre github")
	rootCmd.Flags().StringVarP(&decrypt, "decrypt", "d", "", "Decrypte une chaine")

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
	cmdClone.Flags().BoolVarP(&cli.Provider, "github", "p", false, "Provider Github")

	var commitCmd = &cobra.Command{
		Use:   "commit",
		Short: "Effectue un commit Git",
		Run:   cli.CommitRepo,
	}

	var executeScriptCmd = &cobra.Command{
		Use:   "zsh",
		Short: "Exécute un script shell spécifique",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cli.ExecuteShellScript()
		},
	}

	var addAlias = &cobra.Command{
		Use:   "alias",
		Short: "Ajoute des alias spécifiques",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cli.AddAlias()
		},
	}

	var createCmd = &cobra.Command{
		Use:   "create",
		Short: "Crée un nouveau dépôt sur GitHub",
		Args:  cobra.ExactArgs(1),
		Run:   cli.CreateRepo,
	}

	createCmd.Flags().BoolVarP(&cli.Private, "public", "p", false, "Private or Public")

	var addCollabCmd = &cobra.Command{
		Use:   "add-collab",
		Short: "Add a collaborator to a repository",
		Args:  cobra.ExactArgs(2),
		Run:   cli.AddCollab,
	}
	rootCmd.AddCommand(addCollabCmd)

	var changeVisibilityCmd = &cobra.Command{
		Use:   "change-visibility",
		Short: "Change the visibility of the repository",
		Args:  cobra.ExactArgs(1),
		Run:   cli.ChangeVisibility,
	}
	rootCmd.AddCommand(changeVisibilityCmd)

	var deleteRepoCmd = &cobra.Command{
		Use:   "delete-repo",
		Short: "Delete a repository",
		Args:  cobra.ExactArgs(1),
		Run:   cli.DeleteRepo,
	}
	rootCmd.AddCommand(deleteRepoCmd)

	// Définition des flags
	commitCmd.Flags().StringSliceVarP(&cli.Files, "files", "f", []string{}, "Fichiers à inclure dans le commit")
	commitCmd.Flags().StringVarP(&cli.Message, "message", "m", "Commit automatique", "Message de commit")

	rootCmd.AddCommand(cmdAdd)
	rootCmd.AddCommand(cmdJump)
	rootCmd.AddCommand(cmdClone)
	rootCmd.AddCommand(commitCmd)
	rootCmd.AddCommand(executeScriptCmd)
	rootCmd.AddCommand(addAlias)
	rootCmd.AddCommand(createCmd)

	// Exécute l'application
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var (
	username        string
	usernameGithub  string
	email           string
	GITHUB_TOKEN    string
	decrypt         string
	saveCredentials bool
)
