package confighandler

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	yaml "gopkg.in/yaml.v2"
)

type gitUser struct {
	Name  string `yaml:"name"`
	Email string `yaml:"email"`
}

type Config struct {
	Users        map[string]gitUser `json:"users"`
	Config       []byte
	pathToConfig string
}

func LoadConfig(path string) (*Config, error) {
	c := &Config{
		Users: map[string]gitUser{
			"nobody": {
				Name:  "nobody",
				Email: "nobody",
			},
		},
		pathToConfig: path,
	}
	if c.pathToConfig == "" {
		c.pathToConfig = getHomeConfigPath()
	}

	f, err := ioutil.ReadFile(c.pathToConfig)
	if err != nil {
		return c, err
	}

	c.Config = f
	err = yaml.Unmarshal(c.Config, &c.Users)
	if err != nil {
		return c, err
	}
	return c, nil
}

func getHomeConfigPath() string {
	home := os.Getenv("HOME")
	return fmt.Sprintf("%v/.gcu.yaml", home)
}

func (c *Config) AddUser(user, name, email string) error {
	if dataLenghtVerification(user, name, email) {
		return errors.New("AddUser: user, name and email cannot be empty")
	}
	_, ok := c.Users[user]
	if ok {
		return errors.New("AddUser: user already exists")
	}

	c.Users[user] = gitUser{
		Name:  name,
		Email: email,
	}
	return nil
}

func (c *Config) DeleteUser(name string) error {
	return nil
}

func (c *Config) EditUser(user, name, email string) error {
	return nil
}

func checkLength(s string) bool {
	return len(s) == 0
}

func dataLenghtVerification(user, name, email string) bool {
	return checkLength(user) && checkLength(name) && checkLength(email)
}
