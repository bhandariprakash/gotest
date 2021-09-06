package tests

import (
	"calculation/calculate"
	"testing"
)

func TestAdd(t *testing.T) {

	got := calculate.Add(2, 2)
	expected := 4
	if got != expected {
		t.Errorf("Got: '%v', wanted: '%v'", got, expected)
	}
}

func TestSub(t *testing.T) {
	got := calculate.Sub(3, 2)
	expected := 1
	if got != expected {
		t.Errorf("Got: '%v', wanted: '%v'", got, expected)
	}
}

func TestMultiplication(t *testing.T) {
	got := calculate.Multification(3, 2)
	expected := 6
	if got != expected {
		t.Errorf("Got: '%v', wanted: '%v'", got, expected)
	}
}

func TestDivision(t *testing.T) {
	got := calculate.Division(6, 2)
	expected := 3.0
	if got != expected {
		t.Errorf("Got: '%v', wanted: '%v'", got, expected)
	}
}
