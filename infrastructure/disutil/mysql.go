package disutil

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"log"
)

const accessTokenTemplate = "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local"

type databaseConfig struct {
	User string `envconfig:"DB_USER" default:"user"`
	Pass string `envconfig:"DB_PASSWORD" default:"pass"`
	IP   string `envconfig:"DB_IP" default:"localhost"`
	Port string `envconfig:"DB_PORT" default:"3306"`
	Name string `envconfig:"DB_NAME" default:"app"`
}

func GetMysqlConnectionInfo() string {
	/* ===== データベースへ接続する. ===== */
	var config databaseConfig
	if err := envconfig.Process("", &config); err != nil {
		log.Fatal("Unable to connect to DB(Insufficient variables)")
	}
	return fmt.Sprintf(accessTokenTemplate, config.User, config.Pass, config.IP, config.Port, config.Name)
}
