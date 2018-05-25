package main

import (
	"time"

	"github.com/brian-armstrong/gpio"
)

type event struct {
	Pin   uint
	Value uint
}

func watchPins(pins []uint) chan event {
	channel := make(chan event)
	watcher := gpio.NewWatcher()

	for _, pin := range pins {
		watcher.AddPin(pin)
	}

	go func() {
		for {
			pin, value := watcher.Watch()
			channel <- event{pin, value}
		}
	}()

	return channel
}

func debounce(input chan event) chan event {
	output := make(chan event)

	go func() {

		var e event
		timer := time.NewTimer(2 * time.Second)
		timer.Stop()

		for {
			select {
			case e = <-input:
				timer = time.NewTimer(500 * time.Millisecond)
			case <-timer.C:
				timer.Stop()
				output <- e
			}
		}
	}()

	return output
}
