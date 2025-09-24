func reverse(x int) int {

    min := -2147483648
    max := 2147483647

    if  !(min <= x  && x<= max) {
        return 0
    } 
    numberString := strconv.Itoa(x)
    numberStringSlice := strings.Split(numberString, "")

    var finalNumberString string 
    var isNegetive bool

    for i:=len(numberStringSlice) -1  ; i>=0 ; i-- {
        if i== len(numberStringSlice) -1 && numberStringSlice[i] == "0" {
                continue 
        }else if i == 0 && numberStringSlice[i] == "-" {
                isNegetive = true 
                continue 
        }else {
            finalNumberString += numberStringSlice[i]
        }
    }
    finalNumberInt , _ := strconv.Atoi(finalNumberString)
    if isNegetive{
        finalNumberInt = finalNumberInt - (2 * finalNumberInt)
    }

    if  !(min <= finalNumberInt  && finalNumberInt<= max) {
        return 0
    } 

    return finalNumberInt
}