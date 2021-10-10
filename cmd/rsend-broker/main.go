package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"

	"github.com/ishansd94/sample-app/internal/app/healthz"
	"github.com/ishansd94/sample-app/internal/app/rsend_broker"
	"github.com/ishansd94/sample-app/internal/app/sample"
	"github.com/ishansd94/sample-app/internal/app/version"
	"github.com/ishansd94/sample-app/pkg/env"
	"github.com/ishansd94/sample-app/pkg/log"
	"github.com/ishansd94/sample-app/pkg/router"
)

func main() {

	log.Info("main", fmt.Sprintf("starting service... commit: %s, build time: %s, release: %s", version.Commit, version.BuildTime, version.Release))

	gin.SetMode(env.Get("GIN_MODE", "debug"))

	var ctx = context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "bhm3dfw74kdaxsakoyww-redis.services.clever-cloud.com:3434",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	clientId := "abcd"
	token := "ojnhgy5ngksuut"

	err := rdb.Set(ctx, fmt.Sprintf("client_%s", clientId), token, time.Hour * 4).Err()
	if err != nil {
		log.Error("Redis", "Connection Error", err)
	}


	apiServer := apiServer()
	apiServer.Start()

	wsServer := wsServer()
	wsServer.Start()

	router.Listen()
}

func apiServer() *router.Handler {

	defaultPort := fmt.Sprintf(":%s", env.Get(rsend_broker.RSEND_BROKER_API_PORT, rsend_broker.DEFULT_RSEND_BROKER_API_PORT))

	r := gin.Default()
	r.Use(cors.Default())

	// api server private endpoints
	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("sample", sample.Get)
		apiv1.POST("sample", sample.Create)
	}

	// api server public endpoints
	r.GET("/version", version.Get)

	r.GET("/healthz", healthz.Health)

	serverConfig := &http.Server{
		Addr:         defaultPort,
		Handler:      r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	server := router.NewRouter("rsend-broker-api", serverConfig)

	return server
}

func wsServer() *router.Handler {

	defaultPort := fmt.Sprintf(":%s", env.Get(rsend_broker.RSEND_BROKER_WS_PORT, rsend_broker.DEFULT_RSEND_BROKER_WS_PORT))

	r := gin.Default()
	r.Use(cors.Default())



	// api server private endpoints
	apiv1 := r.Group("/v1")
	{
		apiv1.GET("ws",  rsend_broker.Get)
		//apiv1.POST("ws", rsend_broker.Post)
	}

	// api server public endpoints
	r.GET("/version", version.Get)

	r.GET("/healthz", healthz.Health)

	serverConfig := &http.Server{
		Addr:         defaultPort,
		Handler:      r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	server := router.NewRouter("rsend-broker-ws", serverConfig)

	return server
}
