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
			fmt.Println(e)
			fmt.Println(fmt.Sprintf("pin/%d", e.Pin))
			mqtt.publish(fmt.Sprintf("pin/%d", e.Pin), fmt.Sprintf("%d", e.Value))
		}
	}
}
