package main

/**
Given a binary tree, determine if it is a valid binary search tree (BST).

Assume a BST is defined as follows:

The left subtree of a node contains only nodes with keys less than the node's key.
The right subtree of a node contains only nodes with keys greater than the node's key.
Both the left and right subtrees must also be binary search trees.
Example 1:

Input:
    2
   / \
  1   3
Output: true
Example 2:

    5
   / \
  1   4
     / \
    3   6
Output: false
Explanation: The input is: [5,1,4,null,null,3,6]. The root node's value
             is 5 but its right child's value is 4.
 */

/**
* Definition for a binary tree node.
* type TreeNode struct {
*     Val int
*     Left *TreeNode
*     Right *TreeNode
* }
*/

type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	result, _, _ := checkBST(root)
	return result
}

func checkBST(root *TreeNode) (result bool, lSmallest, rBiggest int) {
	result, lSmallest, rBiggest = true, root.Val, root.Val
	if root.Left != nil {
		lResult, llSmallest, lrBiggest := checkBST(root.Left)
		if !lResult || lrBiggest >= root.Val {
			return false, 0, 0
		}
		lSmallest = llSmallest
	}
	if root.Right != nil {
		rResult, rlSmallest, rrBiggest := checkBST(root.Right)
		if !rResult || rlSmallest <= root.Val {
			return false, 0, 0
		}
		rBiggest = rrBiggest
	}
	return
}