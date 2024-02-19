package main

import (
	"fmt"
	"time"
)

func main() {
	orders := make(chan string) // Unbuffered channel

	// Customer placing orders
	go func() {
		for i := 1; i <= 5; i++ {
			order := fmt.Sprintf("Coffee order #%d", i)
			orders <- order // Blocks until the barista is ready to accept order
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
