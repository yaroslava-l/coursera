package main

import (
"fmt"
"sort"
	"strconv"
)

func main() {
	var slice []int = make([]int, 3)
	var number string

	for {
		fmt.Println("Enter an  Integer (X-exit)")
		fmt.Scanln(&number)
		if number == "X"{
			break
		}
		ap,err:=strconv.Atoi(number)
		if err != nil {
			continue
		}
		slice = append(slice, ap)
		sort.Ints(slice[:])
		fmt.Println(slice)
	}
}