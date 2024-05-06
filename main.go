package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/Bakarseck/jump/cmd/cli"
	"github.com/Bakarseck/jump/internals/models"
	"github.com/Bakarseck/jump/internals/utils"

	"github.com/spf13/cobra"
)

var usernames = []string{
	"Mirak29",
	"mohaskii",
	"AbubakrSaadiq",
	"NiangOos",
	"Nixa001",
	"superMass14",
	"steb1",
	"Aziz-TheDarkCode",
	"elhadjiibrahima",
	"OumarLAM",
	"dioufra",
	"mouhasup",
	"pro12x",
	"biramendoye",
	"ahbarry07",
	"Whoisozdem",
	"seydi-ahmed",
	"djiby26",
	"SSMM0498",
	"papa-abdoulaye-diop",
	"jeebrail",
	"Bakarseck",
	"moussadiengsala",
	"Tafouiny",
	"aadieng100",
	"yayediop2",
	"alpapie",
	"vincefelix",
	"munikmind",
	"bayerane",
	"ElieJnr",
	"louisisaacdiouf",
	"lamabalde",
	"devousmane",
	"coulou800",
	"Yoows",
	"Madike10",
	"galsen-boy",
	"serignefallou-18",
	"lino-smart",
	"Cheibany",
	"Badoulahi8",
}

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
				utils.AddEnvFile("USERNAME", username)
				utils.AddEnvFile("EMAIL", email)
				utils.LoadEnv(models.HomeDir + "/.env")
				cli.ConfigureGit()
			}

			if usernameGithub != "" {
				utils.AddEnvFile("USERNAME_GITHUB", usernameGithub)
				utils.LoadEnv(models.HomeDir + "/.env")
			}

			if saveCredentials {
				cli.SaveCredentials()
			}

			if GITHUB_TOKEN != "" {
				utils.AddToken(homeDir, GITHUB_TOKEN)
			}

			if generateKey {
				secretKey := utils.GenerateSecretKey()
				utils.AddEnvFile("SECRET_KEY", secretKey)
				fmt.Println("La clé secrète a été générée et enregistrée avec succès.")
			}
		},
	}

	rootCmd.Flags().StringVarP(&email, "email", "e", "", "Votre adresse email")
	rootCmd.Flags().StringVarP(&username, "username", "u", "", "Votre nom d'utilisateur")
	rootCmd.Flags().BoolVarP(&saveCredentials, "save", "s", false, "Sauvegarde les credentials dans un fichier")
	rootCmd.Flags().StringVarP(&usernameGithub, "ugithub", "g", "", "Username github")
	rootCmd.Flags().StringVarP(&GITHUB_TOKEN, "token", "t", "", "Save votre github")
	rootCmd.PersistentFlags().BoolVarP(&generateKey, "generate-key", "k", false, "Génère et enregistre une nouvelle SECRET_KEY dans le fichier .env")

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

	var followCmd = &cobra.Command{
		Use:   "follow [username]",
		Short: "Follow a GitHub user",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			username := args[0]
			if err := cli.FollowUser(username); err != nil {
				fmt.Printf("Failed to follow user %s: %v\n", username, err)
				os.Exit(1)
			}
			fmt.Println("Successfully followed user:", username)
		},
	}

	var followZone = &cobra.Command{
		Use:   "followZone",
		Short: "Follow user who send theirs username github",
		Run: func(cmd *cobra.Command, args []string) {
			for _, v := range usernames {
				if err := cli.FollowUser(v); err != nil {
					fmt.Printf("Failed to follow user %s: %v\n", username, err)
					continue
				}
				fmt.Println("Successfully followed user:", v)
			}
		},
	}

	rootCmd.AddCommand(followCmd)
	rootCmd.AddCommand(followZone)

	var getRepoInfoCmd = &cobra.Command{
		Use:   "get-repo-info [repo]",
		Short: "Get information about a specific GitHub repository",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			repo := args[0]
			if err := cli.GetRepoInfo(repo); err != nil {
				fmt.Printf("Failed to get info for repository %s: %v\n", repo, err)
				os.Exit(1)
			}
		},
	}

	rootCmd.AddCommand(getRepoInfoCmd)

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

	var examCmd = &cobra.Command{
		Use:   "exercice",
		Short: "Creates a new Rust library and sets up the environment",
		Long: `This command creates a new Rust library project using 'cargo new --lib',
			   downloads a README.md, and performs additional setup.`,
		Args: cobra.ExactArgs(1), // Ensures exactly one argument is passed
		Run: func(cmd *cobra.Command, args []string) {
			libName := args[0]
			err := cli.CreateNewRustLib(libName)
			if err != nil {
				fmt.Printf("Error creating Rust library: %v\n", err)
				return
			}

			readmeURL := fmt.Sprintf("https://raw.githubusercontent.com/01-edu/public/master/subjects/%s/README.md", libName)
			readmePath := filepath.Join(libName, "README.md")

			err = cli.DownloadFile(readmeURL, readmePath)
			if err != nil {
				fmt.Printf("Error downloading README.md: %v\n", err)
				return
			}

			fmt.Println("New Rust library created successfully with README.md downloaded.")
		},
	}

	var testCmd = &cobra.Command{
		Use:   "test [projectName]",
		Short: "Tests a Rust project",
		Long: `This command navigates to the Rust project directory, downloads the test file,
			   and runs 'cargo test' for the project.`,
		Args: cobra.ExactArgs(1), // Ensures exactly one argument is passed
		Run: func(cmd *cobra.Command, args []string) {
			projectName := args[0]
			projectPath := filepath.Join(projectName, "src") // Assuming the src directory is at a specific path

			// Change to the project directory
			err := os.Chdir(projectPath)
			if err != nil {
				fmt.Printf("Error changing directory to %s: %v\n", projectPath, err)
				return
			}

			// Define the URL for the test main.rs file
			testFileURL := fmt.Sprintf("https://raw.githubusercontent.com/rgilles42/piscine-rust/main/tests/%s_test/src/main.rs", projectName)
			testFilePath := filepath.Join("main.rs") // Path where the test file should be saved

			// Download the test main.rs file
			err = cli.DownloadFile(testFileURL, testFilePath)
			if err != nil {
				fmt.Printf("Error downloading test file: %v\n", err)
				return
			}

			// Run cargo test in the current directory
			_cmd := exec.Command("cargo", "test")
			_cmd.Stdout = os.Stdout
			_cmd.Stderr = os.Stderr
			err = _cmd.Run()
			if err != nil {
				fmt.Printf("Error running cargo test: %v\n", err)
				return
			}

			fmt.Println("Tests executed successfully.")
		},
	}

	// Définition des flags
	commitCmd.Flags().StringSliceVarP(&cli.Files, "files", "f", []string{}, "Fichiers à inclure dans le commit")
	commitCmd.Flags().StringVarP(&cli.Message, "message", "m", "Commit automatique", "Message de commit")

	rootCmd.AddCommand(testCmd)
	rootCmd.AddCommand(cmdAdd)
	rootCmd.AddCommand(cmdJump)
	rootCmd.AddCommand(cmdClone)
	rootCmd.AddCommand(commitCmd)
	rootCmd.AddCommand(executeScriptCmd)
	rootCmd.AddCommand(addAlias)
	rootCmd.AddCommand(createCmd)
	rootCmd.AddCommand(examCmd)

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
	saveCredentials bool
	generateKey     bool
)
