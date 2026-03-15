package db

import "time"

type User struct {
	ID        int       `db:"id"`
	Email     string    `db:"email"`
	Name      string    `db:"name"`
	AvatarURL string    `db:"avatar_url"`
	Hash      string    `db:"hash"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type OAuthAccount struct {
	ID             int       `db:"id"`
	UserID         int       `db:"user_id"`
	Provider       string    `db:"provider"`
	ProviderUserID string    `db:"provider_user_id"`
	AccessToken    string    `db:"access_token"`
	RefreshToken   string    `db:"refresh_token"`
	TokenExpiresAt time.Time `db:"token_expires_at"`
	CreatedAt      time.Time `db:"created_at"`
}

type SessionAccount struct {
	ID        int       `db:"id"`
	UserID    int       `db:"user_id"`
	Token     string    `db:"token"`
	ExpiresAt time.Time `db:"expires_at"`
	CreatedAt time.Time `db:"created_at"`
}
