type NumMatrix struct {
    sumMatrix [][]int
    rows int
    cols int
}


func Constructor(matrix [][]int) NumMatrix {
    numMatrix := NumMatrix{}
    rows := len(matrix)
    if rows == 0 {
        return numMatrix
    }
    cols := len(matrix[0])
    if cols == 0 {
        return numMatrix
    }
    numMatrix.rows = rows
    numMatrix.cols = cols
    numMatrix.sumMatrix = make([][]int, rows+1)
    for i := 0; i < rows+1; i++ {
        numMatrix.sumMatrix[i] = make([]int, cols+1)
    }
    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            numMatrix.sumMatrix[i+1][j+1] = numMatrix.sumMatrix[i][j+1] + numMatrix.sumMatrix[i+1][j] + matrix[i][j] - numMatrix.sumMatrix[i][j]
        }
    }
    return numMatrix
}


func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
    return this.sumMatrix[row2+1][col2+1] - this.sumMatrix[row1][col2+1] - this.sumMatrix[row2+1][col1] + this.sumMatrix[row1][col1]
}


/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
