//6. go - хэш-таблица (map)
package main

import "fmt"

func main() {
	ages := map[string]int{
		"Alice": 28,
		"Bob":   32,
		"Carol": 25,
	}

	fmt.Println(ages)
	fmt.Println(ages["Bob"]) // 32

	ages["Dave"] = 30

	value, ok := ages["Alice"] // проверка наличия
	if ok {
		fmt.Println(value) // 28
	}

	fmt.Println(ages)
}