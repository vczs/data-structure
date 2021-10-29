package main

import "fmt"

//原始方法：遍历切片

//使用稀疏数组
type ValNode struct {
	row int //行
	col int //列
	val int //值
}

func main() {
	var slice [11][11]int
	slice[1][2] = 1
	slice[2][3] = 2
	//存储节点的切片
	var arr []ValNode
	//创建节点并添加到切片
	valNode := ValNode{ //默认值节点
		row: 11,
		col: 11,
		val: 0,
	}
	arr = append(arr, valNode) //将默认值节点valNode添加到arr
	for k1, v1 := range slice {
		for k2, v2 := range v1 {
			if v2 != 0 {
				valNode := ValNode{ //不同值元素节点
					row: k1,
					col: k2,
					val: v2,
				}
				arr = append(arr, valNode) //将不同值元素节点valNode添加到arr
			}
		}
	}
	fmt.Printf("转成稀疏数组arr:%v", arr)

	//恢复原始数组
	var recoveryArr [11][11]int
	for k, v := range arr {
		if k != 0 {
			recoveryArr[v.row][v.col] = v.val
		}
	}
	fmt.Println("恢复后的数组:")
	for _, v1 := range recoveryArr {
		for _, v2 := range v1 {
			fmt.Printf("%v ", v2)
		}
		fmt.Println()
	}
}
