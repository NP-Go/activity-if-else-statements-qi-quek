package main

import (
	"fmt"
	"strconv"
)

func evaluateValue(value int64) string {
	replyMessage := ""
	//insert your code here
	if value%5 == 0 && value%6 == 0 {
		replyMessage = "Value is divisible by both 5 and 6"
	} else if value%5 == 0 && value%6 != 0 {
		replyMessage = "value is divisible by 5 but not 6"
	} else if value%6 == 0 && value%5 != 0 {
		replyMessage = "value is divisible by 6 but not 5"
	} else {
		replyMessage = "Value is not divisable by both 5 or 6"
	}
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
