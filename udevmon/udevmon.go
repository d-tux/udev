// +build linux

package main

import (
	"log"

	"github.com/deniswernert/udev"
)

func main() {
	monitor, err := udev.NewMonitor()
	if nil != err {
		log.Fatalln(err)
	}

	defer monitor.Close()
	events := make(chan *udev.UEvent)
	monitor.Monitor(events)
	for {
		event := <-events
		log.Println(event.String())
	}
}
