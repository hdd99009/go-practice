package homework01

import (
	"fmt"
	"sort"
)

// 1. 只出现一次的数字
// 给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
func SingleNumber(nums []int) int {
	// TODO: implement
	m := make(map[int]int)
	for _, i := range nums {
		m[i]++
	}
	for j, k := range m {
		if k == 1 {
			fmt.Printf("只出现一次的数为%d", j)
			return j
		}

	}
	return 0

}

// 2. 回文数
// 判断一个整数是否是回文数
func IsPalindrome(x int) bool {
	numtostr := fmt.Sprint(x)
	for i := 0; i < len(numtostr)/2; i++ {
		if numtostr[i] != numtostr[len(numtostr)-1-i] {
			return false
		}
	}
	// TODO: implement
	return true
}

func IsPalindrome2(x int) bool {
	if x < 0 {
		return false
	} else {
		j := x
		c := 0

		for j > 0 {
			k := j % 10
			j = j / 10
			c = 10*c + k
		}
		if c == x {
			return true
		} else {
			return false
		}
	}
}

// 3. 有效的括号
// 给定一个只包括 '(', ')', '{', '}', '[', ']' 的字符串，判断字符串是否有效
// 用字符串做的栈 开销很大
func IsValid(s string) bool {

	var s1 string
	for i := 0; i < len(s); i++ {

		if s[i] == '(' {
			s1 += string(s[i])
		}
		if s[i] == '[' {
			s1 += string(s[i])
		}
		if s[i] == '{' {
			s1 += string(s[i])
		}

		if s[i] == ')' {
			if len(s1) == 0 || s1[len(s1)-1] != '(' {
				return false
			} else {
				s1 = s1[:len(s1)-1]
			}
		}
		if s[i] == ']' {
			if len(s1) == 0 || s1[len(s1)-1] != '[' {
				return false
			} else {
				s1 = s1[:len(s1)-1]
			}
		}
		if s[i] == '}' {
			if len(s1) == 0 || s1[len(s1)-1] != '{' {
				return false
			} else {
				s1 = s1[:len(s1)-1]
			}
		}
	}
	if s1 == "" {
		return true
	} else {
		return false
	}
}

func IsValid2(s string) bool {
	var s1 []byte
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			s1 = append(s1, s[i])
		}
		if s[i] == ')' {
			if len(s1) == 0 || s1[len(s1)-1] != '(' {
				return false
			} else {
				s1 = s1[:len(s1)-1]
			}
		}
		if s[i] == ']' {
			if len(s1) == 0 || s1[len(s1)-1] != '[' {
				return false
			} else {
				s1 = s1[:len(s1)-1]
			}
		}
		if s[i] == '}' {
			if len(s1) == 0 || s1[len(s1)-1] != '{' {
				return false
			} else {
				s1 = s1[:len(s1)-1]
			}
		}
	}
	if len(s1) == 0 {
		return true
	} else {
		return false
	}
}

// 4. 最长公共前缀
// 查找字符串数组中的最长公共前缀
func LongestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}
	x := ""
	for _, x = range strs {
		if x == "" {
			return ""
		}
	}
	ans := ""
	ans1 := ""
	for i := 0; i < 200; i++ {
		for j := 0; j < len(strs)-1; j++ {
			if i >= len(strs[j+1]) || i >= len(strs[j]) || strs[j][i] != strs[j+1][i] {
				return ans
			}
			ans1 = string(strs[j][i])
		}
		ans += ans1
	}
	return ""
}
func LongestCommonPrefix2(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	for i := 0; i < len(strs[0]); i++ {
		for j := 0; j < len(strs); j++ {
			if i == len(strs[j]) || strs[j][i] != strs[0][i] {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}

// 5. 加一
// 给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一
// 这个方法会整型溢出
func PlusOne(digits []int) []int {
	var num int

	for i := 0; i < len(digits); i++ {
		num = 10*num + digits[i]
	}
	num = num + 1
	var newdigits []int
	{
	}
	for {
		newdigits = append([]int{num % 10}, newdigits...)
		num = num / 10
		if num == 0 {
			break
		}
	}
	// TODO: implement
	return newdigits
}
func PlusOne1(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i] += 1
			return digits
		} else {
			digits[i] = 0
		}
		if i == 0 {
			newddigits := append([]int{1}, digits...)
			return newddigits
		}
	}
	return digits
}

// 6. 删除有序数组中的重复项
// 给你一个有序数组 nums ，请你原地删除重复出现的元素，使每个元素只出现一次，返回删除后数组的新长度。
// 不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。
// 自己想的，会超时
func RemoveDuplicates(nums []int) int {
	for i := 0; i < len(nums)-1; i++ {

		for j := i + 1; j < len(nums); j++ {

			if nums[i] == nums[j] {

				if j < len(nums)-1 {
					x := nums[j]
					nums[j] = nums[j+1]
					nums[j+1] = x
					//fmt.Println(i, j)
					//fmt.Println(nums)
				} else {

				}
			}

		}
		for k := i + 1; k < len(nums)-1; k++ {
			if nums[k] == nums[i] && nums[k+1] != nums[i] {
				i--
				break
			}
		}

	}
	k := len(nums)

	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[j] == nums[i] {
				if j < k {
					k = j
					//fmt.Print(k)
				}
			}
		}
	}
	//fmt.Println(nums)
	return k

}
func RemoveDuplicates2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	i := 0
	j := 1
	for j < len(nums) {
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}
		j++
	}
	return i + 1
}

func RemoveDuplicates3(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	i := 0
	j := 1
	for ; j < len(nums); j++ {
		if nums[j] != nums[i] {
			nums[i+1] = nums[j]
			i += 1
		}
	}
	return i + 1
}

// 7. 合并区间
// 以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
// 请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。
func Merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	for i := 0; i < len(intervals)-1; {
		if intervals[i][1] >= intervals[i+1][0] {
			intervals[i+1][1] = max(intervals[i][1], intervals[i+1][1])
			intervals[i+1][0] = intervals[i][0]
			intervals = append(intervals[0:i], intervals[i+1:]...)
		} else {
			i++
		}
	}
	return intervals
}

func Merge2(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res := [][]int{}
	res = append(res, intervals[0])
	j := 0
	for i := 1; i < len(intervals); i++ {
		if res[j][1] >= intervals[i][0] {
			res[j][1] = max(res[j][1], intervals[i][1])
		} else {
			res = append(res, intervals[i])
			j++
		}
	}
	return res
}

// 8. 两数之和
// 给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那两个整数
func TwoSum(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if target == nums[i]+nums[j] {
				ans := []int{i, j}
				return ans
			}
		}
	}
	return nil
}

func TwoSum2(nums []int, target int) []int {
	nummap := make(map[int]int)
	for i, num := range nums {
		ansnum := target - num
		v, ok := nummap[ansnum]
		if ok {
			return []int{v, i}
		} else {
			nummap[num] = i
		}

	}
	return nil
}
