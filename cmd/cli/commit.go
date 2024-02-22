package cli

import (
	"fmt"
	"log"

	"github.com/Bakarseck/jump/internals/utils"
	"github.com/spf13/cobra"
)

var (
	Files   []string
	Message string
)

func CommitRepo(cmd *cobra.Command, args []string) {
	if len(Files) == 0 {
		if err := utils.ExecCommand("git", "add", "."); err != nil {
			log.Fatalf("Erreur lors de l'ajout des fichiers : %v", err)
		}
	} else {
		for _, file := range Files {
			if err := utils.ExecCommand("git", "add", file); err != nil {
				log.Fatalf("Erreur lors de l'ajout du fichier %s : %v", file, err)
			}
		}
	}

	if err := utils.ExecCommand("git", "commit", "-m", Message); err != nil {
		log.Fatalf("Erreur lors du commit : %v", err)
	}

	if err := utils.ExecCommand("git", "push"); err != nil {
		log.Fatalf("Erreur lors du push : %v", err)
	}

	fmt.Println("Commit effectué avec succès.")
}
