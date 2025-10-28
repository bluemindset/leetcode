func isPowerOfThree(n int) bool {
    nn := int64(n)

    pp := strconv.FormatInt(nn, 3)
    if string(pp[0])=="1"{
        for i:=1 ; i<len(pp);i++{
            if string(pp[i]) != "0"{
                return false
            }
        }
    }else{
        return false
    }
    return true
}
