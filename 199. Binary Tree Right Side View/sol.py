class TreeNode:
    def __init__(self, val: int = 0, left: 'Optional[TreeNode]' = None, right: 'Optional[TreeNode]' = None):
        self.val = val 
        self.left = left 
        self.right = right 


class Solution:
    def rightSideView(self, root: Optional[TreeNode]) -> List[int]:
        res = []
        if not root:
            return res
        queue = [root]
        while len(queue) > 0 :
            queLen = len(queue)
            res.append(queue[queLen - 1].val)
            for i in range(queLen):
                node = queue[i]
                if node.left:
                    queue.append(node.left)
                if node.right:
                    queue.append(node.right)
            queue = queue[queLen:]
        return res