package fizzbuzz_test

import (
	"testing"

	"github.com/PaiizZ/golang/fizzbuzz"
)

func TestFizzBuzz(t *testing.T) {
	if fizzbuzz.FizzBuzz(1) != 1 {
		t.Errorf("FizzBuzz(1) should be equal 1")
	}
}

func TestFizzBuzz1(t *testing.T) {
	if fizzbuzz.FizzBuzz(2) != 2 {
		t.Errorf("FizzBuzz(1) should be equal 2")
	}
}

func TestFizzBuzz2(t *testing.T) {
	if fizzbuzz.FizzBuzz(3) != "fizz" {
		t.Errorf("FizzBuzz(1) should be equal fizz")
	}
}

func TestFizzBuzz3(t *testing.T) {
	if fizzbuzz.FizzBuzz(5) != "buzz" {
		t.Errorf("FizzBuzz(1) should be equal buzz")
	}
}

func TestFizzBuzz4(t *testing.T) {
	if fizzbuzz.FizzBuzz(15) != "fizzbuzz" {
		t.Errorf("FizzBuzz(1) should be equal fizzbuzz")
	}
}
