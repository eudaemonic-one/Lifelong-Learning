func findMissingRanges(nums []int, lower int, upper int) []string {
    l, u := lower, lower
    ranges := make([]string, 0)
    for _, n := range nums {
        if n < lower || n < l {
            continue
        } else if n == l {
            l = n + 1
            u = l
        } else if n == l+1 {
            ranges = append(ranges, strconv.Itoa(l))
            l = n + 1
            u = l
        } else {
            u = n - 1
            ranges = append(ranges, strconv.Itoa(l) + "->" + strconv.Itoa(u))
            l = n + 1
            u = l
        }
    }
    if l == upper {
        ranges = append(ranges, strconv.Itoa(l))
    } else if l < upper {
        ranges = append(ranges, strconv.Itoa(l) + "->" + strconv.Itoa(upper))
    }
    return ranges
}
