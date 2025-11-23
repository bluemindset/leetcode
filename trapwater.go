func trap(height []int) int {
    maxRight := 0 
    maxLeft :=0 
    //indexes
    r := len(height)-1
    l :=0
    rain := 0 
    for r >= l{
        if maxLeft < maxRight{ 
            if height[l] > maxLeft{
                maxLeft = height[l]
            }else{
                rain+=maxLeft - height[l]
            }
            l+=1
        }else{
            if height[r] > maxRight{
                maxRight = height[r]
            }else{
                rain+=maxRight - height[r]
            }
            r-=1
        }
    }
    return rain
}
