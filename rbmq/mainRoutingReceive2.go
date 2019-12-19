package main

import "RBMQ/RBMQ"

func main()  {
	rbmq := RBMQ.NewRBMQRouting("simple", "simple_two")
	rbmq.ReceiveRouting()
}
