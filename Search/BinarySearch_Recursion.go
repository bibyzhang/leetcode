package main

import "fmt"

func binary_search(arr []int, low, high, hkey int) int {
	if(low>high) {
		return -1
	}
	mid := low + (high-low)/2
	if(arr[mid]>hkey) {
		return binary_search(arr, low, mid-1, hkey)
	} else if(arr[mid]<hkey) {
		return binary_search(arr, mid+1, high, hkey)
	}

	return mid
}

func main()  {
	arr := []int{1,2,3,4,5,6,7,8,9}
	v := binary_search(arr, 0, 8, 5)
	fmt.Println("The hkey is: ", v)
}