package main

import (
	"context"
	"go.uber.org/fx"
	"os"

	"github.com/pnnguyen58/go-project-layout/configs"
	"github.com/pnnguyen58/go-project-layout/internal/adapters/controllers"
	"github.com/pnnguyen58/go-project-layout/internal/adapters/repositories"
	portRepository "github.com/pnnguyen58/go-project-layout/internal/core/ports/repositories"
	"github.com/pnnguyen58/go-project-layout/pkg/clients"
	"github.com/pnnguyen58/go-project-layout/pkg/logger"
	"github.com/pnnguyen58/go-project-layout/pkg/persistence"
)

func main() {
	configs.ReadConfig()
	ctx := context.Background()
	ap := fx.New(
		fx.Provide(
			logger.New,
			configs.LoadTempoConfig,
			configs.LoadDatabaseConfig,
			clients.NewTemporal,
			persistence.NewPostgresDB,
			context.TODO,
			// TODO add all providers

			repositories.NewLoan,
			repositories.NewRepayment,
			controllers.NewLoan,
		),
		fx.Invoke(
			portRepository.Wire,
			registerLoan,
		),
	)
	if err := ap.Start(ctx); err != nil {
		os.Exit(1)
	}
	<-ctx.Done()
}
