package cli

import (
	"fmt"
	"os/exec"
)

func ExecuteShellScript() error {
	cmd := exec.Command("bash", "install.sh")
	err := cmd.Run() // Utilisez cmd.Output() si vous souhaitez capturer la sortie
	if err != nil {
		return fmt.Errorf("erreur lors de l'exécution du script shell: %w", err)
	}
	fmt.Println("Script exécuté avec succès.")
	return nil
}