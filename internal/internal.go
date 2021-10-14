package internal

import (
	"fmt"
	"strings"
)

func MakeCommit(commitType string, scope string, words []string) (string, error) {
	if scope != "" {
		scope = fmt.Sprintf("(%s)", scope)
	}

	msg := fmt.Sprintf(
		"%s%s: %s",
		commitType,
		scope,
		strings.Join(words, " "),
	)

	return msg, nil
}
