package main

import "fmt"

type insert_sort struct {
	num_arr [16]int
}

func (ins *insert_sort) InserSort() {
	for i:=1; i<=len(ins.num_arr)-1; i++  {
		//对比,找到插入位置
		//假设a[j]是有序的,初始只有(a[j],j=0)是有序的
		//因为已经有序,找到对比小的位置就是要插入的位置
		insert_key := i//插入位置是本身,不用移动数据
		for j:=0; j<i; j++ {
			if(ins.num_arr[i]<ins.num_arr[j]) {
				insert_key = j
				break//找到了就不用再继续对比
			}
		}
		
		//插入数据
		if(i!=insert_key) {
			tmp_val := ins.num_arr[i]
			for k:=i; k>insert_key; k-- {
				//移动数据
				ins.num_arr[k] = ins.num_arr[k-1]
			}
			ins.num_arr[insert_key] = tmp_val
		}
	}
}

func main()  {
	var ins = insert_sort{}
	//ins.num_arr = [9]int{5,6,2,4,1,7,3,9,8}
	ins.num_arr = [16]int{5,6,2,4,1,7,3,9,8,55,33,66,99,22,11,77}
	ins.InserSort()
	fmt.Print("%v", ins.num_arr)
}