from typing import Optional
class TreeNode:
    def __init__(self, val: int = 0, left: 'Optional[TreeNode]' = None, right: 'Optional[TreeNode]' = None):
        self.val = val
        self.left = left
        self.right = right 


class Solution:
    def lowestCommonAncestor(self, root: 'TreeNode', p: 'TreeNode', q: 'TreeNode') -> 'TreeNode':

        while root and (root.val > p.val and root.val > q.val) or (root.val < p.val and root.val < q.val):
            if root.val > p.val:
                root = root.left
            else:
                root = root.right
        return root
