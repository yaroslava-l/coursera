package main
import (
	"bufio"
	"fmt"
	"os"
	"strings"
)
type user struct {
	firstname string
	lastname string
}

func main() {

	var fileName string
	var users = make([]user,0,10)

	in := bufio.NewScanner(os.Stdin)
	fmt.Printf("Enter a file name: ")
	in.Scan()
	fileName = in.Text()
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		var userdata = user {"",""}
		userdata.firstname = (strings.Split(scanner.Text(), " "))[0]
		userdata.lastname = (strings.Split(scanner.Text(), " "))[1]
		users = append(users, userdata )
	}

	for _, currentValue := range users{
		fmt.Println(currentValue.firstname, currentValue.lastname)
	}
}
