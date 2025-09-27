func isValid(s string) bool {
    a := map[rune]int{
        '(' :  1,
        '[' :  2,
        '{' :  3,
        ')' : -1,
        ']' : -2,
        '}' : -3,
    }

    stack := []int{}

    for _, char := range s {
        val := a[char]

        if val > 0 {
            stack = append(stack, val)
        } else {
            if len(stack) == 0 {
                return false
            }
            top := stack[len(stack)-1]
            if top+val != 0 { 
                return false
            }
            stack = stack[:len(stack)-1]
        }
    }

    return len(stack) == 0
}
