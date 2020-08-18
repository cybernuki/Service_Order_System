package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/cybernuki/Service_Order_System/graph/generated"
	"github.com/cybernuki/Service_Order_System/graph/model"
	"github.com/cybernuki/Service_Order_System/internal/database/models"
	"github.com/cybernuki/Service_Order_System/internal/jwt"
)

// CreateUser - mutation that creates an user and returns the generated token
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (string, error) {
	var user models.SchemaUser
	user.Email = input.Email
	user.Password = input.Password
	user.FirstName = input.FirstName
	user.LastName = input.LastName

	err := user.Create()
	if err != nil {
		return "", err
	}
	token, err := jwt.GenerateToken(user.Email)
	return token, err
}

func (r *mutationResolver) CreateTechnician(ctx context.Context, input model.NewTechnician) (*model.Technician, error) {
	fmt.Errorf("not implemented technician")
	return nil, nil
}

func (r *mutationResolver) CreateTelevision(ctx context.Context, input model.NewTelevision) (*model.Television, error) {
	var newTv models.SchemaTelevision

	newTv.ModelTV = input.Model
	newTv.Brand = input.Brand
	err := newTv.Create()
	if err != nil {
		return nil, err
	}
	return &model.Television{ID: fmt.Sprint(newTv.ID), Model: newTv.ModelTV, Brand: newTv.Brand}, nil
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
	var television models.SchemaTelevision

	allSchemas, err := television.All()

	var allTelevisions = make([]*model.Television, len(allSchemas))

	for i := 0; i < len(allSchemas); i++ {
		var tv model.Television
		tv.ID = fmt.Sprint(allSchemas[i].ID)
		tv.Model = allSchemas[i].ModelTV
		tv.Brand = allSchemas[i].Brand
		allTelevisions[i] = &tv
	}
	return allTelevisions, err
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
