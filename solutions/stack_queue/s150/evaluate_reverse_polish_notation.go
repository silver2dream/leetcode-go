package s150

import "strconv"

type Calc func(a, b int) int

func evalRPN(tokens []string) int {
	stack := make([]int, 0, len(tokens))
	for _, tok := range tokens {
		switch tok {
		case "+", "-", "*", "/":
			a := stack[len(stack)-1] // 先彈出的是 a（右操作數）
			stack = stack[:len(stack)-1]
			b := stack[len(stack)-1] // 再彈出的是 b（左操作數）
			stack = stack[:len(stack)-1]
			var c int
			switch tok {
			case "+":
				c = b + a
			case "-":
				c = b - a
			case "*":
				c = b * a
			case "/":
				c = b / a // Go 整數除法朝 0 取整，符合題意
			}
			stack = append(stack, c)
		default:
			v, _ := strconv.Atoi(tok) // 題目保證有效
			stack = append(stack, v)
		}
	}
	return stack[len(stack)-1]
}
