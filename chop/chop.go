package chop

import "math"

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
	if len(list) == 1 && list[0] != num {
		return -1, true
	}
	if list[0] == num {
		return 0, true
	}
	if list[len(list)-1] == num {
		return len(list) - 1, true

	}
	if !inRange(num, list) {
		return -1, true
	}

	return -1, false
}

//inRange test if an int is in the range of a sorted slice of int
func inRange(num int, list []int) bool {
	if len(list) == 0 {
		return false
	}
	if num < list[0] || num > list[len(list)-1] {
		return false
	}
	return true
}

// Chop2 takes an integer search target and a sorted, in ascending order, slice of integers.
// It returns the integer index of the target in the array, or -1 if the target is not in the slice.
func Chop2(num int, list []int) int {
	ret, done := precheck(num, list) //maybe I should rename precheck
	if done {
		return ret
	}
	iter := int(math.Ceil(math.Log2(float64(len(list))))) //worst case
	var split, offset, next int
	for i := 0; i <= iter; i++ { //if we dont find our number in the worst case we are doomed!
		split = len(list) / 2
		next = list[split]
		switch {
		// case len(list) <= 1 && next != num:
		// 	return -1
		case next == num:
			return split + offset
		case num > next: //right
			offset += split
			list = list[split:]
		case num < next: //left
			list = list[:split]
		}
	}
	return -1
}

// Chop3 takes an integer search target and a sorted, in ascending order, slice of integers.
// It returns the integer index of the target in the array, or -1 if the target is not in the slice.
func Chop3(num int, list []int) int {
	ret, done := precheck(num, list)
	if done {
		return ret
	}
	iter := int(math.Ceil(math.Log2(float64(len(list))))) //worst case
	var (
		split        = len(list)
		offset, next int
	)
	for i := 0; i <= iter; i++ { //if we dont find our number in the worst case we are doomed!
		split /= 2
		next = list[split+offset]
		switch {
		case next == num:
			return split + offset
		case num > next: //right
			offset += split
		}
	}
	return -1
}

// Chop4 takes an integer search target and a sorted, in ascending order, slice of integers.
// It returns the integer index of the target in the array, or -1 if the target is not in the slice.
// This version is (an attempt to) an euristic approach, we look where the number should be
// in the list if the numbers were spread evenly.
func Chop4(num int, list []int) int {
	ret, done := precheck(num, list)
	if done {
		return ret
	}
	split := linearScale(num, list)
	if split == 0 {
		split = 1 //we already chacked the extremes, we need to narrow it down
	}
	left := list[:split]
	right := list[split:]
	offset := 0
	switch {
	case inRange(num, left):
		offset = Chop4(num, left)
		if offset >= 0 { //we have a winner
			return offset
		}
	case inRange(num, right):
		offset = Chop4(num, right)
		if offset >= 0 { //we have a winner
			return offset + split
		}
	default:
		return -1
	}
	return -1
}

//from https://stats.stackexchange.com/questions/281162/scale-a-number-between-a-range/281164
//we scale the number we are looking for from numeric interval of the number of the list to the position
func linearScale(n int, list []int) int {
	var (
		m    = float64(n)
		rmin = float64(list[0])
		rmax = float64(list[len(list)-1])
		// tmin = 0.0
		tmax = float64(len(list) - 1)
	)
	// ret := (m-rmin)/(rmax-rmin)*(tmax-tmin) + tmin
	ret := (m - rmin) / (rmax - rmin) * tmax

	return int(ret)
}
