package main

import (
	"fmt"
	"os"

	"github.com/go-clix/cli"
)

// use a func to return a Command instead of
// a global variable and `init()`
func applyCmd() *cli.Command {
	cmd := &cli.Command{
		Use: "apply",
		Short: "apply the changes",
	}

	cmd.Run = func(cmd *cli.Command, args []string) error {
		fmt.Println("applied", args[0])
		return nil
	}

	return cmd
}

func main() {
	rootCmd := &cli.Command{
		Use: "subcommands",
		Short: "This command has sub commands.",
	}
	rootCmd.AddCommand(applyCmd())
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
