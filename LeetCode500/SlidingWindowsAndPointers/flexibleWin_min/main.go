package main

import "fmt"

func main() {
	mystr := "10011"

	for _, v := range mystr {
		fmt.Printf("%d", int(v&0))
		break
	}
	// fmt.Println()
	// fmt.Printf("%d", '0')
}
