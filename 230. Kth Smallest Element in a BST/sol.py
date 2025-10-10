from Typing import Optional
class TreeNode:
    def __init__(self, val: int = 0, left: 'Optional[TreeNode]' = None, right: 'Optional[TreeNode]' = None):
        self.val = val
        self.left = left 
        self.right = right 
    
class Solution:
    def kthSmallest(self, root: Optional[TreeNode], k: int) -> int:
        count = 0
        result = None

        def inorder(node: Optional[TreeNode]):
            nonlocal count, result
            if not node or result is not None:
                return 
            inorder(node.left)
            count += 1
            if count == k:
                result = node.val
                return 
            inorder(node.right)
        inorder(root)
        return result
        
        from Typing import Optional
class TreeNode:
    def __init__(self, val: int = 0, left: 'Optional[TreeNode]' = None, right: 'Optional[TreeNode]' = None):
        self.val = val
        self.left = left 
        self.right = right 
    
class Solution:
    def kthSmallest(self, root: Optional[TreeNode], k: int) -> int:
        count = 0
        result = None

        def inorder(node: Optional[TreeNode]):
            nonlocal count, result
            if not node or result is not None:
                return 
            inorder(node.left)
            count += 1
            if count == k:
                result = node.val
                return 
            inorder(node.right)
        inorder(root)
        return result
        
        