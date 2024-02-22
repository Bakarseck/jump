package cli

import (
	"encoding/json"
	"fmt"
	"jump/internals/models"
	"jump/internals/utils"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func Add(cmd *cobra.Command, args []string) {
	dir, path := "", ""
	if len(args) == 1 {
		if args[0] == "." {
			dir = "."
			if p, err := os.Getwd(); err != nil {
				fmt.Println(err.Error())
				return
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

	dirs, ok := utils.LoadDirs()
	if ok {
		return
	}

	exists, differentPath := false, false
	for _, d := range dirs {
		if d.Dir == dir && d.Path == path {
			fmt.Println("Error: Directory with the same name and path already exists.")
			return
		} else if d.Dir == dir && d.Path != path {
			exists, differentPath = true, true
			break
		}
	}

	if exists && differentPath {
		parentDir := strings.Split(path, "/")
		if len(parentDir) > 1 {
			dir = fmt.Sprintf("%s-%s", parentDir[len(parentDir)-2], dir)
		}
	}

	newDir := models.Dirs{Dir: dir, Path: path}
	dirs = append(dirs, newDir)

	updatedContent, err := json.Marshal(dirs)
	if err != nil {
		log.Println(err.Error())
		return
	}

	os.WriteFile(models.PathJson, []byte(updatedContent), 0777)
}
