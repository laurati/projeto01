package configuration

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
)

func EnvironmentSetup() {

	// returns the path of the executable being executed
	exePath, err := os.Executable()
	if err != nil {
		log.Fatalf("%v", err)
	}
	// returns the parent directory of the executable path
	exeDir := filepath.Dir(exePath)
	// This Git command returns the path of the project's root directory
	cmd := exec.Command("git", "rev-parse", "--show-toplevel")
	cmd.Dir = exeDir
	// project root directory path
	output, err := cmd.Output()
	if err != nil {
		log.Fatalf("%v", err)
	}

	projectRoot := strings.TrimSpace(string(output))
	envFilePath := filepath.Join(projectRoot, ".env")

	err = godotenv.Load(envFilePath)
	if err != nil {
		log.Fatalf("Error loading env file %v", err)
		return
	}
	log.Println("YOUR ENV FILE IS LOADED")
}
