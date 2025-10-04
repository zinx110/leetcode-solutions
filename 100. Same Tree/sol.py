from typing import Optional

class TreeNode:
    def __init__(self, val: int = 0, left: 'Optional[TreeNode]' = None, right: 'Optional[TreeNode]' = None):
        self.val = val 
        self.left = left 
        self.right = right 

class Solution: # dfs
    def isSameTree(self, p: Optional[TreeNode], q: Optional[TreeNode]) -> bool:

        def dfs(t1:Optional[TreeNode], t2: Optional[TreeNode]) -> bool:
            if not t1 and not t2:
                return True
            if not t1 or not t2:
                return False
            if t1.val != t2.val:
                return False
            left = dfs(t1.left, t2.left)
            right = dfs(t1.right, t2.right)
            return left and right 
        return dfs(p, q)



class Solution2: #bfs
    def isSameTree(self, p: Optional[TreeNode], q: Optional[TreeNode]) -> bool:

        stack = [[p, q]]
        while len(stack) > 0 :
            a, b = stack.pop()
            if a or b:
                if not a or not b:
                    return False
                if a.val != b.val:
                    return False
                stack.append([a.left, b.left])
                stack.append([a.right, b.right])
        return True
        