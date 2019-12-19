package main

import (
	"RBMQ/RBMQ"
	"strconv"
	"time"
)

func main()  {
	rbmq := RBMQ.NewRBMQSimple("nuclearSimple")
	rbmq.PublishSimple("Hello")
	for i := 0; i < 100; i++ {
		rbmq.PublishSimple("Hello" + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
		println(i)
	}
}