package main

import "fmt"
import "time"

func main() {
	//var name = "ou"
	//number := 56
	fmt.Println("hello world!!!!", validate("hello", 5))
	fmt.Println("time", time.Now())

}
func validate(input string, number int) int {
	if input == "hello" {
		return 4 * number

	}
	return 0 * number

}
