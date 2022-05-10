package router

import (
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"github.com/rs/zerolog/log"
)

var Connector *gorm.DB

func Connect(dsn string) error {
	var err error
	Connector, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		log.Fatal().Msgf("Error connecting to MySQL server: %s", err)
	}
	log.Debug().Msg("Connected to MySQL")
	return nil
}