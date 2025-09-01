package service

import (
	"errors"
	"math"
)

func Sum(numbers []int) (result int, err error) {
	if len(numbers) < 1 {
		err = errors.New("empty request body")
		return
	}

	for _, num := range numbers {
		if (num > 0 && result+num > math.MaxInt-num) || (num < 0 && result < math.MinInt-num) {
			return 0, errors.New("number out of range")
		}
		result += num
	}
	return
}

func Multiply(numbers []int) (int, error) {
	if len(numbers) == 0 {
		return 0, errors.New("empty request body")
	}

	result := 1
	for _, num := range numbers {
		// проверка на переполнение перед умножением
		if num != 0 {
			product := result * num
			if product/num != result {
				return 0, errors.New("number out of range")
			}
		}
		result *= num
	}
	return result, nil
}
