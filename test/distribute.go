package main

import (
	"fmt"
)

// Order represents a single order
type Order struct {
	ID int
}

// distributeOrders divides the orders between regular users and a special user
func distributeOrders(totalOrders int, regularUsers int, specialUser bool) ([][]*Order, []*Order) {
	// Initialize slices for regular users and special user
	var regularUserOrders [][]*Order
	var specialUserOrders []*Order

	// Create orders
	orders := make([]*Order, totalOrders)
	for i := 0; i < totalOrders; i++ {
		orders[i] = &Order{ID: i + 1}
	}

	// Distribute orders to regular users
	ordersPerUser := 5
	for i := 0; i < regularUsers && i*ordersPerUser < totalOrders; i++ {
		start := i * ordersPerUser
		end := start + ordersPerUser
		if end > totalOrders {
			end = totalOrders
		}
		regularUserOrders = append(regularUserOrders, orders[start:end])
	}

	// Check if there are remaining orders and a special user
	if len(orders) > regularUsers*ordersPerUser && specialUser {
		specialUserOrders = orders[regularUsers*ordersPerUser:]
	}

	return regularUserOrders, specialUserOrders
}

func main() {
	// Total number of orders
	const totalOrders = 30
	// Number of regular users (x)
	const regularUsers = 3
	// Whether there is a special user
	const hasSpecialUser = true

	// Call the distribution function
	regularUserOrders, specialUserOrders := distributeOrders(totalOrders, regularUsers, hasSpecialUser)

	// Print the results
	fmt.Println("Regular Users Orders:")
	for i, userOrders := range regularUserOrders {
		fmt.Printf("User %d Orders:\n", i+1)
		for _, order := range userOrders {
			fmt.Println(order.ID)
		}
		fmt.Println() // Print an empty line for better visualization
	}

	if len(specialUserOrders) > 0 {
		fmt.Println("Special User Orders:")
		for _, order := range specialUserOrders {
			fmt.Println(order.ID)
		}
	} else {
		fmt.Println("No remaining orders for the special user.")
	}
}
