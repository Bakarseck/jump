package cli

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/Bakarseck/jump/internals/utils"
)

func ExecuteShellScript() error {

	homeDir, err := os.UserHomeDir()

	githubScriptUrl := "https://raw.githubusercontent.com/Bakarseck/jump/master/install.sh"

	if err != nil {
		return fmt.Errorf("impossible de trouver le répertoire personnel: %v", err)
	}

	scriptPath, err := utils.DownloadFile(githubScriptUrl, homeDir)

	if err != nil {
		return fmt.Errorf("erreur lors du téléchargement du script: %w", err)
	}

	// Rend le script exécutable.
	err = os.Chmod(scriptPath, 0755)
	if err != nil {
		return fmt.Errorf("erreur lors du changement des permissions du fichier: %w", err)
	}

	cmd := exec.Command("bash", scriptPath)
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("erreur lors de l'exécution du script shell: %w", err)
	}
	fmt.Println("Script exécuté avec succès.")
	return nil
}
