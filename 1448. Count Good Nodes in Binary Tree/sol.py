class TreeNode:
    def __init__(self, val: int = 0, left: 'Optional[TreeNode]'= None, right : 'Optional[TreeNode]' = None):
        self.val = val
        self.left = left 
        self.right = right 
        
class Solution:
    def goodNodes(self, root: TreeNode) -> int:
        res = 0
        if not root:
            return res
        stack = [[root, root.val]]
        while len(stack) > 0 :
            node, highest = stack.pop()
            if node:
                if node.val >= highest:
                    res += 1
                stack.append([node.left, max(node.val, highest)])
                stack.append([node.right, max(node.val, highest)])
        return res