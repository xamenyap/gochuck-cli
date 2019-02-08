package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/xamenyap/gochuck"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "gochuck-cli",
		Short: "Get Chuck Norris Jokes",
		RunE: func(cmd *cobra.Command, args []string) error {
			f, err := gochuck.GetRandom()
			if err != nil {
				return err
			}

			cmd.Println(f.Value)
			return nil
		},
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
