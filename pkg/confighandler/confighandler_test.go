package confighandler

import (
	"fmt"
	"io/ioutil"
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

	err := c.DeleteUser("Foo")
	if err == nil {
		t.Errorf("should return error because user does not exist")
	}

	err = c.DeleteUser("John")
	if _, ok := c.Users["John"]; ok && err != nil {
		t.Errorf("should delete user John")
	}
}

func TestEditUser(t *testing.T) {
	c, _ := LoadConfig(".gcu.example.yaml")

	err := c.EditUser("Foo", "John Smith", "john@smith.com")
	if err == nil {
		t.Errorf("should return error because user does not exist")
	}

	err = c.EditUser("", "", "")
	if err == nil {
		t.Errorf("should return error because data cannot be an empty string")
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
	c, _ := LoadConfig(".gcu.example.yaml")
	path := ".test.config.yaml"

	_ = c.SaveConfig(path)
	defer os.Remove(path)
	if _, err := os.Stat(".test.config.yaml"); os.IsNotExist(err) {
		t.Errorf("file should exist after saving")
	}

	savedData, _ := ioutil.ReadFile(path)
	if !testSlice(savedData, c.Config) {
		t.Errorf("saved data should be the same as before")
	}
}
func TestUserString(t *testing.T) {
	c, _ := LoadConfig(".gcu.example.yaml")

	if func(a interface{}) bool {
		_, ok := a.(fmt.Stringer)
		return !ok
	}(c.Users["John"]) {
		t.Errorf("git user should implement Stringer interface")
	}

	str := c.Users["John"].String()
	expected := "\tname: John Titor\n\temail: john@titor.com\n"
	if str != expected {
		t.Errorf("should return string: %v, was: %v", expected, str)
	}
}

func testSlice(a, b []byte) bool {
	if (a == nil) || (b == nil) {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
