// https://leetcode.com/problems/invert-binary-tree/

// Test cases:
// root = [4,2,7,1,3,6,9]
// root = [2,1,3]
// root = []

// Expected results
// [4,7,2,9,6,3,1]
// [2,3,1]
// []

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func invertTree(root *TreeNode) *TreeNode {
  if root == nil {
    return nil
  }

  root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)

  return root
}
