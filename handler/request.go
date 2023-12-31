package handler

import (
	"fmt"
	"net/mail"
)

func errParamIsRequired(name string, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required.", name, typ)
}

type registerUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func validEmail(email string) bool {
	_, err := mail.ParseAddress(email)

	return err == nil
}

func (user *registerUserRequest) validate() error {

	if user.Name == "" && user.Email == "" && user.Password == "" {
		return fmt.Errorf("request body is empty or malformed")
	}

	if user.Name == "" {
		return errParamIsRequired("name", "string")
	}
	if user.Email == "" {
		return errParamIsRequired("email", "string")
	}
	if !validEmail(user.Email) {
		return fmt.Errorf("The email format is invalid.")
	}
	if user.Password == "" {
		return errParamIsRequired("password", "string")
	}

	return nil
}

type loginUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (user *loginUserRequest) validate() error {

	if user.Email == "" && user.Password == "" {
		return fmt.Errorf("request body is empty or malformed")
	}
	if user.Email == "" {
		return errParamIsRequired("email", "string")
	}
	if user.Password == "" {
		return errParamIsRequired("password", "string")
	}

	return nil
}
