func majorityElement(nums []int) []int {
    threshold := len(nums) / 3
    cand1, cand2 := 0, 1
    cnt1, cnt2 := 0, 0
    res := make([]int, 0)
    // there will be at most 2 elements appear more than 3/n times
    for i := 0; i < len(nums); i++ {
        num := nums[i]
        if num == cand1 {
            cnt1 += 1
        } else if num == cand2 {
            cnt2 += 1
        } else if cnt1 == 0 {
            cand1 = num
            cnt1 = 1
        } else if cnt2 == 0 {
            cand2 = num
            cnt2 = 1
        } else {
            cnt1 -= 1
            cnt2 -= 1
        }
    }
    cnt1, cnt2 = 0, 0
    for i := 0; i < len(nums); i++ {
        if nums[i] == cand1 {
            cnt1 += 1
        } else if nums[i] == cand2 {
            cnt2 += 1
        }
    }
    if cnt1 > threshold {
        res = append(res, cand1)
    }
    if cnt2 > threshold {
        res = append(res, cand2)
    }
    return res
}
