package main

import (
	"context"
	"fmt"
	"go.uber.org/fx"
	"os"
	"os/signal"

	"github.com/pnnguyen58/go-project-layout/configs"
	"github.com/pnnguyen58/go-project-layout/internal/adapters/controllers"
	"github.com/pnnguyen58/go-project-layout/internal/adapters/repositories"
	portRepository "github.com/pnnguyen58/go-project-layout/internal/core/ports/repositories"
	"github.com/pnnguyen58/go-project-layout/internal/core/usecases"
	"github.com/pnnguyen58/go-project-layout/pkg/clients"
	"github.com/pnnguyen58/go-project-layout/pkg/logger"
	"github.com/pnnguyen58/go-project-layout/pkg/persistence"
)

func main() {
	configs.ReadConfig()
	ctx, cancel := context.WithCancel(context.Background())
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan)

	go func() {
		// Wait for termination signal
		<-signalChan
		// Trigger cancellation of the context
		cancel()
		// Wait for goroutine to finish
		fmt.Println("The service terminated gracefully")
	}()

	ap := fx.New(
		fx.Provide(
			logger.New,
			configs.LoadTempoConfig,
			configs.LoadDatabaseConfig,
			clients.NewTemporal,
			persistence.NewPostgresDB,
			context.TODO,
			// TODO add all providers

			usecases.NewLoan,
			repositories.NewLoan,
			repositories.NewRepayment,
			controllers.NewLoan,
		),
		fx.Invoke(
			portRepository.Wire,
			controllers.ListenAndServe,
		),
	)
	if err := ap.Start(ctx); err != nil {
		os.Exit(1)
	}
}
