 func twoSum(nums []int, target int) []int {
        seen := make(map[int]int)

        for i, n := range nums {
                complement := target - n
                if j, ok := seen[complement]; ok {
                        return []int{j, i}
                }
                seen[n] = i
        }

        return []int{}
  }