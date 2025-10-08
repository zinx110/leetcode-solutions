from Typing import Optional 

class TreeNode:
    def __inti__(self, val: int = 0, left : 'Optional[TreeNode]' = None, right: 'Optional[TreeNode]' = None):
        self.val = val 
        self.left = left 
        self.right = right 


class Solution:
    def isValidBST(self, root: Optional[TreeNode]) -> bool:

        def dfs(node: Optional[TreeNode], highVal: Optional[int] = None, lowVal: Optional[int] = None):
            if not node:
                return True
            if lowVal != None and node.val <= lowVal:
                return False
            if highVal != None and highVal <= node.val:
                return False
            return dfs(node.left, highVal = node.val, lowVal = lowVal) and dfs(node.right, highVal = highVal, lowVal = node.val)

        return dfs(root)


from collections import deque

class Solution2: #bfs
    def isValidBST(self, root: Optional[TreeNode]) -> bool:
        
        que = deque()
        que.append((root, None, None))
        while len(que) > 0:
            node, low, high = que.popleft()
            if low != None and low >= node.val:
                return False
            if high != None and high <= node.val:
                return False
            if node.left != None:
                que.append((node.left, low, node.val))
            if node.right != None:
                que.append((node.right, node.val, high))
        return True
            