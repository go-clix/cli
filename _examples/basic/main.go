package main

import (
	"fmt"
	"os"

	"github.com/go-clix/cli"
)

func main(){
	// create the root command
	rootCmd := cli.Command{
		Use: "greet",
		Short: "print a message",
		Run: func(cmd *cli.Command, args []string) error {
			fmt.Println("Hello from Cli-X!")
			return nil
		},
	}

	// run and check for errors
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
