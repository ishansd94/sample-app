package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/ishansd94/sample-app/internal/pkg/router"

	"github.com/ishansd94/sample-app/internal/app/sample"
	"github.com/ishansd94/sample-app/internal/pkg/version"
	"github.com/ishansd94/sample-app/pkg/env"
	"github.com/ishansd94/sample-app/pkg/log"
)

func main() {

	log.Info("main", fmt.Sprintf("starting service... commit: %s, build time: %s, release: %s", version.Commit, version.BuildTime, version.Release))

	r1 := gin.Default()

	r1.Use(cors.Default())

	apiv1 := r1.Group("/api/v1")
	{
		apiv1.GET("", sample.Get)
		apiv1.POST("", sample.Create)
	}

	r1.GET("/version", version.Get)


	r2 := gin.Default()

	r2.GET("/metrics", func(context *gin.Context) {
		context.JSON(http.StatusOK, "OK")
	})


	server1 := &http.Server{
		Addr:         fmt.Sprintf(":%s", env.Get("PORT", "8000")),
		Handler:      r1,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	server2 := &http.Server{
		Addr:         ":9000",
		Handler:      r2,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}


	serv1 := router.NewRouter("app-server", server1)
	serv1.Start()

	serv2 := router.NewRouter("metrics-server", server2)
	serv2.Start()


	router.Wait()

}
