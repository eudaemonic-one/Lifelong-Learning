func moveZeroes(nums []int)  {
    lastNonZeroIdx := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] != 0 {
            nums[lastNonZeroIdx] = nums[i]
            lastNonZeroIdx++
        }
    }
    for i := lastNonZeroIdx; i < len(nums); i++ {
        nums[i] = 0
    }
}
