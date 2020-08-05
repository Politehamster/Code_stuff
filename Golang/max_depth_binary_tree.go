baseline: if root == nil {return 0}
recursive : count = 1 + Max(maxDepth(root.Left),maxDepth(root.Right))

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxDepth(root *TreeNode) int {
    if root == nil{
        return 0
    }
    
    count := 1 + Max(maxDepth(root.Left),maxDepth(root.Right))
    return count
    
}

func Max(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}

