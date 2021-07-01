package main

import "fmt"

func binarySearch(arr []int, hkey int) int {
	low, high := 0, len(arr)-1
	for low <= high {
		mid := low + (high-low)/2
		if(arr[mid] == hkey) {
			return mid
		} else if(arr[mid] < hkey) {
			low = mid+1
		} else if(arr[mid] > hkey) {
			high = mid-1
		}
	}
	return -1
}

func main()  {
	arr := []int{1,2,3,4,5,6,7,8,9}
	v := binarySearch(arr, 5)
	fmt.Println("The hkey is: ", v)
}