func twoSum(nums []int, target int) []int {
    for index , _ := range nums {
        for i := index + 1 ; i < len(nums) ; i++ {
            if target == (nums[index] + nums[i]) {
                return []int{index ,  i}
            }
        }
    }
    return []int{}
}
