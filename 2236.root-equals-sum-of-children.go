// https://leetcode.com/problems/root-equals-sum-of-children/

// Test cases:
// root = [10,4,6]
// root = [5,3,1]

// Expected results:
// true
// false

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

package leetcode

func checkTree(root *TreeNode) bool {
	return root.Val == root.Left.Val+root.Right.Val
}
