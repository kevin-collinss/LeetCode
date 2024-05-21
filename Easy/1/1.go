package Easy

type Set map[int]int

func twoSum(nums []int, target int) []int {

	set := make(Set)

	for i, num := range nums {

		remaining := target - num

		if set.contains(remaining) {
			j := set[remaining]
			return []int{j, i}
		}
		set[num] = i
	}

	return []int{}
}

func (s Set) contains(num int) bool {
	_, ok := s[num]
	return ok
}
