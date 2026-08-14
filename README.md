# 🎓 UniGrade — University GPA Calculator

A web-based GPA and CGPA calculator built with Go, designed for Nigerian university students.

## 📋 Features

- ✅ Calculate GPA for 6 courses per semester
- ✅ Supports CA (40 marks) + Exam (60 marks) grading system
- ✅ Automatic grade calculation (A, B, C, D, E, F)
- ✅ Quality points and GPA calculation out of 5.0
- ✅ Degree class classification
- ✅ Saves Semester 1 results automatically
- ✅ Calculates CGPA after Semester 2
- ✅ Print / Save results as PDF
- ✅ Clean and responsive web interface

## 🎯 Grading System

| Score | Grade | Grade Point |
|-------|-------|-------------|
| 70 - 100 | A | 5 |
| 60 - 69 | B | 4 |
| 50 - 59 | C | 3 |
| 45 - 49 | D | 2 |
| 40 - 44 | E | 1 |
| 0 - 39 | F | 0 |

## 🏆 Degree Classification

| CGPA | Class |
|------|-------|
| 4.5 - 5.0 | First Class 🏆 |
| 3.5 - 4.4 | Second Class Upper 🥇 |
| 2.5 - 3.4 | Second Class Lower 🥈 |
| 1.5 - 2.4 | Third Class 🥉 |
| 0 - 1.4 | Fail ❌ |

## 📚 Courses & Credit Units

| Course | Credit Units |
|--------|-------------|
| Mathematics | 3 |
| English | 3 |
| Physics | 2 |
| Chemistry | 2 |
| Biology | 2 |
| Computer Science | 1 |
| **Total** | **13** |

## 🚀 How to Run

### Prerequisites
- Go 1.21 or higher installed

### Steps

1. Clone the repository:
```bash
git clone https://github.com/MubarakAdesola1/unigarde.git
cd unigarde
```

2. Initialize the Go module:
```bash
go mod init unigarde
```

3. Run the server:
```bash
go run .
```

4. Open your browser and go to:

## 📁 Project Structure

unigarde/
├── main.go → Server setup and routes
├── handlers.go → Form display and GPA calculation
├── grades.go → Grade and degree class logic
├── storage.go → Save and load student records
├── static/
│ ├── index.html → Student input form
│ └── style.css → Styling
├── data/ → Stores student JSON files
├── README.md → Project documentation
└── go.mod → Go module file

## 📖 How to Use

### Calculate Semester 1 GPA
1. Open `http://localhost:8080`
2. Fill in your student information
3. Select **"First Semester"**
4. Enter CA and Exam scores for each course
5. Click **"Calculate GPA"**
6. Your results are automatically saved!

### Calculate Semester 2 GPA & CGPA
1. Open `http://localhost:8080`
2. Fill in the **same** name and matric number as Semester 1
3. Select **"Second Semester"**
4. Enter CA and Exam scores for each course
5. Click **"Calculate GPA"**
6. Your CGPA will be calculated automatically!

### Print / Save as PDF
1. After calculating your GPA
2. Click **"🖨️ Print / Save as PDF"**
3. Select **"Save as PDF"** in the print dialog
4. Click **Save**

## ⚠️ Important Notes

- Name and matric number must match **exactly** between Semester 1 and Semester 2
- CA marks cannot exceed **40**
- Exam marks cannot exceed **60**
- Total marks cannot exceed **100**

## 🛠️ Built With

- **Go** — Backend server and logic
- **HTML/CSS** — Frontend interface
- **JSON** — Student record storage
- **JavaScript** — Print functionality

## 👨‍💻 Author

**MubarakAdesola1**
github.com/MubarakAdesola1

- Learn2Earn Nigeria Fellowship
- Federal University of Technology Minna

## 📄 License

This project is open source and available under the [MIT License](LICENSE).