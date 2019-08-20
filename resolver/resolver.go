package resolver

import (
	"context"

	"github.com/kanehara/prisma-template/gqlgen"
	"github.com/kanehara/prisma-template/prisma"
)

type Resolver struct {
	Prisma *prisma.Client
}

func (r *Resolver) Query() gqlgen.QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Empty(ctx context.Context) (*string, error) {
	return nil, nil
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) Empty(ctx context.Context) (*string, error) {
	return nil, nil
}

type subscriptionResolver struct{ *Resolver }

func (r *subscriptionResolver) Empty(ctx context.Context) (<-chan *string, error) {
	return nil, nil
}
