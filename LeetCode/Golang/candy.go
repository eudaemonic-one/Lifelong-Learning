func candy(ratings []int) int {
    if (len(ratings) == 1) {
        return 1
    }
    candies := make([]int, len(ratings))
    sum := 0
    for i := range candies {
        candies[i] = 1
    }
    for i := 1; i < len(ratings); i++ {
        if (ratings[i] > ratings[i-1]) {
            candies[i] = candies[i-1] + 1
        }
    }
    for i := len(ratings)-1; i > 0; i-- {
        if (ratings[i-1] > ratings[i]) {
            if candies[i-1] < candies[i] + 1 {
                candies[i-1] = candies[i] + 1
            }
        }
    }
    for i := range candies {
        sum += candies[i]
    }
    return sum
}
