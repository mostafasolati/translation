package persian

const ()

func TranslateWeekDay(day int) string {
	switch day {
	case 0:
		return "شنبه"
	case 1:
		return "یکشنبه"
	case 2:
		return "دوشنبه"
	case 3:
		return "سه شنبه"
	case 4:
		return "چهارشنبه"
	case 5:
		return "پنجشنبه"
	case 6:
		return "جمعه"
	default:
		return "نامشخص"
	}
}
