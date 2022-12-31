package main

import (
	"flag"
	"fmt"
)

func main() {

	broker := flag.String("b", "", "Specify an mqtt broker (tcp://mqtt-broker-host.foobar.com:1883)")
	user := flag.String("u", "", "Specify an mqtt user")
	pass := flag.String("p", "", "Specify an mqtt password")
	flag.Parse()

	pins := []uint{26} // GPIO pin numbers, NOT the position on the raspberry pi
	channel := debounce(watchPins(pins))

	mqtt := MqttClient{Broker: *broker, User: *user, Pass: *pass}
	mqtt.connect("raspi-gpio-client")

	fmt.Println("Ready...")
	for {
		select {
		case e := <-channel:
			if e.Value == 0 {
				mqtt.publish("garage/door/status", "open")
			} else {
				mqtt.publish("garage/door/status", "closed")
			}
		}
	}
}
