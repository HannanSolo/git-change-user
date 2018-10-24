package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var version = "v0.1.0"
var config, user, name, email string

var rootCmd = &cobra.Command{
	Use:     "gcu",
	Version: version,
	Short:   "Gcu provides easy handling of git global user.",
	Long: `Fast and easy way to handle your git global user by using simple yaml
file in your $HOME/.gcu.yaml as a storage.`,
}

func Execute() {
	add.Flags().StringVarP(&user, "user", "u", "", "choose user to be add to storage")
	add.Flags().StringVarP(&name, "name", "n", "", "choose name for git user")
	add.Flags().StringVarP(&email, "email", "e", "", "choose email for git user")
	edit.Flags().StringVarP(&user, "user", "u", "", "choose user to be edited to storage")
	edit.Flags().StringVarP(&name, "name", "n", "", "choose name for git user")
	edit.Flags().StringVarP(&email, "email", "e", "", "choose email for git user")
	delete.Flags().StringVarP(&user, "user", "u", "", "choose user to be deleted")
	became.Flags().StringVarP(&user, "user", "u", "", "choose user to set as a git global user")
	rootCmd.AddCommand(add, edit, delete, show, became)
	rootCmd.PersistentFlags().StringVar(&config, "config", "", "Choose file to be used as a config.")
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
