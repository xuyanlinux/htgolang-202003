package main

import (
	"fmt"
)

/*
2. 切片，获取切片中第二个最大元素(相同大小的元素) 1 2 4 5 5 并排第一 4 不并排 5
第2条说明：
1 2 4 5 5
要计算最大的第二个数字，有两种
1. 5 5 相同 算并列第一，所以4就是第二个最大的数字
2. 5 5 相同，但不考虑并列情况，5就是第二个最大的数字
这两种功能都实现一下
*/

var slice = []int{2, 1, 5, 4, 5}
var newSlice  []int
var myMap = make(map[int]int)

//当不去重时的情况,冒泡算法排序，然后取len(slice)-1索引位置的值。同时得考虑最大值可能并列
func bubbleAlgorithmV1(slice []int) int {
	for j := 0; j < len(slice); j++ {
		for k := 0; k < len(slice)-1-j; k++ {
			if slice[k] > slice[k+1] {
				tmp := slice[k]
				slice[k] = slice[k+1]
				slice[k+1] = tmp
			}
		}
	}
	return slice[len(slice)-1]
}

//当最大值算并列（存在并列）时，对sliec切片去重并排序后再取第二个最大值
func secMaxNum(slice []int) int {
	//去重
	for v, k := range slice {
		myMap[k] = v
	}
	for k := range myMap {
		newSlice=append(newSlice,k)
	}
	//排序
	bubbleAlgorithmV1(newSlice)
	return newSlice[len(newSlice)-2]
}

func main() {
	//当最大值不算并列（存在并列）时，那么第二个最大值依然是最大值
	fmt.Printf("算并列后的最大值是：%d\n",bubbleAlgorithmV1(slice))
	//当最大值算并列（存在并列）时，那么第二个最大值就是仅小于最大值的那个值
	fmt.Printf("不算并列后的最大值是：%d\n",secMaxNum(slice))
}
