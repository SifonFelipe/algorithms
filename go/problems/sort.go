package problems

import env "algorithms/environment"

var Sort = env.Problem{
	Name: "Sort",
	TestCases: []env.TestCase{
		env.TestCase{InputData: map[string]any{"arr": []int{5, 2, 9, 1, 5, 6}}, ExpectedOutput: []int{1, 2, 5, 5, 6, 9}},
		env.TestCase{InputData: map[string]any{"arr": []int{3, 0, -2, 5, -1}}, ExpectedOutput: []int{-2, -1, 0, 3, 5}},
		env.TestCase{InputData: map[string]any{"arr": []int{10}}, ExpectedOutput: []int{10}},
		env.TestCase{InputData: map[string]any{"arr": []int{}}, ExpectedOutput: []int{}},
		env.TestCase{InputData: map[string]any{"arr": []int{5, 4, 3, 2, 1}}, ExpectedOutput: []int{1, 2, 3, 4, 5}},
		env.TestCase{InputData: map[string]any{"arr": []int{1, 2, 3, 4, 5}}, ExpectedOutput: []int{1, 2, 3, 4, 5}},
	},
}
