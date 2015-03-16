package main

import (
	"fmt"
)

func main() {
	var month int = 6
	var season string
	switch month {
	case 1, 2, 3:
		season = "spring"
	case 4, 5, 6:
		season = "summer"
	case 7, 8, 9:
		season = "fall"
	case 10, 11, 12:
		season = "winter"
	}
	fmt.Println(season)
}
