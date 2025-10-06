from typing import Optional 
class TreeNode:
    def __init__(self, val: int = 0, left: 'Optional[TreeNode]' = None, right: 'Optional[TreeeNode]' = None):
        self.val = val 
        self.left = val
        self.right = right 
        
class Solution:
    def levelOrder(self, root: Optional[TreeNode]) -> List[List[int]]:
        res = []
        if not root:
            return res

        stack = [[root]]
        while len(stack) > 0 :
            nodes = stack.pop()
            currValues = []
            currStack = []
            for node in nodes:
                currValues.append(node.val)
                if node.left:
                    currStack.append(node.left)
                if node.right:
                    currStack.append(node.right)
            if len(currValues) == 0:
                break
            res.append(currValues)
            stack.append(currStack)
        return res

        