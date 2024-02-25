package main

import (
	"testing"
)

func TestFindFirstDigit(t *testing.T) {
	// Set up the test cases
	testCases := []struct {
		input    string
		expected int
	}{
		{"1abc2", 1},
		{"pqr3stu8vwx", 3},
		{"a1b2c3d4e5f", 1},
		{"treb7uchet", 7},
	}

	// Loop through the test cases
	for _, tc := range testCases {
		// Call the function
		actual := findFirstDigit(tc.input)
		// Compare the actual and expected results
		if actual != tc.expected {
			t.Errorf("findFirstDigit(%s): expected %d, actual %d", tc.input, tc.expected, actual)
		}
	}
}

func TestFindLastDigit(t *testing.T) {
	// Set up the test cases
	testCases := []struct {
		input    string
		expected int
	}{
		{"1abc2", 2},
		{"pqr3stu8vwx", 8},
		{"a1b2c3d4e5f", 5},
		{"treb7uchet", 7},
	}

	// Loop through the test cases
	for _, tc := range testCases {
		// Call the function
		actual := findLastDigit(tc.input)
		// Compare the actual and expected results
		if actual != tc.expected {
			t.Errorf("findLastDigit(%s): expected %d, actual %d", tc.input, tc.expected, actual)
		}
	}
}

func TestFindCalibrationValue(t *testing.T) {
	// Set up the test cases
	testCases := []struct {
		input    string
		expected int
	}{
		{"1abc2", 12},
		{"pqr3stu8vwx", 38},
		{"a1b2c3d4e5f", 15},
		{"treb7uchet", 77},
	}
	// Loop through the test cases
	for _, tc := range testCases {
		// Call the function
		actual := findCalibrationValue(tc.input)
		// Compare the actual and expected results
		if actual != tc.expected {
			t.Errorf("findCalibrationValue(%s): expected %d, actual %d", tc.input, tc.expected, actual)
		}
	}
}
