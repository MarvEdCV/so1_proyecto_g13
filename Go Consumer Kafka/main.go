package main

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"time"
)

func main() {
	log.Println("Estoy en el consumer :)")
	conn, _ := kafka.DialLeader(context.Background(), "tcp", "kafka-cluster-kafka-bootstrap:9092", "input-kafka", 0)
	conn.SetWriteDeadline(time.Now().Add(time.Second * 10))

	message, _ := conn.ReadMessage(1e6)
	/*bytes := make([]byte, 1e3)
	for {
		_, err := batch.Read(bytes)
		if err != nil {
			break
		}*/
	fmt.Println(message.Value)
	//}
}
