func summaryRanges(nums []int) []string {
    res := make([]string, 0)
    for i, j := 0, 0; j < len(nums); j++ {
        if j + 1 < len(nums) && nums[j+1] == nums[j] + 1 {
            continue
        }
        if i == j {
            res = append(res, strconv.Itoa(nums[i]))
        } else {
            res = append(res, strconv.Itoa(nums[i]) + "->" + strconv.Itoa(nums[j]))
        }
        i = j + 1
    }
    return res
}
