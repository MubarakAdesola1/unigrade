package main

// getGradePoint takes a score and returns grade and grade point
func getGradePoint(score int) (string, int) {
	if score >= 70 {
		return "A", 5
	} else if score >= 60 {
		return "B", 4
	} else if score >= 50 {
		return "C", 3
	} else if score >= 45 {
		return "D", 2
	} else if score >= 40 {
		return "E", 1
	}
	return "F", 0
}

// getOverallClass returns degree class based on GPA
func getOverallClass(gpa float64) string {
	if gpa >= 4.5 {
		return "First Class"
	} else if gpa >= 3.5 {
		return "Second Class Upper"
	} else if gpa >= 2.5 {
		return "Second Class Lower"
	} else if gpa >= 1.5 {
		return "Third Class"
	}
	return "Fail"
}
