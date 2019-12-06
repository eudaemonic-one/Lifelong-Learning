func threeSumSmaller(nums []int, target int) int {
    res := 0
    sort.Ints(nums)
    for i := 0 ; i < len(nums)-2; i++ {
        for j, k := i+1, len(nums)-1; j < k; {
            sum := nums[i] + nums[j] + nums[k]
            if sum < target {
                res += k - j
                j++
            } else {
                k--
            }
        }
    }
    return res
}

/*
func threeSumSmaller(nums []int, target int) int {
    if len(nums) < 3 {
        return 0
    }
    pairs := make(map[[2]int]int)
    res := 0
    for i := 0 ; i < len(nums)-2; i++ {
        for j := i+1; j < len(nums)-1; j++ {
            pairs[[2]int{i, j}] = nums[i] + nums[j]
        }
    }
    for pair, sum := range pairs {
        j := pair[1]
        for k := j+1; k < len(nums); k++ {
            if nums[k] + sum < target {
                res++
            }
        }
    }
    return res
}
*/
