//go:generate go run github.com/99designs/gqlgen generate
package graph

import db "github.com/ottolauncher/church-app/graph/db/mongo"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	TM *db.TaskManager
	UM *db.UserManager
}
