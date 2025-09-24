func lengthOfLongestSubstring(s string) int {
	var max int
	var record int 
    mapString := make(map[byte]bool)

    if len(s) == 0 {
        return 0
    }

    if len(s) == 1 {
        return 1
    }

	for i:=0; i<len(s); i++ {
        mapString = make(map[byte]bool)
        mapString[s[i]] = true
        if record > max {
            max = record 
        }
        record = 1 
		for j:=i+1; j<len(s); j++ {
            _ , ok := mapString[s[j]]
			if !ok {
			record++
            mapString[s[j]] = true
			}else{
				if record > max {
					max = record 
				}
				record = 1
                break
			}
		}
	}
	return max
}
