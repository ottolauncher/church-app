package main

import (
	"context"
	"log"
	"os"
	"sync"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/ottolauncher/church-app/graph"
	db "github.com/ottolauncher/church-app/graph/db/mongo"
	"github.com/ottolauncher/church-app/graph/generated"
	"go.mongodb.org/mongo-driver/mongo"
)

const defaultPort = "8080"

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	var (
		once sync.Once
		dao  *mongo.Client
	)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	once.Do(func() {
		dao = db.Init()
	})

	src := dao.Database("churchApp")
	defer func() {
		if err := dao.Disconnect(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()

	um := db.NewUserManager(src)
	tm := db.NewTaskManager(src)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{UM: um, TM: tm}}))

	e.POST("/query", func(c echo.Context) error {
		srv.ServeHTTP(c.Response(), c.Request())
		return nil
	})
	e.GET("/playground", func(c echo.Context) error {
		playground.Handler("GraphQL playground", "/query").ServeHTTP(c.Response(), c.Request())
		return nil
	})

	err := e.Start(":" + port)
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	if err != nil {
		log.Fatalln(err)
	}

}
