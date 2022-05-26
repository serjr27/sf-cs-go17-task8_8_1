package main

import (
	"fmt"

	"phones/electronic"
)

func printCharacteristics(phone interface{}) {
	switch phone := phone.(type) {
	case electronic.Phone:
		fmt.Println("Brand:\t", phone.Brand())
		fmt.Println("Model:\t", phone.Model())
		fmt.Println("Type:\t", phone.Type())

		switch phone := phone.(type) {
		case electronic.Smartphone:
			fmt.Println("OS:\t", phone.OS())
		case electronic.StationPhone:
			fmt.Println("Buttons:", phone.ButtonCounts())
		}

	default:
		fmt.Println("Неизвестный телефон")
	}

	fmt.Println()
}

func main() {
	apple := electronic.NewApplePhone("iPhone13", "iOS 15")
	android := electronic.NewAndroidPhone("Huawai", "P30 Lite", "Android 10")
	radio := electronic.NewRadioPhone("Panasonic", "KX-TG2512", 17)

	printCharacteristics(apple)
	printCharacteristics(android)
	printCharacteristics(radio)
}
