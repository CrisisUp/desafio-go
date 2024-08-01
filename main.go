package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println("\n===============================================================================================")

	const reset = "\033[0m"
	const green = "\033[32m"
	const red = "\033[31m"
	const yellow = "\033[33m"

	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Print(red + "pin-pan " + reset)
		} else if i%3 == 0 {
			fmt.Print(green + "pin " + reset)
		} else if i%5 == 0 {
			fmt.Print(yellow + "pan " + reset)
		} else {
			fmt.Printf(" %d ", i)
		}
		if i%20 == 0 {
			fmt.Println()
		}
	}
	fmt.Println()

}
