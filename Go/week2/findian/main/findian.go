package main

import (
	"fmt"
	"strings"
)
func main() {
	var scanner string
	fmt.Scanln(&scanner)
	fmt.Println("Enter a word")
	if scanner!="" {
		userword := strings.ToLower(scanner)
		i := strings.Index(userword, "i")
		n := strings.LastIndex(userword, "n")

		if i == 0 && strings.Contains(userword,"a") && n ==(len(userword)-1)  {
			fmt.Println("Found!")
		} else {
			fmt.Println("Not Found!")
		}
	}
}