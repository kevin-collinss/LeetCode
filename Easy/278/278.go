package Easy

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
	low := 1
	high := n

	for low <= high {

		med := low + (high-low)/2
		current := isBadVersion(med)
		prev := isBadVersion(med - 1)

		if current == true && prev == false {
			return med
		} else if current == false && prev == false {
			low = med + 1
		} else {
			high = med
		}
	}

	return -1
}
