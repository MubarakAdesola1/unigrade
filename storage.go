package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// SemesterResult holds the results of one semester
type SemesterResult struct {
	QualityPoints int     `json:"qualityPoints"`
	Units         int     `json:"units"`
	GPA           float64 `json:"gpa"`
}

// StudentRecord holds the full student data saved to file
type StudentRecord struct {
	Name      string         `json:"name"`
	Matric    string         `json:"matric"`
	Level     string         `json:"level"`
	Semester1 SemesterResult `json:"semester1"`
}

// generateFileName creates a safe filename from name and matric
// e.g. "Mubarak" + "ABC/2021/001" → "ABC_2021_001_Mubarak.json"
func generateFileName(name, matric string) string {
	// Replace "/" and spaces with "_" to make a safe filename
	safematric := strings.ReplaceAll(matric, "/", "_")
	safematric = strings.ReplaceAll(safematric, " ", "_")
	safeName := strings.ReplaceAll(name, " ", "_")
	return fmt.Sprintf("data/%s_%s.json", safematric, safeName)
}

// saveRecord saves a student's semester 1 results to a JSON file
func saveRecord(record StudentRecord) error {
	// Convert the record to JSON format
	data, err := json.MarshalIndent(record, "", "    ")
	if err != nil {
		return err
	}

	// Generate the filename
	filename := generateFileName(record.Name, record.Matric)

	// Write the JSON to the file
	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}

	fmt.Println("Saved record to", filename)
	return nil
}

// loadRecord loads a student's semester 1 results from a JSON file
func loadRecord(name, matric string) (StudentRecord, error) {
	filename := generateFileName(name, matric)

	// Read the file
	data, err := os.ReadFile(filename)
	if err != nil {
		return StudentRecord{}, fmt.Errorf("no semester 1 record found for %s (%s)", name, matric)
	}

	// Convert JSON back to StudentRecord
	var record StudentRecord
	err = json.Unmarshal(data, &record)
	if err != nil {
		return StudentRecord{}, err
	}

	return record, nil
}

// recordExists checks if a semester 1 record already exists for a student
func recordExists(name, matric string) bool {
	filename := generateFileName(name, matric)
	_, err := os.Stat(filename)
	return err == nil
}
