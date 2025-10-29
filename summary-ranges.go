import (
  "fmt"
  "strconv"
)

func summaryRanges(nums []int) []string {
    
    ranges := make([]string, 0)
    if nums == nil || len(nums)==0{
        return ranges
    }
    start := strconv.Itoa(nums[0])
    for i:=0 ; i < len(nums)-1; i++{
        end := ""
        fmt.Println( nums[i+1] , nums[i], start)
        if nums[i+1] - nums[i] != 1{
            end = strconv.Itoa(nums[i])
            v := start + "->" + end
            if end == start{
                v = end
            }
            ranges = append(ranges, v)
            start = strconv.Itoa(nums[i+1])
        }
    }
    end := strconv.Itoa(nums[len(nums)-1])
    v := start + "->" + end
    if end == start{
        v = end
    }
    ranges = append(ranges, v)
    return ranges
}
