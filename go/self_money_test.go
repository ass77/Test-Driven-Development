package main

import (
	"testing"
)

type Dollar struct {
	amount int
}

func (d Dollar) Times(multiplier int) Dollar {
	return Dollar{amount: d.amount * multiplier}
}

func TestMultiplication(t *testing.T) {
	fiveDollars := Dollar{amount: 5}
	actualResult := fiveDollars.Times(2)
	expectedResult := Dollar{10}
	if actualResult != expectedResult {
		t.Errorf("Expected %d but got %d", expectedResult, actualResult)
	}
}
