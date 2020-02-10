package chop

// Chop takes an integer search target and a sorted, in ascending order, slice of integers.
// It returns the integer index of the target in the array, or -1 if the target is not in the slice.
func Chop(num int, list []int) int {
	ret, done := precheck(num, list)
	if done {
		return ret
	}
	split := len(list) / 2
	left := list[:split]
	right := list[split:]
	offset := 0
	switch {
	case inRange(num, left):
		offset = Chop(num, left)
		if offset >= 0 { //we have a winner
			return offset
		}
	case inRange(num, right):
		offset = Chop(num, right)
		if offset >= 0 { //we have a winner
			return offset + split
		}
	default:
		return -1
	}

	return -1
}

//check of corner cases/termination conditions
func precheck(num int, list []int) (int, bool) {
	//some corner cases
	if len(list) == 0 {
		return -1, true
	}
	if len(list) == 1 {
		if list[0] == num {
			return 0, true
		}
		if list[0] != num {
			return -1, true
		}
	}
	if !inRange(num, list) {
		return -1, true
	}

	return -1, false
}

//inRange test if an int is in the range of a sorted slice of int
func inRange(num int, list []int) bool {
	if num < list[0] || num > list[len(list)-1] {
		return false
	}
	return true
}
