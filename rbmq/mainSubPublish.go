package main

import (
	"RBMQ/RBMQ"
	"strconv"
	"time"
)

func main()  {
	rbmq := RBMQ.NewRBMQPubSub("nuclearSimple")
	for i := 0; i < 100; i++ {
		rbmq.PublishSub("Hello" + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
		println("Hello" + strconv.Itoa(i))
	}
}