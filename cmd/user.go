package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/Azaradel/git-change-user/pkg/confighandler"
	"github.com/Azaradel/git-change-user/pkg/userchanger"
	"github.com/spf13/cobra"
)

var add = &cobra.Command{
	Use:   "add",
	Short: "add user",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := confighandler.LoadConfig(config)
		if err != nil {
			log.Fatal(err)
		}
		err = c.AddUser(user, name, email)
		if err != nil {
			log.Fatal(err)
		}
		err = c.SaveConfig(c.Path)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("add successful", user, c.Users[user])
		os.Exit(0)
	},
}

var delete = &cobra.Command{
	Use:   "delete",
	Short: "delete user",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := confighandler.LoadConfig(config)
		if err != nil {
			log.Fatal(err)
		}
		err = c.DeleteUser(user)
		if err != nil {
			log.Fatal(err)
		}
		err = c.SaveConfig(c.Path)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("delete successful", user)
		os.Exit(0)
	},
}

var edit = &cobra.Command{
	Use:   "edit",
	Short: "edit user",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := confighandler.LoadConfig(config)
		if err != nil {
			log.Fatal(err)
		}
		err = c.EditUser(user, name, email)
		if err != nil {
			log.Fatal(err)
		}
		err = c.SaveConfig(c.Path)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("edit successful", user, c.Users[user])
		os.Exit(0)
	},
}

var show = &cobra.Command{
	Use:   "show",
	Short: "show users in storage",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := confighandler.LoadConfig(config)
		if err != nil {
			log.Fatal(err)
		}
		for k, v := range c.Users {
			fmt.Println(k, v)
		}
	},
}

var became = &cobra.Command{
	Use:   "became",
	Short: "choose user preset to use as your git global user",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := confighandler.LoadConfig(config)
		if err != nil {
			log.Fatal(err)
		}
		if _, ok := c.Users[user]; !ok {
			log.Fatal("become: user not found")
		}
		name, email = c.Users[user].Name, c.Users[user].Email
		err = userchanger.BecomeUser(name, email)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	},
}
