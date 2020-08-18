func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0{
		return nil
	}
	return &TreeNode{
		Val: nums[len(nums)/2],
		Left: sortedArrayToBST(nums[:len(nums)/2]),
		Right: sortedArrayToBST(nums[len(nums)/2+1:]),
	}
}

func sortedArrayToBST(nums []int) *TreeNode {

    if len(nums) == 0 {
        return nil
    }
    
    return constructBST(nums, 0, len(nums)-1)
    
}

func constructBST(nums []int, left int, right int) *TreeNode {
    
    if left > right {
        return nil
    }
    
    mid := left + (right - left) / 2
    Node := &TreeNode{
        Val: nums[mid],
        Left: constructBST(nums, left, mid - 1),
        Right: constructBST(nums, mid + 1, right),
    }
    
    return Node
    
}
