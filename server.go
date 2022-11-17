package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/ottolauncher/church-app/graph"
	db "github.com/ottolauncher/church-app/graph/db/mongo"
	"github.com/ottolauncher/church-app/graph/generated"
	auth "github.com/ottolauncher/church-app/graph/middleware"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
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

	e.Use(auth.AuthMiddleware(um))

	config := generated.Config{Resolvers: &graph.Resolver{UM: um, TM: tm}}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(config))

	e.POST("/login", um.Login)
	e.GET("/playground", func(c echo.Context) error {
		playground.Handler("GraphQL playground", "/query").ServeHTTP(c.Response(), c.Request())
		return nil
	})

	cfg := middleware.JWTConfig{
		Claims:     &model.JWTCustomClaims,
		SigningKey: []byte(model.SecretKey),
	}
	r := e.Group("/query")
	r.Use(middleware.JWTWithConfig(cfg))

	r.POST("/query", func(c echo.Context) error {
		srv.ServeHTTP(c.Response(), c.Request())
		return nil
	})

	h2s := &http2.Server{
		MaxConcurrentStreams: 250,
		MaxReadFrameSize:     1048576,
		IdleTimeout:          10 * time.Second,
	}

	s := http.Server{
		Addr:    ":" + port,
		Handler: h2c.NewHandler(e, h2s),
	}

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal(err)
	}

}
