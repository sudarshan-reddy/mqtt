# MQTT Client

## Publish

```
package main

import (
	"log"

	"github.com/sudarshan-reddy/mqtt/mq"
)

const (
	mqttURL = "mqtt://username:password@host:port"
)

func main() {
	client, err := mq.NewClient("clientID", mqttURL, "test_topic", false)
	if err != nil {
		log.Fatal("new client error", err)
	}

	defer client.Close()
	client.Publish("message 12313212")
}
```

## Subscribe
```
package main

import (
	"log"
	"os"

	"github.com/sudarshan-reddy/mqtt/mq"
)

const (
	mqttURL = "mqtt://username:hostname@host:port"
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
```
