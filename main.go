package main

import (
	"TestTask/network"
	"fmt"
)

func main() {
	net := network.MakeNetwork(10)

	net.Connect(9, 2)
	net.Connect(1, 3)
	net.Connect(8, 4)
	net.Connect(3, 2)
	net.Connect(1, 6)

	fmt.Printf("query: %+v\n", net.Query(9, 6)) //should print 'true'
}
