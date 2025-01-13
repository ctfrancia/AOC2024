package day2

import (
	"fmt"
	"testing"
)

func TestDay1(t *testing.T) {
	result, err := Day2()
	if err != nil {
		t.Errorf("error returned, %v", err)
	}
	fmt.Println(result)
}
