package main

import "RBMQ/RBMQ"

func main()  {
	rbmq := RBMQ.NewRBMQSimple("nuclearSimple")
	rbmq.ReceiveSimple()
}
