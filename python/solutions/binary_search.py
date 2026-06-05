from environment import Solver

class BinarySearchSolver(Solver):
    def solve(self, **input_data):
        """
        Input expected: lista=list sorted ascendent, target=value to search
        Output: index of value to search in list or -1
        """
        lista = input_data.pop("lista")
        target = input_data.pop("target")

        low = 0
        high = len(lista) - 1

        while low <= high:
            mid = (low + high) // 2
            guess = lista[mid]

            if guess == target:
                return mid  # index of position

            if guess > target:
                high = mid-1
            else:
                low = mid+1

        return -1

