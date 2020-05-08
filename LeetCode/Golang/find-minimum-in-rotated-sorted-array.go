func findMin(nums []int) int {
    l, r := 0, len(nums)-1
    // if there is only one element
    if len(nums) == 1 {
        return nums[0]
    }
    // if there is no rotation
    if nums[r] >= nums[0] {
        return nums[0]
    }
    for l <= r {
        m := (l + r) / 2 // m for turning point
        if nums[m] > nums[m+1] {
            return nums[m+1]
        }
        if nums[m-1] > nums[m] {
            return nums[m]
        }
        if nums[m] > nums[0] {
            l = m + 1
        } else {
            r = m - 1
        }
    }
    return -1
}
