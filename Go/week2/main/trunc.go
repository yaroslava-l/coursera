package main

import "fmt"

func main() {
	var floatnumber float64
	fmt.Println("Type float number")
	_, _ = fmt.Scan(&floatnumber)
	fmt.Print(int(floatnumber))
}