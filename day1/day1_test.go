package day1

import (
	"testing"
)

func TestDay1(t *testing.T) {
	result, err := Day1()
	if err != nil {
		t.Errorf("Day1() = %d; want 0", result)
	}

	if result != 0 {
		t.Errorf("Day1() = %d; want 0", result)
	}
}
