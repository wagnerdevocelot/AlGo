package main

func sumRange(num int) int {
	if num == 1 {
		return 1
	}
	return num + sumRange(num-1)
}
