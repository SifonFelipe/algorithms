from environment import TestCase, Problem

BINARY_SEARCH_PROBLEM = Problem(
    name="Binary Search",
    description=(
        "Given a sorted array of integers and a target value, return the index of the"
        "target if it exists in the array, otherwise return -1."
    ),
    test_cases=[
        TestCase(input_data={"lista": [-1,0,3,5,9,12], "target": 9}, expected_output=4),
        TestCase(input_data={"lista": [-1,0,3,5,9,12], "target": 2}, expected_output=-1),
        TestCase(input_data={"lista": [-1,0,3,5,9,12], "target": 3}, expected_output=2),
        # Larger test cases
        TestCase(input_data={"lista": [1,2,3,4,5,6,7,8,9,10], "target": 10}, expected_output=9),
    ],
)

