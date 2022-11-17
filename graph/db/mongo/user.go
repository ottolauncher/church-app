package db

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/ottolauncher/church-app/graph/model"
	"github.com/ottolauncher/church-app/preloads"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

type IUser interface {
	Login(c echo.Context) error
	Register(c echo.Context) error
	Create(ctx context.Context, args model.NewUser) (*model.User, error)
	Update(ctx context.Context, args model.UpdateUser) (*model.User, error)
	Delete(ctx context.Context, filter map[string]interface{}) error
	Get(ctx context.Context, filter map[string]interface{}) (*model.User, error)
	All(ctx context.Context, filter map[string]interface{}, limit int, page int) ([]*model.User, error)
	Search(ctx context.Context, query string, filter map[string]interface{}, limit int, page int) ([]*model.User, error)
}

var SecretKey []byte = []byte("DOi4UCP0UQ6jmEbUvVWNCSJTGfM8/LE5UH4XRK3QW92kArVKnGxmwexgQUoXi6Bpido=")

type JWTCustomClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
type UserManager struct {
	Col *mongo.Collection
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(password string, hash []byte) bool {
	err := bcrypt.CompareHashAndPassword(hash, []byte(password))
	return err == nil
}

func NewUserManager(d *mongo.Database) *UserManager {
	users := d.Collection("users")
	return &UserManager{Col: users}
}

func (um *UserManager) Login(c echo.Context) error {
	l, cancel := context.WithTimeout(c.Request().Context(), 350*time.Millisecond)
	defer cancel()

	username := c.FormValue("username")
	password := c.FormValue("password")

	var user model.User
	err := um.Col.FindOne(l, map[string]interface{}{"username": username}).Decode(&user)

	if err != nil {
		return err
	}

	if CheckPasswordHash(password, user.Password) {
		claims := &JWTCustomClaims{
			user.Name,
			jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			},
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		t, err := token.SignedString(SecretKey)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, echo.Map{"token": t})
	}
	return errors.New("invalid credentials")
}

func (um *UserManager) Register(c echo.Context) error {
	l, cancel := context.WithTimeout(c.Request().Context(), 350*time.Millisecond)
	defer cancel()

	username := c.FormValue("username")
	name := c.FormValue("name")
	email := c.FormValue("email")
	phone := c.FormValue("phone")
	password1 := c.FormValue("password1")
	password2 := c.FormValue("password2")

	if password1 == password2 {
		now := time.Now().Unix()
		pwd, _ := HashPassword(string(password2))
		usr := model.User{
			Name:      name,
			Username:  username,
			Email:     email,
			Phone:     phone,
			Password:  []byte(pwd),
			CreatedAt: now,
		}
		_, err := um.Col.InsertOne(l, usr)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, echo.Map{"status": "OK!", "message": "login address: http://localhost:8080/login"})
	}

	return errors.New("invalid credentials")
}

func (um *UserManager) Create(ctx context.Context, args model.NewUser) (*model.User, error) {
	l, cancel := context.WithTimeout(ctx, 350*time.Millisecond)
	defer cancel()
	now := time.Now().Unix()

	if args.Password1 == args.Password2 {
		pwd, _ := HashPassword(string(args.Password2))
		user := model.User{
			Name:      args.Name,
			Email:     args.Email,
			Phone:     args.Phone,
			Password:  []byte(pwd),
			CreatedAt: now,
		}
		res, err := um.Col.InsertOne(l, user)
		if err != nil {
			return nil, err
		}
		user.ID = res.InsertedID.(primitive.ObjectID).Hex()
		return &user, nil
	}
	return nil, errors.New("password missmatch")
}

func (um *UserManager) Update(ctx context.Context, args model.UpdateUser) (*model.User, error) {
	l, cancel := context.WithTimeout(ctx, 500*time.Millisecond)
	defer cancel()
	now := time.Now().Unix()

	user, err := um.Get(ctx, map[string]interface{}{"email": args.Email})
	if err != nil {
		return nil, err
	}
	if CheckPasswordHash(*args.OldPassword, user.Password) {
		pwd, _ := HashPassword(*args.NewPassword)
		usr := model.User{
			Name:      args.Name,
			Email:     args.Email,
			Password:  []byte(pwd),
			CreatedAt: user.CreatedAt,
			UpdatedAt: now,
		}
		res, err := um.Col.UpdateByID(l, user.ID, usr)
		if err != nil {
			return nil, err
		}
		usr.ID = res.UpsertedID.(primitive.ObjectID).Hex()
		return &usr, nil
	}
	return nil, errors.New("password missmatch")

}

func (um *UserManager) Delete(ctx context.Context, filter map[string]interface{}) error {
	l, cancel := context.WithTimeout(ctx, 350*time.Millisecond)
	defer cancel()
	if value, ok := filter["id"]; ok {
		pk, err := primitive.ObjectIDFromHex(fmt.Sprintf("%s", value))
		if err != nil {
			return err
		}
		_, err = um.Col.DeleteOne(l, pk)
		if err != nil {
			return err
		}
	}
	return nil
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
