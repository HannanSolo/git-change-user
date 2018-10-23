package userchanger

import (
	"errors"
	"fmt"
	"os/exec"
)

func BecomeUser(name, email string) error {
	err := useUsername(name)
	if err != nil {
		return err
	}
	err = useEmail(email)
	if err != nil {
		return err
	}
	return nil
}

func useUsername(name string) error {
	nameWithQuotes := parseUsername(name)
	cmd := exec.Command("git", "config", "--global", "user.name", nameWithQuotes)
	if err := cmd.Run(); err != nil {
		return errors.New("gcu: check if the git is installed and in $PATH")
	}
	return nil
}

func useEmail(email string) error {
	cmd := exec.Command("git", "config", "--global", "user.email", email)
	if err := cmd.Run(); err != nil {
		return errors.New("gcu: check if the git is installed and in $PATH")
	}
	return nil
}

func parseUsername(name string) string {
	return fmt.Sprintf("\"%s\"", name)
}
