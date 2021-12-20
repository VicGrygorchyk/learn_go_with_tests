package adder

// Add two integers and returns the sum of them
func Add(first int, second int) (int, error) {
	return first + second, nil
}

func Sum(numbers[] int) int {
	res := 0
	for _, num := range(numbers) {
		res += num
	}
	return res
}

func AddAll(nums...[]int) (summed []int) {
	// sum := make([]int, len(nums)) // not good as we may need slice with bigger capacity
	var sums []int

	for _, num := range nums {
		sums = append(sums, Sum(num))
	}

	return sums
}