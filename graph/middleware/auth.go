package middleware

import (
	"github.com/labstack/echo"
	db "github.com/ottolauncher/church-app/graph/db/mongo"
	"github.com/ottolauncher/church-app/graph/model"
	"golang.org/x/net/context"
)

var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

func AuthMiddleware(um *db.UserManager) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			arg, err := c.Cookie("auth-cookie")
			if err != nil || arg == nil {
				return err
			}
			user, err := um.Get(context.TODO(), map[string]interface{}{"id": c})
			if err != nil {
				c.Error(err)
			}

			c.Set("user", user)
			return next(c)
		}
	}
}

// func AuthMiddleware(um *db.UserManager) func(http.Handler) http.Handler {
// 	return func(next http.Handler) http.Handler {
// 		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 			c, err := r.Cookie("auth-cookie")
// 			if err != nil || c == nil {
// 				next.ServeHTTP(w, r)
// 				return
// 			}
// 			user, err := um.Get(context.TODO(), map[string]interface{}{"id": c})
// 			if err != nil {
// 				http.Error(w, "Invalid cookie", http.StatusForbidden)
// 				return
// 			}
// 			ctx := context.WithValue(r.Context(), userCtxKey, user)

// 			r = r.WithContext(ctx)
// 			next.ServeHTTP(w, r)
// 		})
// 	}
// }

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func ForContext(ctx context.Context) *model.User {
	raw, _ := ctx.Value(userCtxKey).(*model.User)
	return raw
}
