package maximumDepthOfBinaryTree

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// recursive
func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }

    depth := 1
    leftDepth := maxDepth(root.Left)
    rightDepth := maxDepth(root.Right)
    return depth + max(leftDepth, rightDepth)

    
}




// iterative
type Pair struct {
    node *TreeNode
    depth int
}
func maxDepth2(root *TreeNode) int {
    res := 0
    stack := []Pair {Pair{root, 1}}

    for len(stack) > 0 {
        pair := stack[len(stack) - 1]
        stack = stack[:len(stack) - 1]
        if pair.node != nil {
            res = max(res, pair.depth)
            stack = append(stack, Pair{pair.node.Left, pair.depth + 1})
            stack = append(stack, Pair{pair.node.Right, pair.depth + 1})
        }
    }
    return res

}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}