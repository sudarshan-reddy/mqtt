package main

import (
	"log"
	"os"

	"github.com/sudarshan-reddy/mqtt/mq"
)

const (
	mqttURL = "mqtt://ykiyaquh:3brkxyleocTi@m10.cloudmqtt.com:14328"
)

func main() {
	client, err := mq.NewClient("sudarshan", mqttURL, "test_topic", false)
	if err != nil {
		log.Fatal("new client error", err)
	}

	defer client.Close()
	for msg := range client.Subscribe() {
		msg.WritePayload(os.Stdout)
	}
}
