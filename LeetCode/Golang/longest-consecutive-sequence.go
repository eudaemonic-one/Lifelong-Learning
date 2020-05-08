func longestConsecutive(nums []int) int {
    set := make(map[int]int, len(nums))
    var count, ans = 0, 0
    for i := range nums {
        set[nums[i]] = 1
    }
    for key := range set {
        _, ok := set[key-1]
        if !ok {
            count = 1
            for true {
                _, ok := set[key+1]
                if ok {
                    count += 1
                    key += 1
                } else {
                    break
                }
            }
            if (count > ans) {
                ans = count
            }
        }
    }
    return ans
}