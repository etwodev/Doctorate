package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Etwodev/Doctorate/server"
	"github.com/Etwodev/Doctorate/server/router"
	"github.com/Etwodev/Doctorate/server/router/announce"
	"github.com/joho/godotenv"
)


func main () {
	err := godotenv.Load()
	if err != nil {
	  log.Fatal("Error loading .env file")
	}

	cfg := server.Config {
		Version:    os.Getenv("VERSION"),
		Port:       os.Getenv("PORT"),
		Connection: fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=true&multiStatements=true", os.Getenv("USR"), os.Getenv("PASS"), os.Getenv("SERVER"), os.Getenv("DB")),
	}

	var s = server.New(&cfg)

	routers := []router.Router{
		announce.NewRouter(true),
	}

	s.InitRouter(routers...)

	var r = s.InitRouters(true)

	log.Fatal(http.ListenAndServe(cfg.Port, r))
}