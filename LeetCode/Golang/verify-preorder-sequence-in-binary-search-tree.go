func verifyPreorder(preorder []int) bool {
    low := -1
    cur := -1
    for i := 0; i < len(preorder); i++ {
        if preorder[i] < low {
            return false
        }
        for cur >= 0 && preorder[i] > preorder[cur] {
            low = preorder[cur]
            cur--
        }
        cur++
        preorder[cur] = preorder[i]
    }
    return true
}
