package cli

import (
	"fmt"
	"log"
	"os"

	"github.com/Bakarseck/jump/internals/utils"
)

// Récupérer et déchiffrer le token
func GetToken(encryptedToken string) string {
	// Vous devez générer et utiliser une clé secrète constante et sécurisée
	secretKey := os.Getenv("SECRET_KEY") // Doit être de 32 bytes pour AES-256

	// Supposons que encryptedToken est la version chiffrée du token que vous avez récupérée

	token, err := utils.DecryptString(encryptedToken, secretKey)
	if err != nil {
		log.Fatalf("Erreur lors du déchiffrement du token: %v", err)
	}

	fmt.Println(token)

	return token
}
