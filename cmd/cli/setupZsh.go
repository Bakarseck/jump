package cli

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/Bakarseck/jump/internals/utils"
)

func ExecuteShellScript() error {

	homeDir, err := os.UserHomeDir()

	if err != nil {
		return fmt.Errorf("impossible de trouver le répertoire personnel: %v", err)
	}

	scriptPath := homeDir + "/install.sh"

	fmt.Println(scriptPath)

	utils.WriteFile(scriptPath, "alias.sh")

	cmd := exec.Command("bash", scriptPath)
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("erreur lors de l'exécution du script shell: %w", err)
	}
	fmt.Println("Script exécuté avec succès.")
	return nil
}
