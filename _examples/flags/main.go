package main

import (
	"fmt"
	"os"

	"github.com/go-clix/cli"
)

// A `pflag.FlagSet` can be accessed per command using `*Command.Flags()`:

func applyCmd() *cli.Command {
    cmd := &cli.Command{
        Use: "apply",
        Short: "apply the changes",
    }

    force := cmd.Flags().BoolP("force", "f", false, "skip checks")

    cmd.Run = func(cmd *cli.Command, args []string) error {
        fmt.Println("applied", args[0])
        if *force {
            fmt.Println("The force was with us.")
        }
        return nil
    }
    return cmd
}


func main() {
	rootCmd := &cli.Command{
		Use: "kubectl",
		Short: "Kubernetes management tool",
	}

	// add the child command
	rootCmd.AddCommand(
		applyCmd(),
	)

	// run and check for errors
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
