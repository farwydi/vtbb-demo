// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package graphql

import (
	"github.com/farwydi/vtbb-demo/domain"
)

type Accses struct {
	RefreshToken string `json:"refreshToken"`
	AccsesToken  string `json:"accsesToken"`
}

type Authoraze struct {
	Accses *Accses      `json:"accses"`
	User   *domain.User `json:"user"`
}
