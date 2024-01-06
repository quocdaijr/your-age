package your_age

import "time"

func CalculateAgeFromBirthDay(birthDay string) int {
	birthDayDatetime, _ := time.Parse("2006-01-02", date)
	return time.Now().Year() - birthDayDatetime.Year()
}

func PrintYourAge(birthDay string) {
	println(CalculateAgeFromBirthDay(birthDay))
}
