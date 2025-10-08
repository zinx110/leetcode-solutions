package validateBinarySearchTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	var dfs func(node *TreeNode, highVal, lowVal *int) bool
	dfs = func(node *TreeNode, highVal, lowVal *int) bool {
		if node == nil {
			return true
		}
		if lowVal != nil && *lowVal >= node.Val {
			return false
		}
		if highVal != nil && node.Val >= *highVal {
			return false
		}
		return dfs(node.Left, &node.Val, lowVal) && dfs(node.Right, highVal, &node.Val)
	}
	return dfs(root, nil, nil)
}
