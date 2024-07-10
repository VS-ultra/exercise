package main

import "fmt"

func hashstr(val string) uint64 {
	var strUint uint64 = 1
	hashFactor := uint64(11)
	r := []rune(val)
	for _, runeValue := range r {
		strUint *= hashFactor
		strUint += uint64(runeValue)
	}
	hashStr := strUint % 1000
	return hashStr
}

func main() {
	var str string
	fmt.Println("Enter num :")
	fmt.Scan(&str)
	if str == "" {
		return
	}
	if str != "" {
		fmt.Println("Get your hash :", hashstr(str))
	}

}
