package utils

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/Bakarseck/jump/internals/models"
)

func LoadEnv(path string) error {
	file, err := os.Open(path)
	if err != nil {
		log.Println(err.Error())
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			log.Println("Your env file must be set")
		}
		key := parts[0]
		value := parts[1]
		err := os.Setenv(key, value)
		if err != nil {
			return err
		}
	}
	return scanner.Err()
}

func LoadDirs() ([]models.Dirs, bool) {
	var dirs []models.Dirs
	if _, err := os.Stat(models.PathJson); err == nil {
		content, err := os.ReadFile(models.PathJson)
		if err != nil {
			log.Println(err.Error())
			return nil, true
		}
		json.Unmarshal(content, &dirs)
	}
	return dirs, false
}

func UpdateEnvFile(username, email string) {
	envContent := fmt.Sprintf("USERNAME=%s\nEMAIL=%s\n", username, email)
	if err := os.WriteFile(models.HomeDir+"/.env", []byte(envContent), 0644); err != nil {
		log.Fatalf("Erreur lors de l'écriture dans le fichier .env : %v", err)
	}
}

func ExecCommand(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func WriteFile(filePath, fileContentPath string) {

	content, err := os.ReadFile(fileContentPath)

	if err != nil {
		log.Fatalf("Impossible de lire le fichier : %v", err)
	}

	err = os.WriteFile(filePath, content, 0644)
	if err != nil {
		log.Fatalf("Impossible d'écrire dans le fichier : %v", err)
	}
}
