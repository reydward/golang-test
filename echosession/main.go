package main

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/rbcervilla/redisstore/v8"
)

var ctx = context.Background()

// Get redis address, password and database from environment variables
func CreateRedisClient() (*redis.Client, error) {

	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       4,
	})

	err := redisClient.Ping(ctx).Err()

	if err != nil {
		log.Info("redis client not created")
		return nil, err
	}

	return redisClient, nil
}

func main() {
	serv := echo.New()

	redisClient, err := CreateRedisClient()

	redisStore, err := redisstore.NewRedisStore(context.Background(), redisClient)
	if err != nil {
		panic(err)
	}

	serv.Use(session.MiddlewareWithConfig(session.Config{
		Store: redisStore,
	}))

	serv.GET("/", func(ctx echo.Context) error {
		sess, _ := session.Get("session", ctx)
		sess.Options = &sessions.Options{
			Path:     "/",
			MaxAge:   86400 * 7,
			HttpOnly: true,
		}

		var count int
		v := sess.Values["count"]
		if v == nil {
			count = 0
		} else {
			count = v.(int)
			count += 1
		}
		sess.Values["count"] = count
		sess.Save(ctx.Request(), ctx.Response())
		ctx.JSON(200, map[string]interface{}{
			"visit": count,
		})
		return nil
	})
	serv.Start(":8081")
}
