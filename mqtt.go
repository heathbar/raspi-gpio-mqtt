package main

import (
	"fmt"
	"log"
	"os"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// MqttClient is a wrapper for the eclipse mqtt library
type MqttClient struct {
	Broker string
	client mqtt.Client
}

func (c *MqttClient) connect(clientID string) bool {

	hostname, _ := os.Hostname()
	client := fmt.Sprintf("%s-%s", hostname, clientID)
	fmt.Printf("Connecting to %s as %s\n", c.Broker, client)

	mqtt.ERROR = log.New(os.Stderr, "", 0)
	opts := mqtt.NewClientOptions().AddBroker(c.Broker)
	opts.SetKeepAlive(2 * time.Second)
	opts.SetPingTimeout(1 * time.Second)
	opts.SetClientID(client)
	opts.SetAutoReconnect(true)
	opts.SetMaxReconnectInterval(2 * time.Minute)

	c.client = mqtt.NewClient(opts)
	if token := c.client.Connect(); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		return false
	}
	return true
}

func (c *MqttClient) publish(topic string, message interface{}) {
	c.client.Publish(topic, 0, false, message)
}
