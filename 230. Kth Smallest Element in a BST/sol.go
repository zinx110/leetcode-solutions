package kthSmallestElementInABST

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	count := 0
	var result int
	found := false

	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil || found {
			return
		}
		inorder(node.Left)
		count++
		if count == k {
			result = node.Val
			found = true
			return
		}
		inorder(node.Right)
	}
	inorder(root)
	return result
}
