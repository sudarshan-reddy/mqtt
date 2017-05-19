package mq

import (
	"net"
	"net/url"

	"github.com/sudarshan-reddy/mqtt"
	proto "github.com/sudarshan-reddy/mqtt/mqttproto"
)

// Client is a very light wrapper on mqtt
type Client struct {
	client *mqtt.ClientConn
	topic  string
	retain bool
}

// NewClient returns a new instance of Client
func NewClient(clientID, rawURL, topic string, retain bool) (*Client, error) {
	uri, _ := url.Parse(rawURL)
	username := uri.User.Username()
	password, _ := uri.User.Password()

	conn, err := net.Dial("tcp", uri.Host)
	if err != nil {
		return nil, err
	}

	client := mqtt.NewClientConn(conn)
	if err := client.Connect(username, password); err != nil {
		return nil, err
	}

	return &Client{client: client, topic: topic, retain: retain}, nil
}

// Publish publishes a message to mqtt
func (m *Client) Publish(message string) {
	m.client.Publish(&proto.Publish{
		Header:    proto.Header{Retain: m.retain},
		TopicName: m.topic,
		Payload:   proto.BytesPayload([]byte(message)),
	})
}

// Subscribe subscribes to an mqtt message
func (m *Client) Subscribe() chan proto.Payload {
	messages := make(chan proto.Payload)
	tq := []proto.TopicQos{proto.TopicQos{Topic: m.topic, Qos: 0}}
	m.client.Subscribe(tq)
	go func() {
		defer close(messages)
		for message := range m.client.Incoming {
			messages <- message.Payload
		}
	}()
	return messages
}

// Close disconnects
func (m *Client) Close() {
	m.client.Disconnect()
}
