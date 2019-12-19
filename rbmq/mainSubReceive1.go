package main

import "RBMQ/RBMQ"

func main()  {
	rbmq := RBMQ.NewRBMQPubSub("nuclearSimple")
	rbmq.ReceiveSub()
}
