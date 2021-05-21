package main

import (
	"adinunno.fr/ubiquiti-influx-monitoring/src/infra"
	"adinunno.fr/ubiquiti-influx-monitoring/src/service"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

var waitlock = &sync.WaitGroup{}

func tick() {
	http.DefaultTransport.(*http.Transport).CloseIdleConnections()
	service.Tick()
}

func catchSigInt() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for sig := range c {
			if sig == syscall.SIGINT {
				waitlock.Done()
			}
		}
	}()
}

func main() {
	infra.LoadEnv()
	config := infra.LoadConfig()
	onesignalConfig := infra.LoadOnesignalConfig()

	db := service.CreateDB()
	service.LoadService(config, onesignalConfig, db)

	ticker := time.NewTicker(10 * time.Second)
	quit := make(chan struct{})
	waitlock.Add(1)

	catchSigInt()

	tick()
	go func() {
		for {
			select {
			case <-ticker.C:
				tick()
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

	waitlock.Wait()
	println("Cleaning up...")
}
