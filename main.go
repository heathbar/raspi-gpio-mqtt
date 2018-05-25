package main

import (
	"flag"
	"fmt"
)

func main() {

	broker := flag.String("b", "", "Specify an mqtt broker (tcp://mqtt-broker-host.foobar.com:1883)")
	flag.Parse()

	pins := []uint{26} // GPIO pin numbers, NOT the position on the raspberry pi
	channel := debounce(watchPins(pins))

	mqtt := MqttClient{Broker: *broker}
	mqtt.connect("raspi-gpio-client")

	fmt.Println("Ready...")
	for {
		select {
		case e := <-channel:
			if e.Value == 0 {
				fmt.Println("closed")
				mqtt.publish("garage/door/status", "closed")
			} else {
				fmt.Println("open")
				mqtt.publish("garage/door/status", "open")
			}
		}
	}
}
