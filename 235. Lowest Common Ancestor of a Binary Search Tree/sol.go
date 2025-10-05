package lowestCommonAncestorOfABinarySearchTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
	for root != nil && (root.Val > p.Val && root.Val > q.Val) || (root.Val < p.Val && root.Val < q.Val) {
		if root.Val > p.Val {
			root = root.Left
		} else {
			root = root.Right
		}
	}
	return root

}
