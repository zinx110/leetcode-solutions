

class TreeNode:
    def __init__(self, val: int = 0, left: 'Optional[TreeNode]' = None, right: 'Optional[TreeNode]' = None):
        self.val = val 
        self.left = left 
        self.right = right 

class Solution:
    def buildTree(self, preorder: List[int], inorder: List[int]) -> Optional[TreeNode]:
        if len(preorder) == 0 :
            return None
        root = TreeNode(preorder[0])
        leftLen = inorder.index(preorder[0])
        root.left = self.buildTree(preorder[1: leftLen + 1], inorder[:leftLen])
        root.right = self.buildTree(preorder[leftLen + 1: ], inorder[leftLen + 1:])
        return root




