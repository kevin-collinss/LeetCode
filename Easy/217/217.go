package Easy

type Set map[int]struct{}

func containsDuplicate(nums []int) bool {

	set := make(Set)

	for _, num := range nums {
		if set.contains(num) {
			return true
		}
		set[num] = struct{}{}
	}
	return false
}

func (s Set) contains(num int) bool {
	_, ok := s[num]
	return ok
}
