package git

import (
	"os/exec"
	"strings"
)

func Exec(args []string) (string, error) {
	gitCmd := exec.Command("git", args...)

	bout, err := gitCmd.CombinedOutput()
	sout := strings.TrimSpace(string(bout))

	return sout, err
}
