package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func AddUsernameGithub(homeDir, usernameGithub string) {
	envPath := homeDir + "/.env"
	content, err := readEnvFile(envPath)
	if err != nil {
		log.Fatalf("Erreur lors de la lecture du fichier .env : %v", err)
	}

	updatedContent := updateOrAddEnvVar(content, "USERNAME_GITHUB", usernameGithub)

	if err := os.WriteFile(envPath, []byte(updatedContent), 0644); err != nil {
		log.Fatalf("Erreur lors de l'écriture dans le fichier .env : %v", err)
	}
}

func readEnvFile(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	content := ""
	for scanner.Scan() {
		content += scanner.Text() + "\n"
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return content, nil
}

func updateOrAddEnvVar(content, varName, varValue string) string {
	var updated bool
	updatedLines := ""
	lines := strings.Split(content, "\n")

	for _, line := range lines {
		if strings.HasPrefix(line, varName+"=") {
			updatedLines += fmt.Sprintf("%s=%s\n", varName, varValue)
			updated = true
		} else if line != "" {
			updatedLines += line + "\n"
		}
	}

	if !updated {
		updatedLines += fmt.Sprintf("%s=%s\n", varName, varValue)
	}

	return updatedLines
}
