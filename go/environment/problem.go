package environment

import "fmt"

type TestCase struct {
	InputData 	   map[string]any
	ExpectedOutput any
}

type Problem struct {
	Name 		string
	TestCases 	[]TestCase
}

func (p *Problem) RunTests(s Solver) {
	fmt.Printf("Running tests for problem: %s\n", p.Name)

	for i, testCase := range p.TestCases {
		fmt.Printf("Test case %d: Input: %v\n", i+1, testCase.InputData)

		output := s.Solve(testCase.InputData)
		fmt.Printf("Output: %d\n", output)

		if output == testCase.ExpectedOutput {
			fmt.Println("Result: PASS")
		} else {
			fmt.Printf("Result: FAIL (Expected: %v)\n", testCase.ExpectedOutput)
		}
		fmt.Println()
	}
}

