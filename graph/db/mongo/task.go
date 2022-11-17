package db

import (
	"context"
	"fmt"
	"time"

	"github.com/ottolauncher/church-app/graph/model"
	"github.com/ottolauncher/church-app/preloads"
	"github.com/ottolauncher/church-app/utils/text"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ITask interface {
	Create(ctx context.Context, args model.Task) (*model.Task, error)
	Update(ctx context.Context, args model.Task) (*model.Task, error)
	Delete(ctx context.Context, filter map[string]interface{}) (bool, error)
	Get(ctx context.Context, filter map[string]interface{}) (*model.Task, error)
	All(ctx context.Context, filter map[string]interface{}, limit int, page int) ([]*model.Task, error)
	Search(ctx context.Context, query string, filter map[string]interface{}, limit int, page int) ([]*model.Task, error)
}

type TaskManager struct {
	Col *mongo.Collection
}

func NewTaskManager(d *mongo.Database) *TaskManager {
	tasks := d.Collection("tasks")
	return &TaskManager{Col: tasks}
}

func (tm *TaskManager) Create(ctx context.Context, args model.Task) (*model.Task, error) {
	l, cancel := context.WithTimeout(ctx, 350*time.Millisecond)
	defer cancel()
	slug := text.Slugify(args.Title)
	now := time.Now().Unix()

	task := model.Task{
		Title:     args.Title,
		Note:      args.Title,
		Slug:      slug,
		CreatedAt: now,
	}
	res, err := tm.Col.InsertOne(l, task)
	if err != nil {
		return nil, err
	}
	task.ID = res.InsertedID.(primitive.ObjectID).Hex()
	return &task, nil
}

func (tm *TaskManager) Update(ctx context.Context, args model.Task) (*model.Task, error) {
	l, cancel := context.WithTimeout(ctx, 350*time.Millisecond)
	defer cancel()
	slug := text.Slugify(args.Title)
	now := time.Now().Unix()

	task := model.Task{
		Title:     args.Title,
		Note:      args.Title,
		Slug:      slug,
		UpdatedAt: now,
	}
	res, err := tm.Col.UpdateByID(l, args.ID, task)
	if err != nil {
		return nil, err
	}
	task.ID = res.UpsertedID.(primitive.ObjectID).Hex()
	return &task, nil
}

func (tm *TaskManager) Delete(ctx context.Context, filter map[string]interface{}) (bool, error) {
	l, cancel := context.WithTimeout(ctx, 350*time.Millisecond)
	defer cancel()
	if value, ok := filter["id"]; ok {
		pk, err := primitive.ObjectIDFromHex(fmt.Sprintf("%s", value))
		if err != nil {
			return false, err
		}
		_, err = tm.Col.DeleteOne(l, pk)
		if err != nil {
			return false, err
		}
	}
	return true, nil
}

func (tm *TaskManager) Get(ctx context.Context, filter map[string]interface{}) (*model.Task, error) {
	load := preloads.GetPreloads(ctx)
	projections := primitive.M{}

	for _, p := range load {
		projections[p] = 1
	}
	opts := options.FindOne().SetProjection(projections)
	l, cancel := context.WithTimeout(ctx, 500*time.Millisecond)
	defer cancel()

	var task model.Task
	err := tm.Col.FindOne(l, filter, opts).Decode(&task)
	if err != nil {
		return nil, err
	}
	return &task, nil

}

func (tm *TaskManager) All(ctx context.Context, filter map[string]interface{}, limit int, page int) ([]*model.Task, error) {
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

	var tasks []*model.Task
	cur, err := tm.Col.Find(l, filter, &opts)

	if err != nil {
		return nil, err
	}
	if err := cur.All(l, &tasks); err != nil {
		return nil, err
	}

	if err := cur.Err(); err != nil {
		return tasks, nil
	}
	_ = cur.Close(l)
	if len(tasks) == 0 {
		return tasks, mongo.ErrNoDocuments
	}
	return tasks, nil
}

func (tm *TaskManager) Search(ctx context.Context, query string, filter map[string]interface{}, limit int, page int) ([]*model.Task, error) {
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

	var tasks []*model.Task
	cur, err := tm.Col.Find(l, search, &opts)

	if err != nil {
		return nil, err
	}
	if err := cur.All(l, &tasks); err != nil {
		return nil, err
	}

	if err := cur.Err(); err != nil {
		return tasks, nil
	}
	_ = cur.Close(l)
	if len(tasks) == 0 {
		return tasks, mongo.ErrNoDocuments
	}
	return tasks, nil
}
