package server

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"testProject/config"
	"time"
)

// TODO: add config
func StartServer(cfg config.Config, g *gin.Engine, infoLog *log.Logger) error {
	srv := &http.Server{
		Addr:         cfg.Host + ":" + cfg.Port,
		Handler:      g,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	infoLog.Printf("server is listening on http://%s:%s", cfg.Host, cfg.Port)

	err := srv.ListenAndServe()

	return err
}
