package constructBinaryTreeFromPreorderAndInorderTraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := TreeNode{Val: preorder[0]}
	leftLen := 0
	for i := range preorder {
		if preorder[0] == inorder[i] {
			leftLen = i
			break
		}
	}
	root.Left = buildTree(preorder[1:leftLen+1], inorder[:leftLen])
	root.Right = buildTree(preorder[leftLen+1:], inorder[leftLen+1:])
	return &root
}
