package config

import (
	"log"

	"github.com/joeshaw/envdecode"
)

type Conf struct {
	Server  serverConf
	DB      dbConf
	Setting settingConf
}

type serverConf struct {
	Port      string `env:"STOCKYARD_APP_PORT,required"`
	IP        string `env:"STOCKYARD_APP_IP,required"`
	SecretKey []byte `env:"STOCKYARD_APP_SECRET_KEY,required"`
}

type dbConf struct {
	Host              string `env:"STOCKYARD_DB_HOST,required"`
	Port              string `env:"STOCKYARD_DB_PORT,required"`
	User              string `env:"STOCKYARD_DB_USER,required"`
	Password          string `env:"STOCKYARD_DB_PASSWORD,required"`
	DBName            string `env:"STOCKYARD_DB_NAME,required"`
	HasAutoMigrations bool   `env:"STOCKYARD_DB_HAS_AUTO_MIGRATIONS,default=true"`
}

type settingConf struct {
	HasDebugging bool `env:"STOCKYARD_HAS_DEBUGGING,default=true"`
	HasAnalyzer  bool `env:"STOCKYARD_HAS_ANALYZER,default=true"`
}

func AppConfig() *Conf {
	var c Conf
	if err := envdecode.StrictDecode(&c); err != nil {
		log.Fatalf("Failed to decode: %s", err)
	}
	return &c
}
