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
	"github.com/cybernuki/Service_Order_System/internal/tools"
)

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

func (r *mutationResolver) CreateTechnician(ctx context.Context, input model.NewTechnician) (string, error) {
	var technician models.SchemaTechnician
	technician.Email = input.Email
	technician.Password = input.Password
	technician.FirstName = input.FirstName
	technician.LastName = input.LastName

	err := technician.Create()
	if err != nil {
		return "", err
	}
	token, err := jwt.GenerateToken(technician.Email)
	return token, err
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

func (r *mutationResolver) CreateOrder(ctx context.Context, input model.NewOrder, login model.Login) (*model.OrderCreated, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateOrder(ctx context.Context, token string, input model.NewOrder) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	var user models.SchemaUser
	user.Email = input.Email
	user.Password = input.Password
	correct := user.Authenticate()
	if !correct {
		return "", &tools.WrongUsernameOrPasswordError{}
	}
	token, err := jwt.GenerateToken(user.Email)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (r *mutationResolver) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error) {
	email, err := jwt.ParseToken(input.Token)
	if err != nil {
		return "", tools.NewError("access denied")
	}
	token, err := jwt.GenerateToken(email)
	if err != nil {
		return "", err
	}
	return token, nil
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

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *mutationResolver) UpdateUser(ctx context.Context, input model.NewUser) (string, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *mutationResolver) UpdateTechnician(ctx context.Context, input model.NewTechnician) (string, error) {
	panic(fmt.Errorf("not implemented"))
}
