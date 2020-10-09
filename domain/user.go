package domain

import (
	"github.com/google/uuid"
	"hash/fnv"
)

var userOffset = []byte{
	0x00,
	0x00,
}

type User struct {
	UserID int
}

func (u User) ID() uuid.UUID {
	return uuid.NewHash(
		fnv.New128a(),
		uuid.NameSpaceOID,
		append(userOffset, IDtoBytes(u.UserID)...),
		5,
	)
}
