import time

from dataclasses import dataclass
from abc import ABC

@dataclass
class TestCase:
    """
    Represents a single test case for a coding problem,
    containing input data and the expected output.
    """
    input_data: dict
    expected_output: any


class Problem(ABC):
    """
    Represents a coding problem with its details and test cases.
    """
    def __init__(self, name, description, test_cases):
        self.name = name
        self.description = description
        self.test_cases = test_cases

    def set_solver(self, solver):
        self.solver = solver

    def run_tests(self):
        """
        Runs the solver against all test cases and checks if the output matches the expected output.
        """
        print(f"\nRunning test cases for problem '{self.name}' with '{self.solver.name}'")
        times = []
        passed = 0

        for i, test_case in enumerate(self.test_cases):
            start_time = time.perf_counter()

            result = self.solver.solve(**test_case.input_data)

            end_time = time.perf_counter()
            exec_time = (end_time - start_time) * 1000  # ms
            times.append(exec_time)

            if result == test_case.expected_output:
                print(f"[{exec_time:.3f} ms] Test case {i + 1} passed")
                passed += 1
            else:
                print(f"[{exec_time:.3f} ms] Test case {i + 1} failed: expected {test_case.expected_output}, got {result}")

        print(f"Passed {passed}/{len(self.test_cases)} test cases")
        print(f"Average execution time: {sum(times) / len(times):.3f} ms")
