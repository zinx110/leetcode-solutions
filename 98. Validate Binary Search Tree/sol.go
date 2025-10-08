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

// BFS
func isValidBSTBFS(root *TreeNode) bool {
	que := deque{}
	que.append(root, nil, nil)
	for len(que.nodes) > 0 {
		node, low, high := que.popleft()
		if low != nil && *low >= node.Val {
			return false
		}
		if high != nil && *high <= node.Val {
			return false
		}
		if node.Left != nil {
			que.append(node.Left, low, &node.Val)
		}
		if node.Right != nil {
			que.append(node.Right, &node.Val, high)
		}

	}

	return true

}

type deque struct {
	nodes []que
}

type que struct {
	node *TreeNode
	low  *int
	high *int
}

func (q *deque) init() *deque {
	q.nodes = make([]que, 0)
	return q
}

func (q *deque) append(node *TreeNode, low, high *int) {

	q.nodes = append(q.nodes, que{node, low, high})
}

func (q *deque) popleft() (*TreeNode, *int, *int) {
	if len(q.nodes) > 0 {
		val := q.nodes[0]
		q.nodes = q.nodes[1:]
		return val.node, val.low, val.high
	}
	return nil, nil, nil
}
