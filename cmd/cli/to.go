package cli

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/Bakarseck/jump/internals/utils"

	"github.com/spf13/cobra"
)

func To(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		log.Println("Nom du répertoire manquant")
		return
	}

	dir := args[0]

	dirs, ok := utils.LoadDirs()
	if ok {
		return
	}

	for _, v := range dirs {
		if v.Dir == dir {
			cmd := exec.Command("bash", "-c", fmt.Sprintf("cd %s && exec $SHELL", v.Path))
			cmd.Stdin = os.Stdin
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr

			err := cmd.Run()
			if err != nil {
				fmt.Println("Erreur lors du changement du répertoire de travail:", err)
				return
			}
			break
		}
	}
}
