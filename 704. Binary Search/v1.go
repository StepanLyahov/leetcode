func search(nums []int, target int) int {
    
    l := 0
    r := len(nums) - 1
    
    for l <= r {
        s := ((r-l)/2) + l    
        
        if target == nums[s] {
            return s
        }
        
        if target < nums[s] {
            r = s - 1
        }
        
        if target > nums[s] {
            l = s + 1
        }
    }
    
    return -1
}
