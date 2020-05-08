func multiply(A [][]int, B [][]int) [][]int {
    if len(A) == 0 || len(A[0]) == 0 {
        return [][]int{}
    }
    if len(B) == 0 || len(B[0]) == 0 {
        return [][]int{}
    }
    p, q, r := len(A), len(A[0]), len(B[0])
    res := make([][]int, p)
    for i := 0; i < p; i++ {
        res[i] = make([]int, r)
    }
    cols := make([][]int, p)
    for i := 0; i < p; i++ {
        elems := make([]int, 0)
        for j := 0; j < q; j++ {
            if A[i][j] != 0 {
                elems = append(elems, j) // column index
                elems = append(elems, A[i][j]) // value
            }
        }
        cols[i] = elems
    }
    for i := 0; i < p; i++ {
        elems := cols[i]
        for k := 0; k < len(elems)-1; k += 2 {
            colA, valA := elems[k], elems[k+1]
            for j := 0; j < r; j++ {
                valB := B[colA][j]
                res[i][j] += valA * valB
            }
        }
    }
    return res
}
