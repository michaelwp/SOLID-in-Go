package main

import (
	"fmt"
)

// Logger handles logging
type Logger struct{}

func (l *Logger) Log(message string) {
	fmt.Println(message)
}

// User handles user data
type User struct {
	logger *Logger
	name   string
}

func (u *User) SetName(name string) {
	u.name = name
	u.logger.Log("User name updated to " + name)
}

func main() {
	logger := &Logger{}
	user := &User{logger: logger}

	user.SetName("John Doe")
}