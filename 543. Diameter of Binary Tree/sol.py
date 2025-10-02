class TreeNode:
    def __init__(self, val = 0 , left = Optional[TreeNode], right = Optional[TreeNode]):
        self.val = val
        self.left = left 
        self.right = right 
        


class Solution:
    def diameterOfBinaryTree(self, root: Optional[TreeNode]) -> int:
        
        
        def dfs(node: Optional[TreeNode]) -> (int, int):
            if not node:
                return 0, 0
            depth = 1
            left, maxLeft = dfs(node.left)
            right, maxRight = dfs(node.right)

            diameter = left + right
            print(left, maxLeft, right, maxRight)
            return 1 + max(left, right) , max(diameter, maxLeft, maxRight)

        depth, maxDepth = dfs(root)

        return maxDepth