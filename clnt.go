package main

import (
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"time"
	"os"
	"math/rand"
	"fmt"
)

func send(word string) {
	var delay int = rand.Intn(10)

	opts := MQTT.NewClientOptions().AddBroker("tcp://localhost:1883")
	c := MQTT.NewClient(opts)

	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	topic := fmt.Sprintf("topic_%s", word)
	fmt.Println("Sending", word, "to", topic)
	for {
		time.Sleep(time.Duration(delay*100) * time.Millisecond)
		token := c.Publish(topic, 0, false, word)
		token.Wait()
	}
}

func main() {
	rand.Seed(int64(time.Now().Nanosecond()))
	for _,str := range os.Args[1:] {
		go send(str)
	}
	for {
	}
}
