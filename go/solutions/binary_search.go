package solutions

type BinarySearchSolver struct{}

func (s BinarySearchSolver) Solve(input map[string]any) any {
	// arr = array, target = target to find
	arr := input["arr"].([]int)  // type assertion
	target := input["target"].(int)

	var high int = len(arr) - 1
	var low int = 0

	var guess, mid int

	for low <= high {
		mid = (low + high) / 2
		guess = arr[mid]

		if guess == target {
			return mid
		}

		if guess < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}
