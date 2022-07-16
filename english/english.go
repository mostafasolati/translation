package english

func TranslateWeekDay(day int) string {
	var days = []string{"Saturday", "Sunday", "Monday", "Tuesday", "Wedensday", "Thursday", "Friday"}
	if day > 6 && day < 0 {
		return "Unknown"
	}
	return days[day]
}
