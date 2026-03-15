package auth

import (
	"errors"
	"fmt"
	"regexp"
	"spotify_mod/internal/db"
	"unicode"
)

var (
	emailRegex    = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	ErrUserExists = errors.New("User already exists")
)

type RegisterRequest struct {
	Email    string
	Name     string
	Password string
}

func (r *RegisterRequest) Validate() error {
	if !emailRegex.MatchString(r.Email) {
		return errors.New("Invalid email format")
	}

	if len(r.Name) < 4 {
		return errors.New("Username error! Must be at least 4 characters long!")
	}

	if len(r.Password) < 8 {
		return errors.New("Passoword error! Must be at least 8 characters long!")
	}

	hasLetter := false
	hasDigit := false
	hasSpecial := false

	for _, char := range r.Password {
		if unicode.IsLetter(char) {
			hasLetter = true
		} else if unicode.IsDigit(char) {
			hasDigit = true
		} else if unicode.IsPunct(char) || unicode.IsSymbol(char) {
			hasSpecial = true
		}
	}

	if !hasLetter {
		return errors.New("Password must contain at least one letter!")
	}

	if !hasDigit {
		return errors.New("Passoword must containt at least one digit!")
	}

	if !hasSpecial {
		return errors.New("Password must contain at least one special character!")
	}

	return nil
}

func Register(database *db.DB, req RegisterRequest) (*db.User, error) {
	if err := req.Validate(); err != nil {
		return nil, fmt.Errorf("Validation failed: %w", err)
	}

	existing, err := database.GetUserByEmail(req.Email)
	if err != nil {
		return nil, err
	}

	if existing != nil {
		return nil, ErrUserExists
	}

	hash, err := HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user, err := database.CreateUser(req.Email, req.Name, "", hash)
	if err != nil {
		return nil, err
	}

	return user, nil
}
