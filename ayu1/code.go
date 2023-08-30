package ayu1

import (
	"fmt"
)

func Playground() {

	var age = 70
	fmt.Println("adulto:", age >= 18 && age <= 60)
	fmt.Println("niño o viejo:", age < 18 || age > 60)

}
