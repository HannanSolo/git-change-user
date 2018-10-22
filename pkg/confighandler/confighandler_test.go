package confighandler

import (
	"os"
	"testing"
)

var expectedUsers = map[string]gitUser{
	"John": {
		Name:  "John Titor",
		Email: "john@titor.com",
	},
	"Jan": {
		Name:  "Jan Kowalski",
		Email: "jan@kowalski.com",
	},
}

func TestLoadConfig(t *testing.T) {
	c, _ := LoadConfig(".gcu.example.yaml")
	if c.Users["John"] != expectedUsers["John"] {
		t.Errorf("first person should be: %v, was: %v", expectedUsers["John"], c.Users["John"])
	}
	_, err := LoadConfig(".gcu.bad.example.yaml")
	if err == nil {
		t.Errorf("should return errror after receiving bad yaml")
	}

	_, err = LoadConfig("foo/bar")
	if err == nil {
		t.Errorf("should return error if file does not exist")
	}

	//mock the $HOME
	os.Setenv("HOME", "/")
	_, err = LoadConfig("")
	if err == nil {
		t.Errorf("should return error because file in $HOME does not exist")
	}
}

func TestAddUser(t *testing.T) {
	c, _ := LoadConfig(".gcu.example.yaml")

	c.AddUser("John", "John Titor", "john@titor.com")
	if c.Users["John"].Email != expectedUsers["John"].Email {
		t.Errorf("email should be: %v, was: %v", expectedUsers["John"].Email, c.Users["John"].Email)
	}

	c.AddUser("Jan", "Kowalski", "jan@kowalski.com")
	if c.Users["Jan"].Name != expectedUsers["Jan"].Name {
		t.Errorf("name should be: %v, was: %v", expectedUsers["Jan"].Name, c.Users["Jan"].Name)
	}

	err := c.AddUser("Jan", "Kowalski", "jan@kowalski.com")
	if err == nil {
		t.Errorf("should return error because user exists")
	}

	err = c.AddUser("", "", "")
	if err == nil {
		t.Errorf("should return error if one of property is empty")
	}

	err = c.AddUser("Foo", "Foo Bar", "foo@bar.com")
	if err != nil {
		t.Errorf("should not return error if every data about user was given")
	}
}

func TestDeleteUser(t *testing.T) {
	c, _ := LoadConfig(".gcu.example.yaml")

	err := c.DeleteUser("John")
	if err == nil {
		t.Errorf("should return error because referenced does not exist")
	}

	err = c.DeleteUser("John")
	if _, ok := c.Users["John"]; !ok && err == nil {
		t.Errorf("should delete user John")
	}
}

func TestEditUser(t *testing.T) {
	c, _ := LoadConfig(".gcu.example.yaml")

	err := c.EditUser("Foo", "John Smith", "john@smith.com")
	if err == nil {
		t.Errorf("should return error because referenced user does not exist")
	}

	err = c.EditUser("John", "John Smith", "john@smith.com")
	if err != nil && c.Users["John"].Name == "John Smith" {
		t.Errorf("should change name of user John")
	}

	err = c.EditUser("John", "John Smith", "john@smith.com")
	if err != nil && c.Users["John"].Email == "john@smith.com" {
		t.Errorf("should change name of user John")
	}
}

func TestSaveConfig(t *testing.T) {

}
