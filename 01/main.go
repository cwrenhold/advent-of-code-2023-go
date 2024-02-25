package main

import (
	"fmt"
	"strconv"
)

func findFirstDigit(s string) int {
	for _, c := range s {
		if c >= '0' && c <= '9' {
			i, err := strconv.Atoi(string(c))

			if err != nil {
				return -1
			}

			return i
		}
	}
	return -1
}

func findLastDigit(s string) int {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] >= '0' && s[i] <= '9' {
			digit, err := strconv.Atoi(string(s[i]))

			if err != nil {
				return -1
			}

			return digit
		}
	}
	return -1
}

func findCalibrationValue(s string) int {
	firstDigit := findFirstDigit(s)
	lastDigit := findLastDigit(s)
	if (firstDigit != -1 && lastDigit != -1) {
		return firstDigit * 10 + lastDigit
	}
	return -1
}

func main() {
	// Set up the calibration strings
	calibrationStrings := []string{
		"1abc2",
		"pqr3stu8vwx",
		"a1b2c3d4e5f",
		"treb7uchet",
	}

	// Initialise the calibration sum
	calibrationSum := 0

	// Loop through the calibration strings
	for _, s := range calibrationStrings {
		calibrationValue := findCalibrationValue(s)

		if (calibrationValue != -1) {
			fmt.Println("Calibration value:", calibrationValue)
			calibrationSum += calibrationValue
		} else {
			fmt.Println("No digits found")
		}
	}

	// Output the calibration sum
	fmt.Println("Calibration sum:", calibrationSum)
}

