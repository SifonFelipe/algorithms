package main

import "algorithms/problems"
import "algorithms/solutions"

func main() {
	BinProblem := problems.BinarySearch
	BinSolver := solutions.BinarySearchSolver{}

	BinProblem.RunTests(BinSolver)
}

