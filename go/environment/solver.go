package environment

type Solver interface {
	Name() string
	Solve(input map[string]any) any
}

