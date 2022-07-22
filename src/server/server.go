package server

import (
	"context"

	"github.com/jfxdev/petsaur/src/infra/db"
	"github.com/jfxdev/petsaur/src/routes"
	v1 "github.com/jfxdev/petsaur/src/routes/v1"

	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func Serve() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	err := db.Setup()
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	router := gin.Default()

	routes.Register(router)
	v1.Register(router)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	// Inicializa o serviço dentro de uma goroutine, então o graceful shutdown não será bloqueado
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Escuta um sinal de interrupção
	<-ctx.Done()

	// Restaura o comportamento de interrupção e notifica sobre o desligamento
	stop()
	log.Println("shutting down gracefully, press Ctrl+C again to force")

	// O contexto é usado para informar que precisa de 5 sengundos para finalizar os serviços
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")
}
