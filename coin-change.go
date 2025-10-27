func coinChange(coins []int, amount int) int {
    if amount<=0{
        return 0
    }
    dp := make([]int, amount+1)
    for i,_ := range dp{
        dp[i] = 99999999
    }
    dp[0] = 1

    for i:=1 ; i< amount+1; i++{
        for _,v := range coins{
            diff := i-v
            ncoins := 0
            if diff>0{
                ncoins = dp[diff]+1 
                dp[i] = min(dp[i], ncoins)
            }
            if diff == 0{
                ncoins = dp[diff]
                dp[i] = min(dp[i], ncoins)
            }
        }
    }
    if dp[amount] == 99999999{
        return -1
    }
    return dp[amount]
}

func print(dp []int){
    for i,v := range dp{
        fmt.Println(i," ",v)
    }
}


func min(a int, b int) int{
    if a > b{
        return b
    }
    return a
}
