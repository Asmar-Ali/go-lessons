package main

import "fmt"

func maxMessages(thresh int) int {
	totalCost := 0
	for i := 0; ; i++ {
		totalCost += 100 + i
		if totalCost > thresh {
			return i
		}
	}

}

func loop() {
	// cost := bulkSend(30)
	maxMessagesAllowed := maxMessages(200)
	fmt.Println("Total cost is ", maxMessagesAllowed)

}
