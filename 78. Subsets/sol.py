from cgi import test
from collections import deque
import json
import os
from typing import Optional, List 


class Solution:
    # DFS
    def subsets(self, nums: List[int]) -> List[List[int]]:
        res = [[]]
        curr = []
        def dfs(j: int):
            if j > len(nums):
                return 
            for i in range(j, len(nums)):
                curr.append(nums[i])
                dfs(i + 1)
                res.append(curr.copy())
                curr.pop()
        dfs(0)
        return res

    # BFS
    def subsetsBFS(self, nums: List[int]) -> List[List[int]]:
        res = [[]]
        for num in nums:
            curr = []
            for subset in res:
                curr.append(subset + [num])
            res.extend(curr)
        return res






if __name__ == "__main__":
    current_filename = os.path.abspath(__file__)
    current_filepath = os.path.dirname(current_filename)
    test_filename = "testcase.json"
    pathToFile = os.path.join(current_filepath, test_filename)

    fileData = open(pathToFile)
    testCases = json.load(fileData)
    sol = Solution()
    passNum = 0
    for test in testCases["tests"]:
        inputs = test["input"]["nums"]
        expectedOutput = test["expectedOutput"]
        output = sol.subsetsBFS(inputs)
        if sorted(map(sorted, output)) == sorted(map(sorted, expectedOutput)):
            passNum += 1
            print("passed")
        else:
            print("failed")
        print(
            "\ninput:      ", inputs, 
            "\noutput:     ", output, 
            "\nexpected:   ", expectedOutput,
            "\n-------------------"
        )
    print("Total tests: ", len(testCases["tests"]), ", Passed: ", passNum, ", Failed: ", len(testCases["tests"]) - passNum)
