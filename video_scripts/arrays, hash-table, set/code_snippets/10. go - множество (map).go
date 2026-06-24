// 10. go - множество (map)
package main

import "fmt"

func main() {
	visited := make(map[string]struct{})

	visited["Alice"] = struct{}{}
	visited["Bob"] = struct{}{}
	visited["Alice"] = struct{}{}

	_, ok := visited["Alice"]
	fmt.Println(ok)           // true

	_, ok = visited["Carol"]
	fmt.Println(ok)           // false
}