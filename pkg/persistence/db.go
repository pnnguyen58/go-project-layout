package persistence

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

type DatabaseConfig struct {
	Host                  string `json:"host"`
	Port                  int    `json:"port"`
	User                  string `json:"user"`
	Password              string `json:"password"`
	DBName                string `json:"dbName"`
	Schema                string `json:"schema"`
	ConnMaxLifetimeSecond int    `json:"connMaxLifetimeSecond"`
	MaxOpenConn           int    `json:"maxOpenConn"`
	MaxIdleConn           int    `json:"maxIdleConn"`
}

const (
	psqlConn = "host=%v user=%v dbname=%v sslmode=disable password=%v port=%v search_path=%v"
)

func NewPostgresDB(conf *DatabaseConfig) (*gorm.DB, error) {
	strConn := fmt.Sprintf(psqlConn, conf.Host, conf.User,
		conf.DBName, conf.Password, conf.Port, conf.Schema)
	db, err := gorm.Open(postgres.Open(strConn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err == nil {
		if sqlDB, e := db.DB(); e == nil {
			sqlDB.SetMaxIdleConns(conf.MaxIdleConn)
			sqlDB.SetMaxOpenConns(conf.MaxOpenConn)
			sqlDB.SetConnMaxLifetime(time.Duration(conf.ConnMaxLifetimeSecond) * time.Second)
		}
	}
	return db, err
}
