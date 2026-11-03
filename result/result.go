package result

// CalculateTotal adds all the marks together.
func CalculateTotal(marks []float64) float64 {
	total := 0.0
	for _, m := range marks {
		total += m
	}
	return total
}

// CalculateAverage returns the average of the marks (0 for an empty slice).
func CalculateAverage(marks []float64) float64 {
	if len(marks) == 0 {
		return 0
	}
	return CalculateTotal(marks) / float64(len(marks))
}

// CalculateGrade assigns a grade based on the average marks.
func CalculateGrade(average float64) string {
	switch {
	case average >= 90:
		return "O"
	case average >= 80:
		return "A+"
	case average >= 70:
		return "A"
	case average >= 60:
		return "B+"
	case average >= 50:
		return "B"
	default:
		return "RA"
	}
}
