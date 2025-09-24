func romanToInt(s string) int {
    mapped := map[rune]int{
        'I' : 1 ,
        'V' : 5 ,
        'X' : 10 ,
        'L' : 50 , 
        'C' : 100 , 
        'D' : 500 , 
        'M' : 1000,
    }

    var value int 
    var skipFlag bool
    for index , letter := range s {

        if skipFlag {
            skipFlag = false
            continue
        }
        if index != len(s) -1{
        if letter == 'I' {
            if s[index+1] == 'V' {
                value += 4
                skipFlag = true
            }

            if s[index+1] == 'X' {
                value += 9
                skipFlag = true
            }
        }

        if letter == 'X' {
            if s[index+1] == 'L' {
                value += 40
                skipFlag = true
            }

            if s[index+1] == 'C' {
                value += 90
                skipFlag = true
            }
        }

        if letter == 'C' {
            if s[index+1] == 'D' {
                value += 400
                skipFlag = true
            }

            if s[index+1] == 'M' {
                value += 900
                skipFlag = true
            }
        }
        } 


        if skipFlag {
            continue
        }

        value += mapped[letter]

    }
    return value
}