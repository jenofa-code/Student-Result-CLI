# Student Result Processing CLI Application

## Project Description

The Student Result Processing CLI Application is developed using Go (Golang). It accepts student details and subject marks, calculates total marks and average, and assigns a grade based on the average.

## Objectives

* Develop a modular application using Go.
* Store student information using structures.
* Calculate total marks and average.
* Assign grades using conditional statements.
* Handle invalid marks and user input.
* Use Git and GitHub for version control.

## Technologies Used

* Go (Golang)
* Visual Studio Code
* Git
* GitHub

## Project Structure

```text
StudentResultCLI/
├── main.go
├── go.mod
├── README.md
├── .gitignore
├── student/
│   └── student.go
└── result/
    └── result.go
```

## Features

* Accepts student name and roll number.
* Accepts marks for five subjects.
* Validates marks between 0 and 100.
* Calculates total and average marks.
* Assigns grades based on average marks.
* Displays the result in the terminal.

## Grade System

| Average   | Grade |
| --------- | ----- |
| 90–100    | O     |
| 80–89.99  | A+    |
| 70–79.99  | A     |
| 60–69.99  | B+    |
| 50–59.99  | B     |
| Below 50  | RA    |

## How to Run

1. Install Go.
2. Clone this repository.
3. Open the project directory in the terminal.
4. Run:

```
go run .
```

## Validation

The project can be checked using:

```
go build .
go vet ./...
```

## Version Control

Git is used to track changes, and GitHub is used to host the source code.

## Conclusion

This project demonstrates modular programming, functions, structures, slices, conditional statements, input validation, and version control using Go.
