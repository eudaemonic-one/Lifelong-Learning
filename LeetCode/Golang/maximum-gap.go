func maximumGap(nums []int) int {
    if len(nums) < 2 {
        return 0
    }
    // find max value in the array
    max := nums[0]
    for i := 1; i < len(nums); i++ {
        if nums[i] > max {
            max = nums[i]
        }
    }
    exp := 1
    radix := 10
    aux := make([]int, len(nums))
    // LSD Radix Sort
    for max/exp > 0{
        // Counting Sort
        count := make([]int, radix)
        for i := 0; i < len(nums); i++ {
            count[(nums[i]/exp) % radix] += 1
        }
        for i := 1; i < radix; i++ {
            count[i] += count[i-1]
        }
        for i := len(nums)-1; i >= 0; i-- {
            count[(nums[i]/exp) % radix] -= 1
            aux[count[(nums[i]/exp) % radix]] = nums[i]
        }
        for i := 0; i < len(nums); i++ {
            nums[i] = aux[i]
        }
        exp *= 10
    }
    // find max gap
    var res int
    for i := 0; i < len(nums)-1; i++ {
        if nums[i+1] - nums[i] > res {
            res = nums[i+1] - nums[i]
        }
    }
    return res
}
