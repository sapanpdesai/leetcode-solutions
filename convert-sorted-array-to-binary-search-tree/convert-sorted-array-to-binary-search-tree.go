/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
    if len(nums) == 0{
        return nil
    }
    middle := len(nums)/2
    return &TreeNode{Val: nums[middle], Left:sortedArrayToBST(nums[0:middle]), Right:sortedArrayToBST(nums[middle+1:len(nums)])}
}