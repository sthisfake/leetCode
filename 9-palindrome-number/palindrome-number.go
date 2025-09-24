
package main

import (
    "strconv"
)

func isPalindrome(x int) bool {
    if x < 0 {
        return false
    }
    if x < 10 {
        return true
    }

    s := strconv.Itoa(x)

    digits := make([]int, len(s))
    for i, ch := range s {
        digits[i] = int(ch - '0')
    }

    for index , _ := range digits {
        if index > len(digits) / 2 {
            break
        } 
        if digits[index] != digits[(len(digits) -1 ) - index] {
            return false
        }
    }
    return true
}
