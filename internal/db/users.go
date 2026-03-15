package db

import (
	"database/sql"
)

func (db *DB) GetUserByEmail(email string) (*User, error) {
	var user User

	err := db.QueryRow("SELECT id, email, name, avatar_url, hash, created_at, updated_at FROM Users WHERE email = $1", email).Scan(&user.ID, &user.Email, &user.Name, &user.AvatarURL, &user.Hash, &user.CreatedAt, &user.UpdatedAt)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (db *DB) CreateUser(email, name, avatar_url, hash string) (*User, error) {
	var user User

	err := db.QueryRow(
		"INSERT INTO Users (email, name, avatar_url, hash) VALUES ($1, $2, $3, $4) RETURNING id, email, name, avatar_url, hash, created_at, updated_at",
		email, name, avatar_url, hash).Scan(&user.ID, &user.Email, &user.Name, &user.AvatarURL, &user.Hash, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
