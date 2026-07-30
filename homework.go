package homework

//给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。
// 找出那个只出现了一次的元素。可以使用 for 循环遍历数组，结合 if 条件判断和 map 数据结构来解决，
// 例如通过 map 记录每个元素出现的次数，然后再遍历 map 找到出现次数为1的元素。
func SingleNumber(data []int) int {

	res := 0

	for _, v := range data {

		res ^= v //res = res ^ v

	}

	return res
}

//判断一个整数是否是回文数
func IsPalindrome(x int) bool {
	//负数不是回文 直接返回false
	if x < 0 {
		return false
	}

	origin := x  //保存原始数字，后面用来对比  （正序）
	reverse := 0 //用来存放反转的数字  （到序）

	//获取倒序
	for x > 0 {
		last := x % 10 //获取最后一位 （ % 10  取模运算）
		reverse = reverse*10 + last
		x = x / 10 //去掉最后一位
	}

	return origin == reverse

}

//给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效
func IsValidParentheses(s string) bool {

	//建立括号映射 key：右括号 value：对应左括号
	match := map[byte]byte{

		')': '(',
		']': '[',
		'}': '{',
	}

	//创建一个切片。初始化长度0  用切片模拟栈 后进先出
	stack := make([]byte, 0)

	for i := 0; i < len(s); i++ {
		char := s[i]
		if left, isRight := match[char]; isRight {

			//栈空 或者 栈顶和预期左括号不一致 left是左括号 ( ] }
			if len(stack) == 0 || stack[len(stack)-1] != left {

				return false
			}

			//匹配成功就弹出栈顶
			stack = stack[:len(stack)-1]

		} else {
			//左括号 入栈
			stack = append(stack, char)
		}

	}

	//栈为空 == 全部匹配成功
	return len(stack) == 0

}

//查找字符串数组中的最长公共前缀
func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	for j := 0; j < len(strs[0]); j++ {
		//取第一个字符 的对应下标字符
		baseChar := strs[0][j]

		//和剩余字符串j 位置字符对比
		for i := 1; i < len(strs); i++ {

			//当前字符串长度不够  或者  字符不相等
			if j >= len(strs[i]) || strs[i][j] != baseChar {
				return strs[0][:j]
			}

		}

	}

	//全部匹配
	return strs[0]

}

//给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一
func PlusOne(digits []int) []int {
	carry := 1
	//从最后一位向前遍历
	for i := len(digits) - 1; i >= 0; i-- {
		sum := digits[i] + carry

		if sum == 10 {
			digits[i] = 0
			carry = 1
		} else {
			digits[i] = sum
			carry = 0
			break //进位消失 不用继续计算
		}

	}

	if carry == 1 {
		//进位
		return append([]int{1}, digits...)
	}

	//不进位
	return digits

}

//两数之和 在数组中找到两个不同下标的数字，二者相加 = target；返回这两个下标。
func TwoSum(nums []int, target int) []int {

	for i := 0; i < len(nums)-1; i++ {

		for j := i + 1; j < len(nums); j++ {

			if nums[i]+nums[j] == target {

				return []int{i, j}
			}

		}

	}

	return nil

}

func TwoSum2(nums []int, target int) []int {

	// [2,7,4,6] 9
	//key 数组字，value 对应下标
	hashMap := make(map[int]int)

	for idx, num := range nums {
		need := target - num

		if pos, exists := hashMap[need]; exists {

			return []int{pos, idx}
		}

		//不存在 就把当前数字存入map
		hashMap[num] = idx

	}

	return nil

}
