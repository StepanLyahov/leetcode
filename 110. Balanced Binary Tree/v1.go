/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func isBalanced(root *TreeNode) bool {
    if root == nil {
        return true
    }

    leftDeep := maxDeep(root.Left)
    rightDeep := maxDeep(root.Right)

    return mod(leftDeep - rightDeep) < 2 && isBalanced(root.Left) && isBalanced(root.Right)
}

func mod(value int32) int32 {
    if value < 0 {
        return -value
    }

    return value
}

func maxDeep(root *TreeNode) int32 {
    if root == nil {
        return 0
    }

    leftDeep := maxDeep(root.Left)
    rightDeep := maxDeep(root.Right)
    

    if leftDeep > rightDeep {
        return leftDeep + 1
    }

    return rightDeep + 1
}