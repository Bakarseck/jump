package cli

import (
	"os"
	"os/exec"
)

func DownloadFile(url string, dest string) error {
	out, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer out.Close()

	cmd := exec.Command("wget", url, "-O", dest)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func CreateNewRustLib(libName string) error {
	cmd := exec.Command("cargo", "new", "--lib", libName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
