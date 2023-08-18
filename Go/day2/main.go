package main

import "fmt"

func main() {

	// congratulation := "happy birthday"
    // fmt.Println(congratulation)

	// we can decleare two variable
	// averageOpenrate, displayMessage := .23, "this is average message rate";
    // fmt.Println(averageOpenrate, displayMessage)

	// accountAge := 1.2
	// accountAgeInt := int(accountAge)
    // fmt.Println("your account has existed for", accountAgeInt , "years");


	// const secondsInMinute = 60
	// const minutesInHour = 60
	// const secondsInHour = secondsInMinute * minutesInHour
    // fmt.Println("no of second in hour is", secondsInHour)



	const name = "sidharth shukla"
	const openRate = 30.5
    msg := fmt.Sprintf(
		"Hii %s, your open rate is %.1f percent",
		name,
		openRate,
	)
	fmt.Println(msg)
}
