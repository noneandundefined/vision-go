package pkg

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

// The GitRoot function allows the program to find
// the main entry point of the root directory of the project,
// helps and simplifies the search for log files.
func GitRoot() string {
	cmd := exec.Command("git", "rev-parse", "--show-toplevel")
	output, err := cmd.Output()
	if err != nil {
		log.Printf("Failed to get Git root directory: %v\n", err)

		log.Println("Search root directory by go.mod")
		gomod_root := GOMODRoot()
		if gomod_root != "" {
			return gomod_root
		}
	}

	gitRoot := string(output)
	gitRoot = gitRoot[:len(gitRoot)-1]

	err = os.Chdir(gitRoot)
	if err != nil {
		log.Printf("Failed to change directory: %v\n", err)
	}

	return gitRoot
}

// The GOMODRoot function allows the program to find
// the main entry point to the root directory of the project,
// helps and simplifies the search for log files.
func GOMODRoot() string {
	dir := "."
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir
		}

		parentDir := filepath.Dir(dir)
		if parentDir == dir {
			break
		}
		dir = parentDir
	}

	log.Println("Project root not found")
	return ""
}
