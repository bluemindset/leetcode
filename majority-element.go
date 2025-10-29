func majorityElement(nums []int) int {
    
    c := 0 
    m := nums[0]

    for _,v:= range nums {
        if c == 0 {
            m = v 
            c++
        }else if v == m{
            c++
        }else{
            c--
        }
    }
    return m 

}
// O(1)space , O(n) time - Boyer–Moore majority vote algorithm
