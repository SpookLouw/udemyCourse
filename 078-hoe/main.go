package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	fmt.Println("this is where initialization for my program occurs")
}

func main() {
	rand.Seed(time.Now().UnixNano())
	x := rand.Intn(251)
	fmt.Printf("the value of x is %d\n", x)

	switch {
	case x >= 0 && x <= 100:
		fmt.Println("between 0 and 100")
	case x >= 101 && x <= 200:
		fmt.Println("between 101 and 200")
	case x >= 201 && x <= 250:
		fmt.Println("between 201 and 250")
	}

	fmt.Println("Demonstrating rand.Intn() with numbers 0-3:")
	for i := 0; i < 10; i++ {
		fmt.Println(rand.Intn(4))
	}
}
