package main

import (
	"fmt"
)

type car struct {
	gasPedal      uint16 // min 0 to 65535
	brakePedal    uint16
	steeringWheel int16 // -32k to +32k
	topSpeed      float64
}

func main() {
	var tesla car = car{
		gasPedal:      100,
		brakePedal:    20,
		steeringWheel: 20000,
		topSpeed:      350.5,
	}
	fmt.Println(tesla)
}
