package main

import (
	"fmt"
)

func hashint64(val int64) uint64 {
	hash_uint := val % 1000
	return uint64(hash_uint)
}
func main() {
	var num int64
	fmt.Println("Enter num :")
	fmt.Scan(&num)
	fmt.Println("Get your hash :", hashint64(num))
}
