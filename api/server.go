package serve

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/pius706975/the-sims-backend/api/routes"
	envConfig "github.com/pius706975/the-sims-backend/config"
	"github.com/pius706975/the-sims-backend/package/database"
	"github.com/pius706975/the-sims-backend/package/utils"

	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
	"github.com/spf13/cobra"
)

var ServeCMD = &cobra.Command{
	Use:   "serve",
	Short: "For Running api server",
	RunE:  serve,
}

func corsHandler(allowedOrigins []string) *cors.Cors {
	return cors.New(cors.Options{
		AllowedOrigins: allowedOrigins,
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodOptions,
		},
		AllowedHeaders: []string{
			"Authorization",
			"Content-Type",
			"Accept",
			"Origin",
			"X-Requested-With",
		},
		ExposedHeaders: []string{
			"Content-Length",
		},
		AllowCredentials: false,
	})
}

func serve(cmd *cobra.Command, args []string) error {
	envCfg := envConfig.LoadConfig()

	// Set Gin mode BEFORE gin.Default()
	switch envCfg.Mode {
	case "production":
		gin.SetMode(gin.ReleaseMode)
	case "staging":
		gin.SetMode(gin.TestMode)
	default:
		gin.SetMode(gin.DebugMode)
	}

	errorLogger, debugLogger := utils.InitLogger()
	debugLogger.Println("Starting server...")

	// Init DB
	db, err := database.NewDB()
	if err != nil {
		errorLogger.Println("DB connection failed:", err)
		return err
	}

	// Router
	router := gin.Default()

	if err := routes.RouteApp(router, db); err != nil {
		errorLogger.Println("Failed to initialize route:", err)
		return err
	}

	// CORS handler
	handler := corsHandler(envCfg.AllowedOrigins).Handler(router)

	// HTTP Server
	srv := &http.Server{
		Addr:    ":" + envCfg.Port,
		Handler: handler,
	}

	// Run server in goroutine
	go func() {
		debugLogger.Printf("Server running on port %s", envCfg.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			errorLogger.Println("Server error:", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	debugLogger.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		errorLogger.Println("Server forced to shutdown:", err)
		return err
	}

	debugLogger.Println("Server exited cleanly")
	return nil
}
