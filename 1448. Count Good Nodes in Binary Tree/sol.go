package countGoodNodesInBinaryTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func goodNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := 0

	stack := []pair{pair{root, root.Val}}
	for len(stack) > 0 {
		index := len(stack) - 1
		curr := stack[index]
		stack = stack[:index]
		if curr.Node != nil {
			if curr.Node.Val >= curr.Highest {
				res++
			}
			stack = append(stack, pair{curr.Node.Left, max(curr.Node.Val, curr.Highest)}, pair{curr.Node.Right, max(curr.Node.Val, curr.Highest)})
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

type pair struct {
	Node    *TreeNode
	Highest int
}
