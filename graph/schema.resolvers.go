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
	task := model.Task{
		Title: input.Title,
		Note:  input.Note,
	}
	res, err := r.TM.Create(ctx, task)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// UpdateTask is the resolver for the updateTask field.
func (r *mutationResolver) UpdateTask(ctx context.Context, input model.UpdateTask) (*model.Task, error) {
	task := model.Task{
		ID:    input.ID,
		Title: input.Title,
		Note:  input.Note,
	}
	res, err := r.TM.Update(ctx, task)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// DeleteTask is the resolver for the deleteTask field.
func (r *mutationResolver) DeleteTask(ctx context.Context, id string) (bool, error) {
	if err := r.TM.Delete(ctx, map[string]interface{}{"id": id}); err != nil {
		return false, err
	}
	return true, nil
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, filter map[string]interface{}) (*model.User, error) {
	res, err := r.UM.Get(ctx, filter)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context, filter map[string]interface{}, limit *int, page *int) ([]*model.User, error) {
	res, err := r.UM.All(ctx, filter, *limit, *page)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Task is the resolver for the task field.
func (r *queryResolver) Task(ctx context.Context, filter map[string]interface{}) (*model.Task, error) {
	res, err := r.TM.Get(ctx, filter)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Tasks is the resolver for the tasks field.
func (r *queryResolver) Tasks(ctx context.Context, filter map[string]interface{}, limit *int, page *int) ([]*model.Task, error) {
	res, err := r.TM.All(ctx, filter, *limit, *page)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
