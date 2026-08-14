package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func showForm(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/index.html")
}

func calculate(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	name := r.FormValue("name")
	matric := r.FormValue("matric")
	level := r.FormValue("level")
	semester := r.FormValue("semester")

	totalQualityPoints := 0
	totalUnits := 0
	results := ""

	for i, subject := range subjects {
		ca, _ := strconv.Atoi(r.FormValue(fmt.Sprintf("ca%d", i)))
		exam, _ := strconv.Atoi(r.FormValue(fmt.Sprintf("exam%d", i)))
		total := ca + exam
		grade, gradePoint := getGradePoint(total)
		qualityPoints := gradePoint * units[i]
		totalQualityPoints += qualityPoints
		totalUnits += units[i]

		results += fmt.Sprintf(`
			<tr>
				<td>%s</td>
				<td>%d</td>
				<td>%d</td>
				<td>%d</td>
				<td>%d</td>
				<td><strong>%s</strong></td>
				<td>%d</td>
				<td>%d</td>
			</tr>
		`, subject, units[i], ca, exam, total, grade, gradePoint, qualityPoints)
	}

	gpa := float64(totalQualityPoints) / float64(totalUnits)
	overallClass := getOverallClass(gpa)

	// CGPA section
	cgpaSection := ""

	if semester == "First Semester" {
		// Save semester 1 results to file
		record := StudentRecord{
			Name:   name,
			Matric: matric,
			Level:  level,
			Semester1: SemesterResult{
				QualityPoints: totalQualityPoints,
				Units:         totalUnits,
				GPA:           gpa,
			},
		}
		saveRecord(record)
		cgpaSection = `
			<div class="cgpa-box">
				<p>✅ Semester 1 results saved!</p>
				<p>Come back after Semester 2 to calculate your CGPA.</p>
			</div>`

	} else if semester == "Second Semester" {
		// Try to load semester 1 results
		if recordExists(name, matric) {
			record, err := loadRecord(name, matric)
			if err == nil {
				// Calculate CGPA
				totalQP := record.Semester1.QualityPoints + totalQualityPoints
				totalU := record.Semester1.Units + totalUnits
				cgpa := float64(totalQP) / float64(totalU)
				cgpaClass := getOverallClass(cgpa)

				cgpaSection = fmt.Sprintf(`
					<div class="cgpa-box">
						<h2>📊 CGPA: %.2f / 5.0</h2>
						<p><strong>Semester 1 GPA:</strong> %.2f</p>
						<p><strong>Semester 2 GPA:</strong> %.2f</p>
						<p><strong>Degree Class:</strong> %s</p>
					</div>`, cgpa,
					record.Semester1.GPA,
					gpa,
					cgpaClass)
			}
		} else {
			cgpaSection = `
				<div class="cgpa-box">
					<p>⚠️ No Semester 1 record found!</p>
					<p>Please calculate Semester 1 first.</p>
				</div>`
		}
	}

	html := fmt.Sprintf(`
<!DOCTYPE html>
<html>
<head>
	<title>UniGrade Results</title>
	<link rel="stylesheet" href="/static/style.css">
	<style>
		@media print {
			.btn, .btn-print { display: none; }
			body { background: white; }
			.gpa-box { background: #333 !important; -webkit-print-color-adjust: exact; }
			.cgpa-box { background: #1a237e !important; -webkit-print-color-adjust: exact; }
		}
		.btn-print {
			background: #2196F3;
			color: white;
			padding: 14px 30px;
			border: none;
			border-radius: 8px;
			cursor: pointer;
			font-size: 16px;
			width: 100%%;
			margin-top: 10px;
		}
		.btn-print:hover { background: #1976D2; }
		.cgpa-box {
			background: #1a237e;
			color: white;
			padding: 25px;
			border-radius: 8px;
			text-align: center;
			margin: 20px 0;
		}
		.cgpa-box h2 { font-size: 32px; margin-bottom: 10px; }
		.cgpa-box p { font-size: 16px; margin: 5px 0; color: #90caf9; }
	</style>
</head>
<body>
	<div class="container">
		<h1>🎓 UniGrade Results</h1>

		<div class="student-info">
			<p><strong>Name:</strong> %s</p>
			<p><strong>Matric Number:</strong> %s</p>
			<p><strong>Level:</strong> %s</p>
			<p><strong>Semester:</strong> %s</p>
		</div>

		<table>
			<tr>
				<th>Course</th>
				<th>Units</th>
				<th>CA (40)</th>
				<th>Exam (60)</th>
				<th>Total</th>
				<th>Grade</th>
				<th>Grade Point</th>
				<th>Quality Points</th>
			</tr>
			%s
			<tr class="total-row">
				<td colspan="7"><strong>Total Quality Points</strong></td>
				<td><strong>%d</strong></td>
			</tr>
		</table>

		<div class="gpa-box">
			<h2>GPA: %.2f / 5.0</h2>
			<p>%s</p>
		</div>

		%s

		<a href="/" class="btn">← Calculate Again</a>
		<button onclick="window.print()" class="btn-print">🖨️ Print / Save as PDF</button>
	</div>
</body>
</html>`, name, matric, level, semester, results, totalQualityPoints, gpa, overallClass, cgpaSection)

	fmt.Fprint(w, html)
}
