package problems

import env "algorithms/environment"

var BinarySearch = env.Problem{
	Name: "Binary Search",
	TestCases: []env.TestCase{
		env.TestCase{InputData: map[string]any{"arr": []int{1, 2, 3, 4, 5}, "target": 3}, ExpectedOutput: 2},
		env.TestCase{InputData: map[string]any{"arr": []int{1, 2, 3, 4, 5}, "target": 6}, ExpectedOutput: -1},
		env.TestCase{InputData: map[string]any{"arr": []int{1, 2, 3, 4, 5}, "target": 1}, ExpectedOutput: 0},
		env.TestCase{InputData: map[string]any{"arr": []int{1, 2, 3, 4, 5}, "target": 5}, ExpectedOutput: 4},
	},
}
