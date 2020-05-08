func evalRPN(tokens []string) int {
    numStack := make([]int, 0)
    tmpResult := 0
    for i := range tokens {
        num, err := strconv.Atoi(tokens[i])
        if err == nil {
            numStack = append(numStack, num)
        } else {
            num2, num1 := numStack[len(numStack)-1], numStack[len(numStack)-2]
            numStack = numStack[:len(numStack)-2]
            switch tokens[i] {
                case "+":
                    tmpResult = num1 + num2
                case "-":
                    tmpResult = num1 - num2
                case "*":
                    tmpResult = num1 * num2
                case "/":
                    tmpResult = num1 / num2
                default:
                    tmpResult = 0
            }
            numStack = append(numStack, tmpResult)
        }
    }
  return numStack[0]
}
