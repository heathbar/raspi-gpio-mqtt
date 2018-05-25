# raspi-gpio-mqtt

A little app written in Go to watch for changes to GPIO pins on the Raspberry Pi and relay them to an MQTT topic. 

The app uses the os.select() to receive changes from the kernel rather than polling to keep CPU usage to a minimum. It also debounces the input so that the MQTT topic does not get spammed. 

Dependencies: github.com/brian-armstrong/gpio, github.com/eclipse/paho.mqtt.golang

## Setup
```bash
git clone git@github.com:heathbar/raspi-gpio-mqtt.git
cd raspi-gpio-mqtt
go get
go build

# root/sudo is required to access the GPIO pins on the raspberry pi. 
sudo ./raspi-gpi-mqtt -b tcp://my-mqtt-server.com:1883
```

## TODO
- mqtt security
- mqtt QoS
- configurable debounce timeout
