func findDuplicate(nums []int) int {
    // 0 cannot appear as a value in nums
    // nums[0] cannot be part of the cycle
    // traversing the array in this manner from nums[0] is equivalent to traversing a cyclic linked list
    slow, fast := nums[0], nums[0]
    for true {
        slow = nums[slow]
        fast = nums[nums[fast]]
        if slow == fast {
            slow = nums[0]
            for slow != fast {
                slow = nums[slow]
                fast = nums[fast]
            }
            return fast
        }
    }
    return -1
}
