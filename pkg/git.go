package pkg

import (
	"log"
	"os"
	"os/exec"
)

// The GitRoot function allows the program to find
// the main entry point of the root directory of the project,
// helps and simplifies the search for log files.
func GitRoot() string {
	cmd := exec.Command("git", "rev-parse", "--show-toplevel")
	output, err := cmd.Output()
	if err != nil {
		log.Printf("Failed to get Git root directory: %v\n", err)
	}

	gitRoot := string(output)
	gitRoot = gitRoot[:len(gitRoot)-1]

	err = os.Chdir(gitRoot)
	if err != nil {
		log.Printf("Failed to change directory: %v\n", err)
	}

	return gitRoot
}
