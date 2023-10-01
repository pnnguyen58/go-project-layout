package configs

import (
	"github.com/pnnguyen58/go-project-layout/pkg/persistence"
)

func LoadDatabaseConfig() *persistence.DatabaseConfig {
	return &persistence.DatabaseConfig{
		Host:                  C.Database.Host,
		Port:                  C.Database.Port,
		User:                  C.Database.User,
		Password:              C.Database.Password,
		DBName:                C.Database.DBName,
		Schema:                C.Database.Schema,
		ConnMaxLifetimeSecond: C.Database.ConnMaxLifetimeSecond,
		MaxOpenConn:           C.Database.MaxOpenConn,
		MaxIdleConn:           C.Database.MaxIdleConn,
	}
}
