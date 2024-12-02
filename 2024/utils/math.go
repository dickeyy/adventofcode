package utils

func SumNums(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
