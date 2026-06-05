import problems
import solutions

if __name__ == "__main__":
    problem = problems.BINARY_SEARCH_PROBLEM
    solver = solutions.BinarySearchSolver(name="Binary Search Solver")

    problem.set_solver(solver)
    problem.run_tests()


