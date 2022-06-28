package main

import "testing"

func TestAdd(t *testing.T) {
	result := Add(1, 2)
	if result != 3 {
		t.Errorf("Failed test %v", result)
	}
}

func Add(a, b int) int {
	return a + b
}
