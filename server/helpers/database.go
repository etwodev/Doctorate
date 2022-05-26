package helpers

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/zerolog/log"
	"xorm.io/xorm"
	"xorm.io/xorm/names"
)

var Engine 		*xorm.Engine

func Connect(dsn string) error {
	var err error
	Engine, err = xorm.NewEngine("mysql", dsn)
	Engine.SetMapper(names.SameMapper{})
	if err != nil {
		log.Fatal().Err(err).Msg("Startup failed!")
	}
	return nil
}
