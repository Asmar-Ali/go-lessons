package main

import "fmt"

/* 1st message: 1.0 + 0.00
2nd message: 1.0 + 0.01
3rd message: 1.0 + 0.02
4th message: 1.0 + 0.03 */

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
