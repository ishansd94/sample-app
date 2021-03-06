package main

import (
	"fmt"
	"github.com/ishansd94/sample-app/internal/app/healthz"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/ishansd94/sample-app/internal/app/sample"
	"github.com/ishansd94/sample-app/internal/app/version"
	"github.com/ishansd94/sample-app/internal/pkg/metrics"
	"github.com/ishansd94/sample-app/pkg/env"
	"github.com/ishansd94/sample-app/pkg/log"
	"github.com/ishansd94/sample-app/pkg/router"
)

func main() {

	log.Info("main", fmt.Sprintf("starting service... commit: %s, build time: %s, release: %s", version.Commit, version.BuildTime, version.Release))

	gin.SetMode(env.Get("GIN_MODE", "debug"))

	apiserver := apiServer()
	apiserver.Start()

	metricsserver := metricsServer()
	metricsserver.Start()

	router.Listen()

}

func apiServer() *router.Handler {

	defaultPort := fmt.Sprintf(":%s", env.Get("PORT", "8000"))

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

	server := router.NewRouter("app-server", serverConfig)

	return server
}

func metricsServer() *router.Handler {
	r := gin.Default()

	metricsPort := fmt.Sprintf(":%s", env.Get("METRICS_PORT", "9090"))

	r.GET("/metrics", metrics.PrometheusMetrics)

	serverConfig := &http.Server{
		Addr:         metricsPort,
		Handler:      r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	server := router.NewRouter("metrics-server", serverConfig)

	return server
}
