package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/farwydi/vtbb-demo/domain"
)

func (r *mutationResolver) RefreshToken(ctx context.Context, refreshToken string, fingerprint string) (*Accses, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) Authorize(ctx context.Context, email string, password string, otp string, fingerprint string) (*Authoraze, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) Register(ctx context.Context, email string, password string, otpSecret string, fingerprint string) (*domain.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
