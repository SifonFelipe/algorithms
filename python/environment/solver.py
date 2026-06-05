from abc import ABC, abstractmethod

class Solver(ABC):
    """
    Abstract base class for solvers of coding problems.
    Each solver must implement the solve method.
    """
    def __init__(self, name):
        self.name = name

    @abstractmethod
    def solve(self, input_data):
        pass
