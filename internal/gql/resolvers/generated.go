package resolvers

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

import (
	"context"

	"github.com/yangfanx/poker-server/internal/gql"
	"github.com/yangfanx/poker-server/internal/gql/models"
)

type Resolver struct{}

func (r *mutationResolver) CreateUser(ctx context.Context, username string, currentIP string) (*models.User, error) {
	panic("not implemented")
}

func (r *mutationResolver) Renew(ctx context.Context, username string, currentIP string) (*models.User, error) {
	panic("not implemented")
}

func (r *mutationResolver) CreateTable(ctx context.Context, user models.ActionUser) (*models.Table, error) {
	panic("not implemented")
}

func (r *mutationResolver) NextHand(ctx context.Context, table *string) (bool, error) {
	panic("not implemented")
}

func (r *mutationResolver) JoinTable(ctx context.Context, user models.ActionUser, tableID *string) (*models.Player, error) {
	panic("not implemented")
}

func (r *mutationResolver) Stand(ctx context.Context, user models.ActionUser, tableID *string) (bool, error) {
	panic("not implemented")
}

func (r *mutationResolver) Sit(ctx context.Context, user models.ActionUser, tableID *string) (bool, error) {
	panic("not implemented")
}

func (r *mutationResolver) RequestBuyIn(ctx context.Context, user models.ActionUser) (bool, error) {
	panic("not implemented")
}

func (r *mutationResolver) ApproveBuyIn(ctx context.Context, user models.ActionUser, actor string) (bool, error) {
	panic("not implemented")
}

func (r *mutationResolver) Check(ctx context.Context, user models.ActionUser) (bool, error) {
	panic("not implemented")
}

func (r *mutationResolver) Bet(ctx context.Context, user models.ActionUser, amount int) (bool, error) {
	panic("not implemented")
}

func (r *mutationResolver) Fold(ctx context.Context, user models.ActionUser) (bool, error) {
	panic("not implemented")
}

func (r *mutationResolver) Straddle(ctx context.Context, user models.ActionUser) (bool, error) {
	panic("not implemented")
}

func (r *mutationResolver) Showcard(ctx context.Context, user models.ActionUser) (bool, error) {
	panic("not implemented")
}

func (r *queryResolver) Tables(ctx context.Context, status *models.TableStatus) ([]*models.Table, error) {
	panic("not implemented")
}

func (r *queryResolver) Users(ctx context.Context, username *string) ([]*models.User, error) {
	panic("not implemented")
}

// Mutation returns gql.MutationResolver implementation.
func (r *Resolver) Mutation() gql.MutationResolver { return &mutationResolver{r} }

// Query returns gql.QueryResolver implementation.
func (r *Resolver) Query() gql.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
