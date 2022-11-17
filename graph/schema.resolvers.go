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
func (r *mutationResolver) Login(ctx context.Context, email string, password string) (string, error) {
	panic(fmt.Errorf("not implemented: Login - login"))
}

// Register is the resolver for the register field.
func (r *mutationResolver) Register(ctx context.Context, input model.NewUser) (*model.User, error) {
	panic(fmt.Errorf("not implemented: Register - register"))
}

// CreateTask is the resolver for the createTask field.
func (r *mutationResolver) CreateTask(ctx context.Context, input model.NewTask) (*model.Task, error) {
	panic(fmt.Errorf("not implemented: CreateTask - createTask"))
}

// UpdateTask is the resolver for the updateTask field.
func (r *mutationResolver) UpdateTask(ctx context.Context, input model.UpdateTask) (*model.Task, error) {
	panic(fmt.Errorf("not implemented: UpdateTask - updateTask"))
}

// DeleteTask is the resolver for the deleteTask field.
func (r *mutationResolver) DeleteTask(ctx context.Context, id string) (bool, error) {
	panic(fmt.Errorf("not implemented: DeleteTask - deleteTask"))
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, filter map[string]interface{}) (*model.User, error) {
	panic(fmt.Errorf("not implemented: User - user"))
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context, filter map[string]interface{}, limit *int, page *int) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented: Users - users"))
}

// Task is the resolver for the task field.
func (r *queryResolver) Task(ctx context.Context, filter map[string]interface{}) (*model.Task, error) {
	panic(fmt.Errorf("not implemented: Task - task"))
}

// Tasks is the resolver for the tasks field.
func (r *queryResolver) Tasks(ctx context.Context, filter map[string]interface{}, limit *int, page *int) ([]*model.Task, error) {
	panic(fmt.Errorf("not implemented: Tasks - tasks"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
