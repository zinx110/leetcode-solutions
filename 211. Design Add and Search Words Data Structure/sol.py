




import os
import json
from typing import Optional



class TrieNode:
    def __init__(self):
        self.children = {}
        self.endOfWord = False

class WordDictionary:

    def __init__(self):
        self.root = TrieNode()
        

    def addWord(self, word: str) -> None:
        curr = self.root 
        for c in word:
            if c not in curr.children:
                curr.children[c] = TrieNode()
            curr = curr.children[c]
        curr.endOfWord = True

    def searchFromNode(self, word: str, curr: Optional[TrieNode]) -> bool:
        for i in range(len(word)):
            c = word[i]
            # print(c, curr.children)
            if c in curr.children:
                curr = curr.children[c]
                continue
            
            if c == "." and len(curr.children) != 0:
                # print("len of childs ", len(curr.children), curr.children)
                valid = False
                for child in curr.children.values():
                    if self.searchFromNode(word[i + 1:], child):
                        return True
                return False
            else:
                # print("dfas",curr)
                return False
        if curr.endOfWord:
            return True
        return False


    def search(self, word: str) -> bool:
      
        return self.searchFromNode(word, self.root)
            
        


# Your WordDictionary object will be instantiated and called as such:
# obj = WordDictionary()
# obj.addWord(word)
# param_2 = obj.search(word)

def main():
    filename = "testcase.json"
    current_filename = os.path.abspath(__file__)
    # print("file:",__file__, "current_filename:", current_filename)
    current_dir = os.path.dirname(current_filename)
    # print("current_dir:", current_dir)
    path = os.path.join(current_dir, filename)
    jsonData = open(path).read()
    testcases = json.loads(jsonData) 
    # print(testcases)
    passNum = 0
    wordDictionary = WordDictionary()
    tests = testcases["tests"]
    for test in tests:
        input = test["input"]
        operations = input["operations"]
        arguments = input["arguments"]
        expectedOutput = test["expectedOutput"]
        failed = False 
        output = []
        for i in range(len(operations)):
            res = None
            if operations[i] == "WordDictionary":
                wordDictionary = WordDictionary()
            if operations[i] == "addWord":
                wordDictionary.addWord(arguments[i][0])
            elif operations[i] == "search":
                res = wordDictionary.search(arguments[i][0])
            
            output.append(res)
        if output != expectedOutput:
            failed = True 
            print(f"Test failed")
        else:
            print(f"Test passed")
            passNum += 1

        print("inputs :" )
        print(" operations :", operations)
        print(" arguments  :", arguments)
        print(" output     :", output)
        print(" expected   :", expectedOutput)
        print("----------------")
    print("Total tests: ", len(tests),", Passed: ", passNum, ", failed: ", len(tests) - passNum)


if __name__ == "__main__":
    main()