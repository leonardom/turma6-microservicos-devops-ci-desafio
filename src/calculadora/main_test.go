package main

import "testing"

func TestSoma(t *testing.T) {
	got := Soma(5, 5)
	expected := 10

	if (got != expected) {
		t.Errorf("Soma(5, 5) \n got: %v, expexted: %v\n", got, expected)
	}
}