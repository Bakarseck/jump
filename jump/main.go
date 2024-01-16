package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

type Dirs struct {
	Dir  string
	Path string
}

var (
	pathJson  string
)

func main() {

	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("Erreur lors de la récupération du répertoire personnel : %v", err)
	}

	pathJson = homeDir + "/jump.json"

	var rootCmd = &cobra.Command{
		Use:   "jump",
		Short: "Un gestionnaire de répertoires flexible et rapide",
		Long: "\n" + `
		Jump est un gestionnaire de répertoires construit avec amour en Go.
		Il permet d'ajouter et de naviguer facilement entre différents répertoires.`,
	}

	// Définit une commande pour ajouter un répertoire
	var cmdAdd = &cobra.Command{
		Use:   "add",
		Short: "Ajouter un répertoire",
		Long: "\n" + `
		Ajoute un nouveau répertoire à la liste de gestion.
		Vous pouvez spécifier un nom et un chemin, ou utiliser '.' pour ajouter le répertoire courant.
		Exemple: 'jump add projet /chemin/vers/projet' ou 'jump add .'
		Pour pouvoir utiliser le 'jump add .', il faut naviguer jusqu'au répértoire que tu veux ajouter.`,

		Run:   addDirectory,
	}

	// Définit une commande pour changer de répertoire
	var cmdJump = &cobra.Command{
		Use:   "to",
		Short: "Jump vers un répertoire",
		Long: "\n"+ `
		Change le répertoire courant en celui spécifié.
		Utilise le nom du répertoire tel qu'ajouté avec la commande 'add'.
		Exemple: 'jump to projet' pour aller au répertoire nommé 'projet'.`,
		Run:   changeDirectory,
	}

	// Attache les commandes à l'application principale
	rootCmd.AddCommand(cmdAdd)
	rootCmd.AddCommand(cmdJump)

	// Exécute l'application
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func addDirectory(cmd *cobra.Command, args []string) {
	dir, path := "", ""

	if len(args) == 1 {
		if args[0] == "." {
			dir = "."
			if p, err := os.Getwd(); err != nil {
				fmt.Println(err.Error())
			} else {
				e := strings.Split(p, "/")
				dir = e[len(e)-1]
				path = p
			}
		}
	} else if len(args) == 2 {
		dir = args[0]
		path = args[1]
	}

	var dirs []Dirs
	if _, err := os.Stat(pathJson); err == nil {
		Content, err := os.ReadFile(pathJson)
		if err != nil {
			log.Println(err.Error())
			return
		}
		json.Unmarshal(Content, &dirs)
	}

	newDir := Dirs{Dir: dir, Path: path}
	dirs = append(dirs, newDir)

	updatedContent, err := json.Marshal(dirs)
	if err != nil {
		log.Println(err.Error())
		return
	}

	os.WriteFile(pathJson, []byte(updatedContent), 0777)
}

func changeDirectory(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		log.Println("Nom du répertoire manquant")
		return
	}

	dir := args[0]

	var dirs []Dirs
	if _, err := os.Stat(pathJson); err == nil {
		Content, err := os.ReadFile(pathJson)
		if err != nil {
			log.Println(err.Error())
			return
		}
		json.Unmarshal(Content, &dirs)
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
