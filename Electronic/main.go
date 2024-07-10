package main

import (
	electronic "Electronic/electro"
	"fmt"
)

func printCharacteristics(p electronic.Phone) {
	fmt.Println("Brand: ", p.Brand())
	fmt.Println("Model: ", p.Model())
	fmt.Println("Type: ", p.Type())
	if s, ok := p.(electronic.StationPhone); ok {
		fmt.Println("Buttons Count ", s.ButtonsCount())
	}
	if s, ok := p.(electronic.Smartphone); ok {
		fmt.Println("OS: ", s.OS())
	}
	fmt.Println()
}

func main() {
	applePhone := electronic.NewApplePhone("iPhone X", "iOS")
	androidPhone := electronic.NewAndroidPhone("Samsung", "Galaxy S10", "Android")
	radioPhone := electronic.NewRadioPhone("Sony", "XYZ", 12)

	printCharacteristics(applePhone)
	printCharacteristics(androidPhone)
	printCharacteristics(radioPhone)
}
