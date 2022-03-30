package main

import (
	"fmt"
	"strconv"
)

func evaluateValue(value int64) string {
	replyMessage := ""
	//insert your code here

	//Do not remove these lines
	fmt.Println((replyMessage))
	return replyMessage

}

func main() {
	var compareValue string
	fmt.Println("Enter an integer: ")
	fmt.Scanln(&compareValue)

	//conversion of value
	valueInt, _ := strconv.ParseInt(compareValue, 10, 0)
	evaluateValue(valueInt)

}
