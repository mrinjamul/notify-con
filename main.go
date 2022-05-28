package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gen2brain/beeep"
	"github.com/kardianos/service"
	"github.com/robfig/cron"
)

const serviceName = "Internet Notify service"
const serviceShortName = "Internet Notify"
const serviceDescription = "Internet notify will notify internet is connected or lost."

var logger service.Logger

type program struct{}

func main() {
	svcConfig := &service.Config{
		Name:        serviceName,
		DisplayName: serviceShortName,
		Description: serviceDescription,
	}

	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatal(err)
	}
	logger, err = s.Logger(nil)
	if err != nil {
		log.Fatal(err)
	}
	err = s.Run()
	if err != nil {
		logger.Error(err)
	}
}

// service interfaces

func (p *program) Start(s service.Service) error {
	// Start should not block. Do the actual work async.
	go p.run()
	return nil
}
func (p *program) run() {
	CheckConnection()
}
func (p *program) Stop(s service.Service) error {
	// Stop should not block. Return with a few seconds.
	return nil
}

// User functions

// CheckConnection check if internet is up
func CheckConnection() {
	fmt.Println("INFO - Starting internet connection notifier...")
	flagConnected := true
	fmt.Println("INFO - Checking connection...")
	c := cron.New()
	// runs every 5 seconds
	c.AddFunc("*/5 * * * * *", func() {
		f := Retry(2, Connected)
		if flagConnected != f {
			flagConnected = f
			if flagConnected {
				fmt.Println("Internet connection is back")
				Notify("Internet connection is back", "", "info")
			} else {
				fmt.Println("Internet connection is lost")
				Notify("Internet connection is lost", "", "warning")
			}
		}
	})
	c.Start()
}

// Connected checks if the client is connected to the server
func Connected() (ok bool) {
	client := http.Client{
		Timeout: 1 * time.Second,
	}
	_, err := client.Get("http://clients3.google.com/generate_204")
	return err == nil
}

// Retry tries to connect to the server
func Retry(times int, f func() bool) bool {
	var ok bool
	for i := 0; i < times; i++ {
		ok = f()
		if ok {
			return ok
		}
	}

	return false
}

// Notify sends a notification to the user
func Notify(title, message string, notiType string) error {
	switch notiType {
	case "info":
		err := beeep.Notify(title, message, "assets/information.png")
		if err != nil {
			return err
		}
		err = beeep.Beep(beeep.DefaultFreq, 1000)
		if err != nil {
			panic(err)
		}

	case "warning":
		err := beeep.Alert(title, message, "assets/warning.png")
		if err != nil {
			return err
		}
		err = beeep.Beep(beeep.DefaultFreq, 1000)
		if err != nil {
			panic(err)
		}

	default:
		err := beeep.Notify(title, message, "assets/information.png")
		if err != nil {
			return err
		}
		err = beeep.Beep(beeep.DefaultFreq, 1000)
		if err != nil {
			panic(err)
		}
	}
	return nil
}
