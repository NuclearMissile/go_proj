package main

import (
	"RBMQ/RBMQ"
	"fmt"
	"strconv"
	"time"
)

func main() {
	r1 := RBMQ.NewRBMQRouting("simple", "simple_one")
	r2 := RBMQ.NewRBMQRouting("simple", "simple_two")
	for i := 0; i < 100; i++ {
		r1.PublishRouting(fmt.Sprintf("Hello one: %s", strconv.Itoa(i)))
		r2.PublishRouting(fmt.Sprintf("Hello two: %s", strconv.Itoa(i)))
		time.Sleep(1 * time.Second)
		println(i)
	}
}
