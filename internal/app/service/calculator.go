package service

import (
	"errors"
)

func Sum(numbers []int) (int, error) {
	if len(numbers) < 1 {
		return 0, errors.New("Empty request body")
	}

	var sum int
	for _, num := range numbers {
		sum += num
	}
	return sum, nil
}

func Multiply(numbers []int) (int, error) {
	if len(numbers) < 1 {
		return 0, errors.New("Empty request body")
	}

	var multiply int
	multiply = 1
	for _, num := range numbers {
		multiply *= num
	}
	return multiply, nil
}
