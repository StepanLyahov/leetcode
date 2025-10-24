/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func isSymmetric(root *TreeNode) bool {
    return compareReq(root.Left, root.Right)
}

func compareReq(l *TreeNode, r *TreeNode) bool {
    if l == nil && r == nil {
        return true
    }

    if l == nil || r == nil {
        return false
    }

    if l.Val != r.Val {
        return false
    }

    return compareReq(l.Left, r.Right) && compareReq(l.Right, r.Left)
}