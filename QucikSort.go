package main

import "fmt"

type quick_sort struct {
	num_arr [9]int
}

func (q *quick_sort) QuickSort(low int, high int) {
	var pivot int
	if(low<high) {
		pivot = q.Partition(low, high)//将q *quick_sort一分为二 算出枢轴值pivot
		q.QuickSort(low, pivot-1)//对低子表递归排序
		q.QuickSort(pivot+1, high)//对高子表递归排序
	}
}

//交换子表(子数组)的记录，使枢轴记录到位，并返回其所在的位置
func (q *quick_sort) Partition(low int, high int) int {
	var pivotkey int
	var temp int//用于交换
	pivotkey = q.num_arr[low]//用子表的第一个记录做枢轴记录
	for {
		if (low>=high) {
			break
		}
		//从右边向中间扫描
		//高位值大于枢轴,继续循环,执行high--
		for {
			if (low >= high || q.num_arr[high] <= pivotkey) {
				break
			}
			high--
		}
		//高位数字小于枢轴,交换到低位
		temp = q.num_arr[low]
		q.num_arr[low] = q.num_arr[high]
		q.num_arr[high] = temp

		//从左边向中间扫描
		//低位小于枢轴,继续循环,执行low--
		for {
			if (low >= high || q.num_arr[low] >= pivotkey) {
				break
			}
			low++
		}
		//低位大于枢轴，交换到高位
		temp = q.num_arr[high]
		q.num_arr[high] = q.num_arr[low]
		q.num_arr[low] = temp

	}
	return low//返回枢轴的位置
}

func main()  {
	var q = quick_sort{}
	q.num_arr = [9]int{50,10,90,30,70,40,80,60,20}
	q.QuickSort(0, len(q.num_arr)-1)
	fmt.Printf("%v", q.num_arr)
}