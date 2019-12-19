package main

import "RBMQ/RBMQ"

func main()  {
	r1 := RBMQ.NewRBMQRouting("simple", "simple_one")
	r2 := RBMQ.NewRBMQRouting("simple", "simple_two")
	r1.ReceiveRouting()
	r2.ReceiveRouting()
}
