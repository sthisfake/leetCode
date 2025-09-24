func shortestLength(arr []string) (int , string) {
	if len(arr) == 0 {
		return 0 , ""
	}
    var chosenIndex int
	shortest := len(arr[0])
	for index , s := range arr {
		if len(s) < shortest {
			shortest = len(s)
            chosenIndex = index 
		}
	}
	return shortest , arr[chosenIndex]
}




func longestCommonPrefix(strs []string) string {
    
    shortestStringLen , shortestString := shortestLength(strs)
    numberOfStrings := len(strs)
    var isSame bool
    var tempChar byte

    var prefix []byte

    if len(strs) == 1 {
        return strs[0]
    }

    for i:= 0 ; i<shortestStringLen ; i++ {

        if i != 0 && !isSame {
            return string(prefix)
        }

        isSame = false

        if i!=0 {
           prefix = append(prefix,strs[0][i-1])
        }

        for j:= 0; j<numberOfStrings ; j++ {
            if j == 0 {
                tempChar = strs[j][i]
            }else {
                if strs[j][i] == tempChar {
                    isSame = true
                }else {
                    isSame = false
                    break
                }
            }
        }
    }

    if shortestStringLen == 1 && isSame {
        return shortestString
    }

    if isSame {
        prefix = append(prefix,strs[0][shortestStringLen - 1])
    }

    return string(prefix)

}