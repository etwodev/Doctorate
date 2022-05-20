package helpers

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/zerolog/log"
	"xorm.io/xorm"
)

var Engine 		*xorm.Engine

func Connect(dsn string) error {
	var err error
	Engine, err = xorm.NewEngine("mysql", dsn)
	if err != nil {
		log.Fatal().Msgf("Error connecting to MySQL server: %s", err)
	}
	log.Debug().Msg("Connected to MySQL")
	return nil
}
