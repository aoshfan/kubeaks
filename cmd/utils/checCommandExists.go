package utils

import (
	"fmt"
	"os"
	"os/exec"
)

func CommandExists(cmd string) string {
	path, err := exec.LookPath(cmd)
	if err != nil {
		fmt.Println("Please install kubelogin first")
		os.Exit(1)
	}

	return path
}
