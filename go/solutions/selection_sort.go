package solutions

type SelectionSortSolver struct{}

func (s SelectionSortSolver) Solve(input map[string]any) any {
	// receives list, sorts it
	arr := input["arr"].([]int)  // assertion

	if len(arr) <= 1 {
		return arr
	}

	var length int = len(arr)
	var j, min, minIdx int

	for i := 0; i < length-1; i++ {
		min = arr[i]
		minIdx = i

		for j = i+1; j < length; j++ {
			if min > arr[j] {
				min = arr[j]
				minIdx = j
			}
		}

		arr[minIdx] = arr[i]
		arr[i] = min
	}

	return arr
}
