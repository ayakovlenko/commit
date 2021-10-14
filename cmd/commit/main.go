package main

import (
	"commit/internal"
	"commit/internal/git"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use: "commit",
	}

	semanticCommitTypes := []string{
		"feat",
		"fix",
		"refactor",
		"test",
		"docs",
		"chore",
		"ci",
		"build",
	}

	commitTypeCmd := &cobra.Command{
		Use: "type",
	}

	for _, commitType := range semanticCommitTypes {
		commitTypeCmd.AddCommand(createSemanticCommitCommand(commitType))
	}

	aliasCmd := &cobra.Command{
		Use: "set-alias <alias>",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return fmt.Errorf("exactly one argument is required; got: %d", len(args))
			}

			alias := fmt.Sprintf("alias.%s", args[0])
			out, err := git.Exec(
				[]string{"config", "--global", alias, "!commit type"},
			)
			fmt.Println(out)
			if err != nil {
				return err
			}

			return nil
		},
	}

	rootCmd.AddCommand(
		commitTypeCmd,
		aliasCmd,
	)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func createSemanticCommitCommand(commitType string) *cobra.Command {
	cmd := &cobra.Command{
		Use: commitType,
		RunE: func(cmd *cobra.Command, args []string) error {
			scope, err := cmd.Flags().GetString("scope")
			if err != nil {
				return err
			}

			msg, err := internal.MakeCommit(commitType, scope, args)
			if err != nil {
				return err
			}

			sout, err := git.Exec([]string{"commit", "-m", msg})
			fmt.Println(sout)
			if err != nil {
				return err
			}

			return nil
		},
	}

	cmd.Flags().StringP("scope", "s", "", "")

	return cmd
}
