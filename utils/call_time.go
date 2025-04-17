package utils

import (
	"fmt"
	"time"
)

func lalala() {
	fmt.Println("lalala")
}

func GetTime() int {
	var now time.Time = time.Now()
	fmt.Println("now: " + now.Format(time.RFC822))
	var year int = now.Year()
	// fmt.Println(year)
	const longForm = "Jan 2, 2006 at 3:04pm (MST)"
	t, _ := time.Parse(longForm, "Feb 3, 2013 at 7:54pm (PST)")
	fmt.Println(t)

	return year
}