package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/ottolauncher/church-app/graph/generated"
	"github.com/ottolauncher/church-app/graph/model"
)

// Login is the resolver for the login field.
func (r *authOpsResolver) Login(ctx context.Context, obj *model.AuthOps, email string, password string) (interface{}, error) {
	panic(fmt.Errorf("not implemented: Login - login"))
}

// Register is the resolver for the register field.
func (r *authOpsResolver) Register(ctx context.Context, obj *model.AuthOps, input model.NewUser) (interface{}, error) {
	panic(fmt.Errorf("not implemented: Register - register"))
}

// Auth is the resolver for the auth field.
func (r *mutationResolver) Auth(ctx context.Context) (*model.AuthOps, error) {
	panic(fmt.Errorf("not implemented: Auth - auth"))
}

// Task is the resolver for the task field.
func (r *mutationResolver) Task(ctx context.Context) (*model.TaskMutation, error) {
	panic(fmt.Errorf("not implemented: Task - task"))
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	panic(fmt.Errorf("not implemented: User - user"))
}

// Protected is the resolver for the protected field.
func (r *queryResolver) Protected(ctx context.Context) (string, error) {
	panic(fmt.Errorf("not implemented: Protected - protected"))
}

// AuthOps returns generated.AuthOpsResolver implementation.
func (r *Resolver) AuthOps() generated.AuthOpsResolver { return &authOpsResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type authOpsResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
