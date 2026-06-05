package environment

type TestCase struct {
	InputData 	   map[string]any
	ExpectedOutput any
}

type Problem struct {
	Name 		string
	Description string
	TestCases 	[]TestCase
}

func (p *Problem) RunTests(s Solver) {
	// ...	
}

