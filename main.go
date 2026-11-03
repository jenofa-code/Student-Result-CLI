package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"studentresult/result"
	"studentresult/student"
)

const numSubjects = 5

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("==================================")
	fmt.Println("   STUDENT RESULT PROCESSING")
	fmt.Println("==================================")

	// Read student name and roll number
	name := readLine(reader, "Enter Student Name: ")
	rollNo := readLine(reader, "Enter Roll Number: ")

	// Read and validate marks for each subject
	marks := make([]float64, 0, numSubjects)
	for i := 1; i <= numSubjects; i++ {
		for {
			input := readLine(reader, fmt.Sprintf("Enter mark for Subject %d: ", i))
			mark, err := strconv.ParseFloat(input, 64)
			if err != nil || mark < 0 || mark > 100 {
				fmt.Println("Invalid mark! Enter a number between 0 and 100.")
				continue
			}
			marks = append(marks, mark)
			break
		}
	}

	s := student.Student{Name: name, RollNo: rollNo, Marks: marks}

	total := result.CalculateTotal(s.Marks)
	average := result.CalculateAverage(s.Marks)
	grade := result.CalculateGrade(average)

	fmt.Println()
	fmt.Println("==================================")
	fmt.Println("        STUDENT RESULT")
	fmt.Println("==================================")
	fmt.Printf("Student Name : %s\n", s.Name)
	fmt.Printf("Roll Number  : %s\n", s.RollNo)
	fmt.Printf("Total Marks  : %.2f / %d\n", total, numSubjects*100)
	fmt.Printf("Average      : %.2f\n", average)
	fmt.Printf("Grade        : %s\n", grade)
	fmt.Println("==================================")
}

// readLine prints a prompt and returns the trimmed line typed by the user.
func readLine(reader *bufio.Reader, prompt string) string {
	fmt.Print(prompt)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}
