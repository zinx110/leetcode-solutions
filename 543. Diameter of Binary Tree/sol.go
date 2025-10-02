package diameterOfBinaryTree

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode 

}
func diameterOfBinaryTree(root *TreeNode) int { 
    res := 0 
    
    var dfs func(node *TreeNode) (int) 
    dfs = func(node *TreeNode) (int) {
        if node == nil {
            return 0
        }
        
        leftHeight := dfs(node.Left)
        rightHeight := dfs(node.Right)

        maxDepth := max(leftHeight, rightHeight) + 1
        currDiameter := leftHeight + rightHeight
        res = max(res, currDiameter)

        return maxDepth
    }

    dfs(root)
    return res

}

func max(args ...int) int {
    m := args[0]
    for _, a := range args {
        if a > m {
            m = a
        }
    }
    return m
}