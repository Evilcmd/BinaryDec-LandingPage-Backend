package models

import (
	"fmt"
)

var (
	ErrorUserExists = fmt.Errorf("user already exists")
)

type User struct {
	Name    string `json:"name" bson:"name"`
	Email   string `json:"email" bson:"email"`
	Company string `json:"company" bson:"company"`
}
