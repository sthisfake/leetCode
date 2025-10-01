func remove(slice []int, s int) []int {
    return append(slice[:s], slice[s+1:]...)
}

func removeDuplicates(nums []int) int {
    temp := []int{}
    existMap := make(map[int]bool)
    for i := 0; i < len(nums); i++ {
        if _, ok := existMap[nums[i]]; !ok {
            existMap[nums[i]] = true
            temp = append(temp, nums[i])
        }
    }
    copy(nums, temp)
    return len(temp)
}

