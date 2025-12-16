package p27

// removeElement 移除切片 nums 中所有与val相等的元素
// 返回值即为去重后的新长度
func removeElement(nums []int, val int) int {
	// 双指针：
	// left指向待处理位置，right指向已移除位置
	// （开始所有元素均未处理right指向末尾元素下一个位置）
	left, right := 0, len(nums)
    for left < right {
        if nums[left] == val {
			// 当前元素为待移除元素
			// right自减 指向下一个移除元素位置
			right--
			// 将right指向元素覆盖left指向元素
            nums[left] = nums[right]
			// 注意：末尾覆盖过来元素也可能是val
			// left指针不自增继续检查当前位置
        } else {
			// 当前元素不为val，left自增，检查下一位置
            left++
        }
    }
    return left
}