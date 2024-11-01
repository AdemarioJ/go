package main

import "fmt"

func weekDay(number int) string {
	switch number {
	case 1:
		return "Monday"
	case 2:
		return "Tuesday"
	case 3:
		return "Wednesday"
	case 4:
		return "Thursday"
	case 5:
		return "Friday"
	case 6:
		return "Saturday"
	case 7:
		return "Sunday"
	default:
		return "Invalid day"
	}
}

func weekDay2(number int) string {
	var weekDay string

	switch {
	case number == 1:
		weekDay = "Monday"
	case number == 2:
		weekDay = "Tuesday"
	case number == 3:
		weekDay = "Wednesday"
	case number == 4:
		weekDay = "Thursday"
	case number == 5:
		weekDay = "Friday"
	case number == 6:
		weekDay = "Saturday"
	case number == 7:
		weekDay = "Sunday"
	default:
		weekDay = "Invalid day"
	}
	return weekDay
}

func main() {
	fmt.Println("Switch")

	fmt.Println(weekDay(1))

	fmt.Println(weekDay2(7))
}
