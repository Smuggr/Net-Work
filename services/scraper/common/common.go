package scraper

import (
	"fmt"
)

type Website int
type DayOfWeek int

type WebsiteData struct {
	Website Website
	Data interface{}
}


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
