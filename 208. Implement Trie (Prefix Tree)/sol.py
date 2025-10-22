# leetcode 208

import json 
import os 

class TrieNode:
    def __init__(self):
        self.children = {}
        self.endOfWord = False
class Trie:

    def __init__(self):
        self.root = TrieNode()
        

    def insert(self, word: str) -> None:
        curr = self.root
        for c in word:
            if c not in curr.children:
                curr.children[c] = TrieNode()
            curr = curr.children[c]
        curr.endOfWord = True

    def search(self, word: str) -> bool:
        curr = self.root 
        for c in word:
            if c not in curr.children:
                return False
            curr = curr.children[c]
        return curr.endOfWord

    def startsWith(self, prefix: str) -> bool:
        curr = self.root 
        for c in prefix:
            if c not in curr.children:
                return False 
            curr = curr.children[c]
        return True


if __name__ == "__main__":
    current_filename = os.path.abspath(__file__)
    current_dir = os.path.dirname(current_filename)
    testcase_path = os.path.join(current_dir, "testcase.json")

    testcases = json.load(open(testcase_path))
    tests = testcases["tests"]
    passNum = 0 
    trie = Trie()
    for test in tests:
        input = test["input"]
        commands = input["commands"]
        arguments = input["arguments"]
        expectedOutput = test["expectedOutput"]
        output = []
        failed = False
        for i in range(len(commands)):
            res = None
            if commands[i] == "Trie":
                trie = Trie()
            elif commands[i] == "insert":
                arg = arguments[i][0] if isinstance(arguments[i], list) else arguments[i]
                trie.insert(arg)
            elif commands[i] == "search":
                arg = arguments[i][0] if isinstance(arguments[i], list) else arguments[i]
                res = trie.search(arg)
            elif commands[i] == "startsWith":
                arg = arguments[i][0] if isinstance(arguments[i], list) else arguments[i]
                res = trie.startsWith(arg)
            output.append(res)
        if output != expectedOutput:
            failed = True
            print(f"Test failed")
            break
        else:
            passNum += 1
            print(f"Test passed")
     
        print(f"commands: {commands}")
        print(f"arguments: {arguments}")
        print(f"output: {output}")
        print(f"expectedOutput: {expectedOutput}")
        print("--------------------------------")
    print(f"Total tests: {len(tests)}, \nPassed tests: {passNum} \nFailed tests: {len(tests) - passNum}")