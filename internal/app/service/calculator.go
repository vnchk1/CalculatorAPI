package service

import (
	"errors"
)

func Sum(numbers []int) (result int, err error) {
	if len(numbers) < 1 {
		err = errors.New("empty request body")
		return
	}

	for _, num := range numbers {
		result += num
	}
	return
}

func Multiply(numbers []int) (result int, err error) {
	if len(numbers) < 1 {
		err = errors.New("empty request body")
		return
	}

	result = 1
	for _, num := range numbers {
		result *= num
	}
	return
}
