package main

import "testing"

func TestIfElse(t *testing.T) {
	get := evaluateValue(100)
	res := "The value is NOT divisible by 5 and 6"
	if get != res {
		t.Errorf("Wanted %s Got %s", res, get)
	}

	get = evaluateValue(30)
	res = "The value is divisible by 5 and 6"
	if get != res {
		t.Errorf("Wanted %s Got %s", res, get)
	}
}
