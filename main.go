package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	s "github.com/batasrki/ppe4peeps/cmd"
	"github.com/batasrki/ppe4peeps/config"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(config.LogLevel())
}

func main() {
	startTime := time.Now()
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		log.WithField("uptime", time.Since(startTime).String()).
			WithField("signal", sig.String()).Error("Interrupt signal detected")
		os.Exit(0)
	}()

	server := s.Server{
		Port: config.Port(),
	}

	log.Fatal(server.ListenAndServe())

}
