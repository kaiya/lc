/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	numMap := make([int]int, len(nums))
	for i, num := range nums {
		if pos, ok := numMap[target - num]; ok{
			return int[]{pos, i}
		}
		numMap[i] = num
	}
	return nil
}

// @lc code=end

