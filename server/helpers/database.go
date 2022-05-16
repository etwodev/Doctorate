package helpers

import (
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/zerolog/log"
	"xorm.io/xorm"
)

var Engine *xorm.Engine

var HostConfigRAW string

var HostConfigSTR string

var PubKey string

var PrivKey string

func Connect(dsn string) error {
	var err error
	HostConfigRAW = os.Getenv("HOSTRAW")
	HostConfigSTR = os.Getenv("HOSTSTR")
	PubKey = os.Getenv("PUBLIC_KEY_PATH")
	PrivKey = os.Getenv("PRIVATE_KEY_PATH")
	Engine, err = xorm.NewEngine("mysql", dsn)
	if err != nil {
		log.Fatal().Msgf("Error connecting to MySQL server: %s", err)
	}
	log.Debug().Msg("Connected to MySQL")
	return nil
}
