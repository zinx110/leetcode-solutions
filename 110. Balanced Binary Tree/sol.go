package BalancedBinaryTree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func isBalanced(root *TreeNode) bool {

    var dfs func(node *TreeNode) int 
    dfs = func(node *TreeNode) int {
        if node == nil {
            return 0
        }
        left := dfs(node.Left)
        right := dfs(node.Right)
        if left < 0 || right < 0 || abs(left - right) > 1 {
            return -1
        }
        return 1 + max(left, right)
    }

    depth := dfs(root)

    return depth >= 0 
    
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}