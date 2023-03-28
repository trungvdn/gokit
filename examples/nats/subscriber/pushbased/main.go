package main

import (
	"encoding/json"
	"fmt"
	"log"
	"runtime"

	"github.com/nats-io/nats.go"
	"github.com/shijuvar/go-distsys/jsdemo/model"
)

func main() {
	nc, _ := nats.Connect(nats.DefaultURL)
	js, err := nc.JetStream()
	if err != nil {
		log.Fatal(err)
	}
	subOpt := []nats.SubOpt{
		nats.ManualAck(),
		nats.MaxDeliver(3),
		nats.DeliverAll(),
	}
	js.Subscribe("ORDERS.*", func(msg *nats.Msg) {
		// msg.Ack()
		var order model.Order
		err := json.Unmarshal(msg.Data, &order)
		if order.OrderID%2 == 0 {
			msg.Ack()
		}
		if err != nil {
			fmt.Println(err)
			log.Printf("monitor service subscribes from subject:%s\n", msg.Subject)
			log.Printf("OrderID:%d, CustomerID: %s, Status:%s\n", order.OrderID, order.CustomerID, order.Status)
		}
		fmt.Println("received data: ", order.OrderID)
	}, subOpt...)
	runtime.Goexit()
}
