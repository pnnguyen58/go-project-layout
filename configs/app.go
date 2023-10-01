package configs

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
)

type config struct {
	Server struct {
		Name           string `mapstructure:"NAME"`
		GRPCPort       string `mapstructure:"GRPC_PORT"`
		HTTPPort       string `mapstructure:"HTTP_PORT"`
		TempoHost      string `mapstructure:"TEMPO_HOST"`
		TempoNameSpace string `mapstructure:"TEMPO_NAMESPACE"`
	} `mapstructure:"SERVER"`
	Database struct {
		Host                  string `mapstructure:"HOST"`
		Port                  int    `mapstructure:"POST"`
		User                  string `mapstructure:"USER"`
		Password              string `mapstructure:"PASSWORD"`
		DBName                string `mapstructure:"DATABASE_NAME"`
		Schema                string `mapstructure:"SCHEMA"`
		ConnMaxLifetimeSecond int    `mapstructure:"CONN_MAX_LIFETIME_SECOND"`
		MaxOpenConn           int    `mapstructure:"MAX_OPEN_CONN"`
		MaxIdleConn           int    `mapstructure:"MAX_IDLE_CONN"`
	} `mapstructure:"DATABASE"`
}

var C config

func ReadConfig() {
	Config := &C
	Config.Server.Name = GetEnv[string]("SERVER_NAME", "aspire-code-challenge")
	Config.Server.GRPCPort = GetEnv[string]("SERVER_GRPC_PORT", "8001")
	Config.Server.HTTPPort = GetEnv[string]("SERVER_HTTP_PORT", "9001")
	Config.Server.TempoHost = GetEnv[string]("TEMPO_HOST", "temporal:7233")
	Config.Server.TempoNameSpace = GetEnv[string]("TEMPO_NAMESPACE", "aspire-code-challenge")

	Config.Database.Host = GetEnv[string]("DB_HOST", "aspire-db")
	Config.Database.Port = GetEnv[int]("DB_PORT", 5432)
	Config.Database.User = GetEnv[string]("DB_USER", "aspire")
	Config.Database.Password = GetEnv[string]("DB_PASSWORD", "1qazxsw23edc")
	Config.Database.DBName = GetEnv[string]("DB_NAME", "code_challenge")
	Config.Database.Schema = GetEnv[string]("DB_SCHEMA", "public")
	Config.Database.ConnMaxLifetimeSecond = GetEnv[int]("DB_CONN_MAX_LIFETIME_SECOND", 300)
	Config.Database.MaxOpenConn = GetEnv[int]("DB_MAX_OPEN_CONN", 100)
	Config.Database.MaxIdleConn = GetEnv[int]("DB_MAX_IDLE_CONN", 100)
}

func rootDir() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	return filepath.Dir(d)
}

type EnvType interface {
	string | int
}

func GetEnv[V EnvType](key string, fallback V) V {
	if value, ok := os.LookupEnv(key); ok {
		var convertVal V
		log.Println(value)
		_, err := fmt.Sscanf(value, "%v", &convertVal)
		if err == nil {
			return convertVal
		}
	}
	return fallback
}

func readConf() {
	Config := &C

	viper.SetConfigName(".env")
	viper.SetConfigType("yml")
	viper.AddConfigPath(filepath.Join(rootDir(), "config"))
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Println(".env not found")
	}
	if err := viper.Unmarshal(&Config); err != nil {
		log.Fatalln(err)
	}
}
