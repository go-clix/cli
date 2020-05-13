package main

import (
	"fmt"
	"os"

	"github.com/go-clix/cli"
)

// Keep in mind that in `--help` outputs the command will still be called `apply`.


func applyCmd() *cli.Command {
	return &cli.Command{
		Use: "apply",
		Aliases: []string{"make", "do"},
		Short: "apply the changes",
	}
}

func main() {
	rootCmd := &cli.Command{
		Use: "aliases",
		Short: "Subcommand has aliases.",
	}
	rootCmd.AddCommand(applyCmd())
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
