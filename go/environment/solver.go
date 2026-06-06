package environment

type Solver interface {
	Solve(input map[string]any) any
}
