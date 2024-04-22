package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func AddToken(homeDir, token string) {
	envPath := homeDir + "/.env"
	LoadEnv(envPath)

	// Charge le contenu du fichier .env
	content, err := os.ReadFile(envPath)
	if err != nil {
		log.Fatalf("Erreur lors de la lecture du fichier .env : %v", err)
	}

	// Assumez que `EncryptString` est votre fonction de chiffrement
	// et que `secretKey` est votre clé secrète définie ailleurs dans votre configuration
	key := os.Getenv("SECRET_KEY")
	encryptedUsername, err := EncryptString(token, key)
	if err != nil {
		log.Fatalf("Erreur lors du chiffrement du username : %v", err)
	}

	// Mise à jour ou ajout de la variable USERNAME_GITHUB dans le .env
	updatedContent := updateOrAddEnvVar(string(content), "GITHUB_TOKEN", encryptedUsername)

	// Écriture du nouveau contenu dans le fichier .env
	if err := os.WriteFile(envPath, []byte(updatedContent), 0644); err != nil {
		log.Fatalf("Erreur lors de l'écriture dans le fichier .env : %v", err)
	}

	log.Println("Le token GitHub chiffré a été enregistré avec succès dans le fichier .env.")
}

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
