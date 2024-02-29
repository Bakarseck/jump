package cli

import (
	"fmt"
	"os"

	"github.com/Bakarseck/jump/internals/utils"
)

func SetupZsh() {
	// Installation de Zsh
	// if err := utils.ExecCommand("sudo", "apt-get", "update"); err != nil {
	// 	fmt.Fprintf(os.Stderr, "Erreur lors de la mise à jour des paquets: %v\n", err)
	// 	return
	// }
	// if err := utils.ExecCommand("sudo", "apt-get", "install", "zsh", "-y"); err != nil {
	// 	fmt.Fprintf(os.Stderr, "Erreur lors de l'installation de Zsh: %v\n", err)
	// 	return
	// }

	// // Changement du shell par défaut pour zsh
	// if err := utils.ExecCommand("chsh", "-s", "$(which zsh)"); err != nil {
	// 	fmt.Fprintf(os.Stderr, "Erreur lors du changement du shell par défaut: %v\n", err)
	// 	return
	// }

	// Installation d'Oh My Zsh
	if err := utils.ExecCommand("bash", "-c", "$(curl -fsSL https://raw.githubusercontent.com/ohmyzsh/ohmyzsh/master/tools/install.sh)"); err != nil {
		fmt.Fprintf(os.Stderr, "Erreur lors de l'installation de Oh My Zsh: %v\n", err)
		return
	}

	// Configuration supplémentaire ici, comme l'installation de plugins

	fmt.Println("Zsh et Oh My Zsh ont été installés et configurés avec succès.")
}
