package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"github.com/spf13/cobra"

	"github.com/mabou-dev/music-library/internal/album"
	"github.com/mabou-dev/music-library/pkg/log"
)

var StartCmd = NewStartCmd()

func NewStartCmd() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "start",
		Short: "start server",
		Run:   StartExecute,
	}

	return cmd
}

func StartExecute(cmd *cobra.Command, args []string) {
	log.Setup()
	defer log.Sync()

	router := gin.Default()
        router.Use(
            cors.New(cors.Config{
                AllowOrigins: []string{"*"},
                AllowMethods: []string{"GET", "POST"},
                AllowHeaders: []string{"Content-Type"},
                AllowCredentials: true,
                MaxAge: 12*time.Hour,
        }))

	repo := album.NewAlbumRepository()
	service := album.NewAlbumService(repo)
	handler := album.NewAlbumHandler(service)

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	router.GET("/albums", handler.GetAlbums)

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM)
	defer stop()

	go func() {
		log.Infof("starting server on port %s", server.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Errorf("server err : %w", err)
		}
	}()

	<-ctx.Done()
	log.Info("shutting done server")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Errorf("server forced to shutdown: %w", err)
	}

	log.Info("server exited gracefully")
}
