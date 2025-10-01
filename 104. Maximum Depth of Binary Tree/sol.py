
class TreeNode: 
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

class Solution: # recursive
    def maxDepth(self, root: Optional[TreeNode]) -> int:
        if root is None:
            return 0
        
        depth = 1
        depth += max(self.maxDepth(root.left), self.maxDepth(root.right))
        return depth
        

class Solution2: # iterative
    def maxDepth(self, root: Optional[TreeNode]) -> int:
        res = 0
        stack = [[root, 1]]
        while stack:
            curr, depth = stack.pop()

            if curr: 
                res = max(depth, res)
                stack.append([curr.left, depth + 1])
                stack.append([curr.right, depth + 1])
        return res