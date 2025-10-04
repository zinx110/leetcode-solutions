package subtreeofanothertree 

type TreeNode struct {
	Val int 
	Left *TreeNode 
	Right *TreeNode 
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	stack := []*TreeNode {root}
    var node *TreeNode
    for len(stack) > 0 {
        index := len(stack) - 1
        node = stack[index]
        stack = stack[:index]
        if node != nil {
            if node.Val == subRoot.Val {
                if isSameTree(node, subRoot) { return true}
            }
            stack = append(stack, node.Left, node.Right)
        }
    }
    return false
}


func isSameTree(p *TreeNode, q *TreeNode) bool {
    stack := [][]*TreeNode { []*TreeNode {p, q}}
    for len(stack) > 0 {
        index := len(stack) - 1
        a, b := stack[index][0], stack[index][1]
        stack = stack[:index]
        if a != nil || b != nil {
            if a == nil || b == nil {
                return false
            }
            if a.Val != b.Val {
                return false
            }
            stack = append(stack, []*TreeNode {a.Left, b.Left}, []*TreeNode {a.Right, b.Right})
        }
    }
    return true
}