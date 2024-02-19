package main

import (
	"fmt"
	"time"
)

func main() {
	orders := make(chan string, 3) // Buffered channel with a capacity of 3

	// Customers placing orders
	go func() {
		for i := 1; i <= 5; i++ {
			order := fmt.Sprintf("Coffee order #%d", i)
			orders <- order // Only blocks if the buffer is full
			fmt.Println("🗒️  Placed:", order)
		}
		close(orders)
	}()

	// Barista processing orders
	for order := range orders {
		fmt.Printf("🫘 Preparing: %s\n", order)
		time.Sleep(2 * time.Second) // Time taken to prepare the order
		fmt.Printf("☕ Served: %s\n", order)
	}
}
