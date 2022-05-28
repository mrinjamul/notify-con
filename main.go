package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gen2brain/beeep"
)

func main() {
	fmt.Println("INFO - Starting internet connection notifier...")
	CheckConnection()
}

// CheckConnection check if internet is up
func CheckConnection() {
	timeTicker := time.NewTicker(5 * time.Second)
	var flagConnected bool = true
	fmt.Println("INFO - Checking connection...")
	for range timeTicker.C {
		f := Retry(Connected)
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
	}
}

// Connected checks if the client is connected to the server
func Connected() (ok bool) {
	_, err := http.Get("http://www.google.com")
	return err == nil
}

// Retry tries to connect to the server
func Retry(f func() bool) bool {
	var ok bool
	for i := 0; i < 2; i++ {
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
	case "warning":
		err := beeep.Alert(title, message, "assets/warning.png")
		if err != nil {
			return err
		}
	default:
		err := beeep.Notify(title, message, "assets/information.png")
		if err != nil {
			return err
		}
	}
	return nil
}
