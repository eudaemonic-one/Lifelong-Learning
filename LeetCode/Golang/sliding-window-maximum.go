func maxSlidingWindow(nums []int, k int) []int {
    res := make([]int, 0)
    if len(nums) == 0 || k == 0 {
        return res
    }
    deque := make([]int, 0)
    max_idx := 0
    for i := 0; i < k; i++ {
        // remove indexes of elements not in sliding window
        if len(deque) > 0 && deque[0] == i-k {
            deque = deque[1:]
        }
        // remove from deq indexes of all elements 
        // which are smaller than current element nums[i]
        for len(deque) > 0 && nums[i] > nums[deque[len(deque)-1]] {
            deque = deque[:len(deque)-1]
        }
        deque = append(deque, i)
        if nums[i] > nums[max_idx] {
            max_idx = i
        }
    }
    res = append(res, nums[max_idx])
    for i := k; i < len(nums); i++ {
        if len(deque) > 0 && deque[0] == i-k {
            deque = deque[1:]
        }
        for len(deque) > 0 && nums[i] > nums[deque[len(deque)-1]] {
            deque = deque[:len(deque)-1]
        }
        deque = append(deque, i)
        res = append(res, nums[deque[0]])
    }
    return res
}
