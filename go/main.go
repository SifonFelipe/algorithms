package main

import "fmt"
import problems "algorithms/problems"
import solutions "algorithms/solutions"
import env "algorithms/environment"

var PROBLEMS = map[string]env.Problem{
	"BinarySearch": problems.BinarySearch,
	"Sort": problems.Sort,
}

var SOLVERS = map[string]map[string]env.Solver{
	"BinarySearch": {
		"BinarySearchSolver": solutions.BinarySearchSolver{},
	},
	"Sort": {
		"SelectionSort": solutions.SelectionSortSolver{},
	},
}

func main() {
	// select problem and solver, then run tests

	problemNames := make([]string, 0, len(PROBLEMS))
	for problemName := range PROBLEMS {
		problemNames = append(problemNames, problemName)
	}

	println("Select a problem to solve:")
	for i, problemName := range problemNames {
		println(i+1, ":", problemName)
	}

	var problemChoice int
	_, err := fmt.Scan(&problemChoice)

	if err != nil || problemChoice < 1 || problemChoice > len(problemNames) {
		println("Invalid choice, exiting.")
		return
	}

	selectedProblem := problemNames[problemChoice-1]
	println("You selected:", selectedProblem)

	solvers := SOLVERS[selectedProblem]
	solverNames := make([]string, 0, len(solvers))
	for solverName := range solvers {
		solverNames = append(solverNames, solverName)
	}

	println("\nSelect a solver to use:")
	for i, solverName := range solverNames {
		println(i+1, ":", solverName)
	}

	var solverChoice int
	_, err = fmt.Scan(&solverChoice)

	if err != nil || solverChoice < 1 || solverChoice > len(solverNames) {
		println("Invalid choice, exiting.")
		return
	}

	selectedSolver := solverNames[solverChoice-1]
	println("You selected:", selectedSolver)

	problem := PROBLEMS[selectedProblem]
	solver := solvers[selectedSolver]
	problem.RunTests(solver)
}

