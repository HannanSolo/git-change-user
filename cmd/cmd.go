package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var version = "v0.1.0"

var rootCmd = &cobra.Command{
	Use:   "gcu",
	Short: "Gcu provides easy handling of git global user.",
	Long: `Fast and easy way to handle your git global user
by using simple yaml file in your $HOME/.gcu.yaml`,
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Check the version",
	Long:  `Prints the version of your gcu binary in semantic versioning`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("gcu version:", version)
	},
}

var createCmd = &cobra.Command{
	Use:   "create [sub]",
	Short: "Create new [user] or [config]",
	Long: `Create new user in config file or creates 
new config file based on present git config`,
	Args: cobra.MinimumNArgs(1),
}

var userCmd = &cobra.Command{
	Use:   "gcu create user [user, name, email]",
	Short: "Create new user in config file",
	Long:  "Create new user in config file",
	Args:  cobra.MinimumNArgs(3),
}

func Execute() {
	rootCmd.AddCommand(versionCmd, createCmd)
	createCmd.AddCommand(userCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Make sure your config file exists.")
		os.Exit(1)
	}
}
