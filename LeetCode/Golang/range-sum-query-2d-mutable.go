type NumMatrix struct {
    tree [][]int
    nums [][]int
    m int
    n int
}


func Constructor(matrix [][]int) NumMatrix {
    if len(matrix) == 0 || len(matrix[0]) == 0 {
        return NumMatrix{}
    }
    m, n := len(matrix), len(matrix[0])
    numMatrix := NumMatrix{}
    numMatrix.m = m
    numMatrix.n = n
    numMatrix.tree = make([][]int, m+1)
    for i := 0; i < m+1; i++ {
        numMatrix.tree[i] = make([]int, n+1)
    }
    numMatrix.nums = make([][]int, m)
    for i := 0; i < m; i++ {
        numMatrix.nums[i] = make([]int, n)
    }
    for i:= 0; i < m; i++ {
        for j := 0; j < n; j++ {
            numMatrix.Update(i, j, matrix[i][j])
        }
    }
    return numMatrix
}


func (this *NumMatrix) Update(row int, col int, val int)  {
    delta := val - this.nums[row][col]
    this.nums[row][col] = val
    for i := row+1; i <= this.m; i += i&(-i) {
        for j := col+1; j <= this.n; j += j&(-j) {
            this.tree[i][j] += delta
        }
    }
}


func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
    return sum(this.tree,row2+1,col2+1)+sum(this.tree,row1,col1)-sum(this.tree,row1,col2+1)-sum(this.tree,row2+1,col1)
}


func sum(tree [][]int, row, col int) int {
    sum := 0
    for i := row; i > 0; i -= i&(-i) {
        for j := col; j > 0; j -= j&(-j) {
            sum += tree[i][j]
        }
    }
    return sum
}


/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * obj.Update(row,col,val);
 * param_2 := obj.SumRegion(row1,col1,row2,col2);
 */
