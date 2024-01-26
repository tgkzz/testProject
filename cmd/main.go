package main

import (
	"fmt"
	"os"
	"testProject/config"
	"testProject/internal/handler"
	"testProject/internal/repository"
	"testProject/internal/server"
	"testProject/internal/service"
	"testProject/logger"
)

// TODO: Start server pls
func main() {
	infoLog, errLog, err := logger.NewLogger()
	if err != nil {
		fmt.Println(err)
		return
	}

	var cfgPath string

	switch len(os.Args[1:]) {
	case 1:
		cfgPath = os.Args[1]
	case 0:
		cfgPath = "./.env"
	default:
		errLog.Fatal("USAGE: go run [CONFIG_PATH]")
	}

	// init config
	cfg, err := config.LoadConfig(cfgPath)
	if err != nil {
		errLog.Fatal(err)
	}

	//init db
	db, err := repository.LoadDB(cfg.DB.DriverName, cfg.DB.DataSourceName)
	if err != nil {
		errLog.Fatal(err)
	}

	r := repository.NewRepository(db)

	s := service.NewService(*r, cfg.URL)

	h := handler.NewHandler(s, infoLog, errLog)

	errLog.Fatal(server.StartServer(cfg, h.Routes(infoLog, errLog), infoLog))

}
