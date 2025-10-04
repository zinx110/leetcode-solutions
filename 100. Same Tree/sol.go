package sameTree



type TreeNode struct {
	Val int
	Left *TreeNode  
	Right *TreeNode 
}


// dfs
func isSameTree(p *TreeNode, q *TreeNode) bool {
    if p == nil && q == nil {
        return true
    }
    if p == nil || q == nil {
        return false
    }
    if p.Val != q.Val {
        return false
    }
    left := isSameTree(p.Left, q.Left)
    right := isSameTree(p.Right, q.Right)
    return left && right 
}

// bfs
func isSameTreeBFS(p *TreeNode, q *TreeNode) bool {

    stack := [][]*TreeNode {[]*TreeNode {p, q}}
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
            stack = append(stack, []*TreeNode {a.Left, b.Left})
            stack = append(stack, []*TreeNode {a.Right, b.Right})
        }
    }


    return true
}
