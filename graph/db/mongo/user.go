package db

import (
	"context"

	"github.com/ottolauncher/church-app/graph/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type IUser interface {
	Create(ctx context.Context, args model.User) (*model.User, error)
	Update(ctx context.Context, args model.User) (*model.User, error)
	Delete(ctx context.Context, filter map[string]interface{}) (bool, error)
	Get(ctx context.Context, filter map[string]interface{}) (*model.User, error)
	All(ctx context.Context, filter map[string]interface{}, limit int, page int) ([]*model.User, error)
	Search(ctx context.Context, query string, filter map[string]interface{}, limit int, page int) ([]*model.User, error)
}

type UserManager struct {
	Col *mongo.Collection
}

func (um *UserManager) Create(ctx context.Context, args model.User) (*model.User, error) {
	panic("not implemented") // TODO: Implement
}

func (um *UserManager) Update(ctx context.Context, args model.User) (*model.User, error) {
	panic("not implemented") // TODO: Implement
}

func (um *UserManager) Delete(ctx context.Context, filter map[string]interface{}) (bool, error) {
	panic("not implemented") // TODO: Implement
}

func (um *UserManager) Get(ctx context.Context, filter map[string]interface{}) (*model.User, error) {
	panic("not implemented") // TODO: Implement
}

func (um *UserManager) All(ctx context.Context, filter map[string]interface{}, limit int, page int) ([]*model.User, error) {
	panic("not implemented") // TODO: Implement
}

func (um *UserManager) Search(ctx context.Context, query string, filter map[string]interface{}, limit int, page int) ([]*model.User, error) {
	panic("not implemented") // TODO: Implement
}
