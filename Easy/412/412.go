package Easy

import (
	"strconv"
)

func fizzBuzz(n int) []string {

	i := 1
	var answer string

	fizz_buzz_array := make([]string, n)

	for i <= n {
		if i%3 == 0 || i%5 == 0 {
			answer = "FizzBuzz"
			if i%5 != 0 {
				answer = "Fizz"
			} else if i%3 != 0 {
				answer = "Buzz"
			}
		} else {
			answer = strconv.Itoa(i)
		}

		fizz_buzz_array[i-1] = answer
		i++
	}

	return fizz_buzz_array
}
