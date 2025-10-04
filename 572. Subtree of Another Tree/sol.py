

class TreeNode:
    def __init__(self, val: int = 0, left: 'Optional[TreeNode]' = None, right: 'Optional[TreeNode]' = None):
        self.val = val 
        self.left = left 
        self.right = right 

class Solution:
    def isSubtree(self, root: Optional[TreeNode], subRoot: Optional[TreeNode]) -> bool:
        
        def isSameTree(p: Optional[TreeNode], q: Optional[TreeNode]) -> bool:
            stack = [[p, q]]
            while len(stack) > 0:
                p, q = stack.pop()
                if p or q:
                    if not p or not q:
                        return False
                    if p.val != q.val:
                        return False
                    stack.append([p.left, q.left])
                    stack.append([p.right, q.right])
            return True


        stack = [root]
        node = root
        while len(stack) > 0 :
            node = stack.pop()
            if node:
                if node.val == subRoot.val:
                    if isSameTree(node, subRoot):
                        return True
                stack.append(node.left)
                stack.append(node.right)
        
        return False
