package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {

	users := make(map[string]string)

	in := bufio.NewScanner(os.Stdin)

	fmt.Printf("Enter name :")
	in.Scan()
	users["name"] = in.Text()

	fmt.Printf("Enter address :")
	in.Scan()
	users["address"] = in.Text()


	userdata, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf(string(userdata))
	}
}