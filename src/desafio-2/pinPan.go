package main

import "fmt"

func main() {
	fmt.Println("Brincadeira Pin Pan")
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println("Pin")
		} else if i%5 == 0 {
			fmt.Println("Pan")
		} else {
			fmt.Println(i)
		}
	}
}
