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

func (g gitUser) String() string {
	return fmt.Sprintf("\tname: %v\n\temail: %v\n", g.Name, g.Email)
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
		return errors.New("add user: user, name and email cannot be empty")
	}
	_, ok := c.Users[user]
	if ok {
		return errors.New("add user: user already exists")
	}

	c.Users[user] = gitUser{
		Name:  name,
		Email: email,
	}
	return nil
}

func (c *Config) DeleteUser(user string) error {
	_, ok := c.Users[user]
	if !ok {
		return errors.New("delete user: user already does not exist")
	}
	delete(c.Users, user)
	return nil
}

func (c *Config) EditUser(user, name, email string) error {
	if dataLenghtVerification(user, name, email) {
		return errors.New("edit user: user, name and email cannot be empty")
	}
	if _, ok := c.Users[user]; !ok {
		return errors.New("edit user: user does not exist")
	}
	g := &gitUser{
		Name:  name,
		Email: email,
	}
	c.Users[user] = *g
	return nil
}

func (c *Config) SaveConfig(path string) error {
	f, err := os.Open(path)
	defer f.Close()
	if err != nil && os.IsNotExist(err) {
		f, err = os.Create(path)
	} else if err != nil {
		return err
	}
	_, err = f.Write(c.Config)
	if err != nil {
		return err
	}
	err = f.Sync()

	return err
}

func checkLength(s string) bool {
	return len(s) == 0
}

func dataLenghtVerification(user, name, email string) bool {
	return checkLength(user) && checkLength(name) && checkLength(email)
}
