package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Seed the random number generator to get different results each time
	rand.Seed(time.Now().UnixNano())

	// Generate a random integer between 0 and 250
	x := rand.Intn(251) // Since Intn(n) generates a number in the range [0, n)

	// Print the name and value of the variable
	fmt.Printf("The value of x is %d\n", x)

	// Use a switch statement to check the range and print the appropriate message
	switch {
	case x >= 0 && x <= 100:
		fmt.Println("between 0 and 100")
	case x >= 101 && x <= 200:
		fmt.Println("between 101 and 200")
	case x >= 201 && x <= 250:
		fmt.Println("between 201 and 250")
	}

	// Demonstrating the behavior of rand.Intn() with numbers 0-3
	fmt.Println("Demonstrating rand.Intn() with numbers 0-3:")
	for i := 0; i < 10; i++ {
		fmt.Println(rand.Intn(4)) // Generates numbers in the range [0, 4)
	}
}
