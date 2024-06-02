package Easy

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
	low := 1
	high := n

	for low <= high {
		myGuess := low + (high-low)/2
		result := guess(myGuess)
		if (result) == 0 {
			return myGuess
		} else if result == -1 {
			high = myGuess - 1
		} else {
			low = myGuess + 1
		}
	}
	return -1
}
