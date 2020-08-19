package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"strconv"

	"github.com/cybernuki/Service_Order_System/graph/generated"
	"github.com/cybernuki/Service_Order_System/graph/model"
	"github.com/cybernuki/Service_Order_System/internal/auth"
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

func (r *mutationResolver) CreateOrder(ctx context.Context, input model.NewOrder) (*model.OrderCreated, error) {
	var tv models.SchemaTelevision
	var newOrder models.SchemaOrder

	// Authorization
	user := auth.ForContext(ctx)
	if user == nil {
		return nil, tools.NewError("access denied")
	}
	models.Db.Find(&user, "email = ?", user.Email)
	// get the tv
	tv.ModelTV = input.Television.Model
	tv.Brand = input.Television.Brand
	tv.Create()

	// create the new service order
	newOrder.Description = input.Description
	err := newOrder.Create(*user, tv)
	if err != nil {
		return nil, err
	}
	token, err := jwt.GenerateToken(strconv.Itoa(int(newOrder.ID)))
	if err != nil {
		return nil, err
	}
	order := model.OrderCreated{
		Token: token,
		URL:   "localhost:8000/orders/" + strconv.Itoa(int(newOrder.ID)),
	}
	return &order, nil
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
