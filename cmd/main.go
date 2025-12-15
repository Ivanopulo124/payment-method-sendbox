package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"payment-sendbox/service/internal/service"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	service.RegisterRoutes(router)

	srv := &http.Server{
		Addr:    "127.0.0.1" + ":" + "8080",
		Handler: router,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
}
