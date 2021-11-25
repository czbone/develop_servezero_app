package server

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"time"
	"web/config"
	"web/modules/log"

	"github.com/gin-gonic/gin"
)

func Run(router *gin.Engine) {
	log.Info("Start http server listening: " + config.GetEnv().ServerPort)

	srv := &http.Server{
		Addr:    ":" + config.GetEnv().ServerPort,
		Handler: router,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 10 seconds.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Info("Server exiting")

	pid := fmt.Sprintf("%d", os.Getpid())
	_, openErr := os.OpenFile("pid", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if openErr == nil {
		_ = ioutil.WriteFile("pid", []byte(pid), 0)
	}
}
