package common

import (
	"fmt"
)

type Website int
type DayOfWeek int


const (
	UonetOptivum Website = iota
)

const (
	Monday DayOfWeek = iota
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)


func Initialize() {
	fmt.Println("Initializing scraper")
}
