// 2. Go - массив(slice)
package main

import "fmt"

func main() {
	nums := []int{10, 20, 30, 40}

	fmt.Println(nums)    // [10 20 30 40]
	fmt.Println(nums[1]) // 20

	nums = append(nums, 50)

	nums = append(nums, 0)
	copy(nums[2:], nums[1:])
	nums[1] = 15

	fmt.Println(nums)      // [10 15 20 30 40 50]
	fmt.Println(len(nums)) // 6
	fmt.Println(cap(nums)) // вместимость может отличаться
}