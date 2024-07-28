package main

import "fmt"

type TreeNode struct {
  Value int
  Left *TreeNode
  Right *TreeNode
}

func main() {
  root := TreeNode{Value: 1}

  root.Left = &TreeNode{Value: 2}
  root.Right = &TreeNode{Value: 3}

  changeValue(&root, 10)

  fmt.Println("Root value: ", root.Value)
}

func changeValue(node *TreeNode, newValue int) {
  node.Value = newValue
}
