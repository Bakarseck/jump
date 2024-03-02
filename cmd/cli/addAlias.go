package cli

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"

	"github.com/Bakarseck/jump/internals/utils"
)

func AddAlias() error {

	homeDir, err := os.UserHomeDir()
	scriptPath := homeDir + "/alias.sh"

	fmt.Println(homeDir + "/alias.sh")

	utils.WriteFile(homeDir+"/alias.sh", "alias.sh")

	if err != nil {
		return fmt.Errorf("impossible de trouver le répertoire personnel: %v", err)
	}

	shellConfigPath := filepath.Join(homeDir, ".zshrc")

	scriptFile, err := os.Open(scriptPath)
	if err != nil {
		return fmt.Errorf("erreur lors de l'ouverture du fichier script: %v", err)
	}
	defer scriptFile.Close()

	configFile, err := os.OpenFile(shellConfigPath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("erreur lors de l'ouverture du fichier de configuration du shell: %v", err)
	}
	defer configFile.Close()

	scanner := bufio.NewScanner(scriptFile)
	for scanner.Scan() {
		_, err := configFile.WriteString(scanner.Text() + "\n")
		if err != nil {
			return fmt.Errorf("erreur lors de l'écriture dans le fichier de configuration du shell: %v", err)
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("erreur lors de la lecture du fichier script: %v", err)
	}

	fmt.Println("Les alias ont été ajoutés avec succès. Veuillez exécuter `source ~/.zshrc` pour appliquer les changements.")

	return nil
}
