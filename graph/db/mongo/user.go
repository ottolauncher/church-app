package db

import (
	"context"
	"fmt"
	"time"

	"github.com/ottolauncher/church-app/graph/model"
	"github.com/ottolauncher/church-app/preloads"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func NewUserManager(d *mongo.Database) *UserManager {
	users := d.Collection("users")
	return &UserManager{Col: users}
}

func (um *UserManager) Create(ctx context.Context, args model.User) (*model.User, error) {
	l, cancel := context.WithTimeout(ctx, 350*time.Millisecond)
	defer cancel()
	now := time.Now().Unix()

	user := model.User{
		Name:      args.Name,
		Email:     args.Email,
		Phone:     args.Phone,
		CreatedAt: now,
	}
	res, err := um.Col.InsertOne(l, user)
	if err != nil {
		return nil, err
	}
	user.ID = res.InsertedID.(primitive.ObjectID)
	return &user, nil
}

func (um *UserManager) Update(ctx context.Context, args model.User) (*model.User, error) {
	l, cancel := context.WithTimeout(ctx, 350*time.Millisecond)
	defer cancel()
	now := time.Now().Unix()

	user := model.User{
		Name:      args.Name,
		Email:     args.Email,
		Phone:     args.Phone,
		UpdatedAt: now,
	}
	res, err := um.Col.UpdateByID(l, args.ID, user)
	if err != nil {
		return nil, err
	}
	user.ID = res.UpsertedID.(primitive.ObjectID)
	return &user, nil
}

func (um *UserManager) Delete(ctx context.Context, filter map[string]interface{}) (bool, error) {
	l, cancel := context.WithTimeout(ctx, 350*time.Millisecond)
	defer cancel()
	if value, ok := filter["id"]; ok {
		pk, err := primitive.ObjectIDFromHex(fmt.Sprintf("%s", value))
		if err != nil {
			return false, err
		}
		_, err = um.Col.DeleteOne(l, pk)
		if err != nil {
			return false, err
		}
	}
	return true, nil
}

func (um *UserManager) Get(ctx context.Context, filter map[string]interface{}) (*model.User, error) {
	load := preloads.GetPreloads(ctx)
	projections := primitive.M{}

	for _, p := range load {
		projections[p] = 1
	}
	opts := options.FindOne().SetProjection(projections)
	l, cancel := context.WithTimeout(ctx, 500*time.Millisecond)
	defer cancel()

	var user model.User
	err := um.Col.FindOne(l, filter, opts).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (um *UserManager) All(ctx context.Context, filter map[string]interface{}, limit int, page int) ([]*model.User, error) {
	l, cancel := context.WithTimeout(ctx, 2000*time.Millisecond)
	defer cancel()
	load := preloads.GetPreloads(ctx)
	projections := primitive.M{}
	for _, p := range load {
		projections[p] = 1
	}
	opts := options.FindOptions{
		Projection: projections,
	}
	opts.SetLimit(int64(limit))

	var users []*model.User
	cur, err := um.Col.Find(l, filter, &opts)

	if err != nil {
		return nil, err
	}
	if err := cur.All(l, &users); err != nil {
		return nil, err
	}

	if err := cur.Err(); err != nil {
		return users, nil
	}
	_ = cur.Close(l)
	if len(users) == 0 {
		return users, mongo.ErrNoDocuments
	}
	return users, nil
}

func (um *UserManager) Search(ctx context.Context, query string, filter map[string]interface{}, limit int, page int) ([]*model.User, error) {
	l, cancel := context.WithTimeout(ctx, 1000*time.Millisecond)
	defer cancel()
	load := preloads.GetPreloads(ctx)
	projections := primitive.M{}
	for _, p := range load {
		projections[p] = 1
	}
	opts := options.FindOptions{
		Projection: projections,
	}

	search := bson.M{
		"$text": bson.M{
			"$search": query,
		},
	}

	var users []*model.User
	cur, err := um.Col.Find(l, search, &opts)

	if err != nil {
		return nil, err
	}
	if err := cur.All(l, &users); err != nil {
		return nil, err
	}

	if err := cur.Err(); err != nil {
		return users, nil
	}
	_ = cur.Close(l)
	if len(users) == 0 {
		return users, mongo.ErrNoDocuments
	}
	return users, nil
}
