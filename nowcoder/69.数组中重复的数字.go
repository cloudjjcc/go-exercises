package main

import "fmt"

//题目描述
//在一个长度为n的数组里的所有数字都在0到n-1的范围内。
//数组中某些数字是重复的，但不知道有几个数字是重复的。
//也不知道每个数字重复几次。请找出数组中任意一个重复的数字。
//例如，如果输入长度为7的数组{2,3,1,0,2,5,3}，那么对应的输出是第一个重复的数字2。

func main() {
	fmt.Println(duplicate([]int{2, 3, 1, 0, 2, 5, 3}))
}
func duplicate(arr []int) int {
	if len(arr) < 2 {
		return 0
	}
	for i := 0; i < len(arr); i++ {
		for arr[i] != i {
			if arr[i] == arr[arr[i]] {
				return arr[i]
			}
			arr[arr[i]], arr[i] = arr[i], arr[arr[i]]
		}
	}
	return 0
}
