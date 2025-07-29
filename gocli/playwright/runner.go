package playwright

import (
	"fmt"
	"gocli/types"
	"os/exec"
)

func RunScript(platform, action string, args ...string) error {
	// Example: node playwright/twitter/post.js --content "Hello"
	cmdArgs := append([]string{"playwright", platform, action}, args...)
	cmd := exec.Command("node", cmdArgs...)
	out, err := cmd.CombinedOutput()
	fmt.Printf("Playwright output: %s\n", string(out))
	return err
}
