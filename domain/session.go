package domain

import (
	"github.com/google/uuid"
	"time"
)

type Session struct {
	ID           int
	User         *User
	RefreshToken uuid.UUID
	UserAgent    string
	Fingerprint  string
	IP           string
	ExpiresIn    time.Time
}
