package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/cybernuki/Service_Order_System/graph/generated"
	"github.com/cybernuki/Service_Order_System/graph/model"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (string, error) {
	user := model.User{FirstName: input.FirstName, LastName: input.LastName}
	return "User: " + user.FirstName + " " + user.LastName, nil
}

func (r *mutationResolver) CreateTechnician(ctx context.Context, input model.NewTechnician) (*model.Technician, error) {
	fmt.Errorf("not implemented technician")
	return nil, nil
}

func (r *mutationResolver) CreateTelevision(ctx context.Context, input model.NewTelevision) (*model.Television, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateOrder(ctx context.Context, input model.NewOrder) (*model.Order, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateUser(ctx context.Context, input model.NewUser) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateTechnician(ctx context.Context, input model.NewTechnician) (*model.Technician, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateOrder(ctx context.Context, input model.NewOrder) (*model.Order, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Technicians(ctx context.Context) ([]*model.Technician, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Televisions(ctx context.Context) ([]*model.Television, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Orders(ctx context.Context) ([]*model.Order, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Order(ctx context.Context, input string) (*model.Order, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
