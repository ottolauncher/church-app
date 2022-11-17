package db

import (
	"context"

	"github.com/ottolauncher/church-app/graph/model"
)

type ITask interface {
	Create(ctx context.Context, args model.Task)
}
