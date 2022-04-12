package main

import "testing"

func TestSum(t *testing.T) {
	if sum(1, 2) != 3 {
		t.Error("Expected 3, got ", sum(1, 2))
	}
}

func TestSumString(t *testing.T) {
	// Table Drive Test
	// https://golang.org/pkg/testing/#hdr-Table_driven_tests
	tests := []struct {
		Escenarie string
		A         string
		B         string
		Error     bool
		Expected  int
	}{
		{
			Escenarie: "sum two numbers",
			A:         "1",
			B:         "2",
			Error:     false,
			Expected:  3,
		},
		{
			Escenarie: "Incorrect sum two numbers, invalid A",
			A:         "a",
			B:         "2",
			Error:     true,
			Expected:  0,
		},
		{
			Escenarie: "Incorrect sum two numbers, invalid B",
			A:         "1",
			B:         "b",
			Error:     true,
			Expected:  0,
		},
	}
	for _, test := range tests {
		t.Run(test.Escenarie, func(t *testing.T) {
			t.Parallel() // Run tests in parallel
			c, error := sumString(test.A, test.B)
			if test.Error != (error != nil) {
				t.Error("Expected error ", test.Error, " got ", error)
			}
			if c != test.Expected {
				t.Error("Expected ", test.Expected, " got ", c)
			}
		})
	}
}
