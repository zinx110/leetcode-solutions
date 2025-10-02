class TreeNode:
    def __init__(self, val = 0, left = Optional[TreeNode], right = Optional[TreeNode]):
        self.val = val
        self.left = left
        self.right = right

        
class Solution:
    def isBalanced(self, root: Optional[TreeNode]) -> bool:

        def dfs(node: Optional[TreeNode]) -> int:
            if not node:
                return 0
            left = dfs(node.left)
            right = dfs(node.right)
            if left == -1 or right == -1 or abs(left - right) > 1:
                return -1
            return 1 + max(left, right)

        depth = dfs(root)
        if depth == -1:
            return False
        return True

        