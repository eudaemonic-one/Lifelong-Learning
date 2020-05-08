func diffWaysToCompute(input string) []int {
    operands := make([]int, 0)
    operators := make([]byte, 0)
    for i := 0; i < len(input); {
        if input[i] < '0' || input[i] > '9' {
            operators = append(operators, input[i])
            i++
            continue
        }
        j := i + 1
        for ; j < len(input); j++ {
            if input[j] < '0' || input[j] > '9' {
                break
            }
        }
        num, _ := strconv.Atoi(input[i:j])
        operands = append(operands, num)
        i = j
    }
    return helper(operands, operators, 0, len(operands)-1)
}

func helper(operands []int, operators []byte, i, j int) []int {
    if i == j {
        return []int{operands[i]}
    }
    result := 0
    tmpRes := make([]int, 0)
    for k := i; k < j; k++ {
        first, second := helper(operands, operators, i, k), helper(operands, operators, k+1, j)
        for m := 0; m < len(first); m++ {
            for n := 0; n < len(second); n++ {
                switch operators[k] {
                    case '+':
                        result = first[m] + second[n]
                    case '-':
                        result = first[m] - second[n]
                    case '*':
                        result = first[m] * second[n]
                    default:
                        continue
                }
                tmpRes = append(tmpRes, result)
            }
        }
    }
    return tmpRes
}
