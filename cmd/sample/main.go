package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-contrib/cache/persistence"

	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/requestid"
	swaggerFiles "github.com/swaggo/files"
	swagger "github.com/swaggo/gin-swagger"

	"github.com/ishansd94/sample-app/internal/app/healthz"
	//"https://github.com/gin-contrib/httpsign"
	"github.com/gin-gonic/gin"

	_ "github.com/ishansd94/sample-app/docs/sample"
	"github.com/ishansd94/sample-app/internal/app/sample"
	"github.com/ishansd94/sample-app/internal/app/version"
	"github.com/ishansd94/sample-app/internal/pkg/metrics"
	"github.com/ishansd94/sample-app/pkg/env"
	"github.com/ishansd94/sample-app/pkg/log"
	"github.com/ishansd94/sample-app/pkg/router"
)

// @title Gin Swagger Example API
// @version 1.0
// @description This is a sample server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8000
// @BasePath /
// @schemes http
func main() {

	log.Info("main", fmt.Sprintf("starting service... commit: %s, build time: %s, release: %s", version.Commit, version.BuildTime, version.Release))

	gin.SetMode(env.Get("GIN_MODE", "debug"))

	apiServer := apiServer()
	apiServer.Start()

	metricServer := metricsServer()
	metricServer.Start()

	docsServer := docsServer()
	docsServer.Start()

	router.Listen()
}

func apiServer() *router.Handler {

	defaultPort := fmt.Sprintf(":%s", env.Get("PORT", "8000"))

	redisStore := persistence.NewRedisCache("bhm3dfw74kdaxsakoyww-redis.services.clever-cloud.com:3434", "9n8IQfM1OvfBvp4HVWo", time.Second)

	r := gin.Default()
	r.Use(cors.Default())
	r.Use(requestid.New())
	// api server private endpoints
	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("sample", sample.Get)
		apiv1.POST("sample", sample.Create)
	}

	apiv2 := r.Group("/api/v2")
	{
		apiv2.GET("/proto", sample.Proto)
	}

	// api server public endpoints
	r.GET("/version", cache.CachePage(redisStore, time.Hour, version.Get))

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

func docsServer() *router.Handler {
	r := gin.Default()

	metricsPort := fmt.Sprintf(":%s", env.Get("SWAGGER_PORT", "3000"))

	url := swagger.URL("http://localhost:3000/swagger/doc.json") // The url pointing to API definition
	r.GET("/*any", swagger.WrapHandler(swaggerFiles.Handler, url))

	serverConfig := &http.Server{
		Addr:         metricsPort,
		Handler:      r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	server := router.NewRouter("docs-server", serverConfig)

	return server
}
