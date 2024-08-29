package main

import "fmt"

// Student struct to hold student information
type Student struct {
	subscore map[string]int
	avg      float64
	grade    rune
}

func main() {
	// Create a map to store subjects and their scores
	subjects := map[string]string{
		"Math":    "M",
		"Science": "S",
		"English": "E",
	}

	// Create a map to store student information
	students := make(map[string]Student)

	// Prompt the user for student information and scores for some students
	for i := 0; i < 1; i++ {
		var name string
		fmt.Printf("\nEnter student %d name: ", i+1)
		fmt.Scanln(&name)

		// Create a map to store scores for each subject
		subjectScores := make(map[string]int)

		// Initialize average and grade variables
		var totalScore, averageScore float64
		var grade rune

		// Prompt the user for scores in each subject
		for subject, subjectCode := range subjects {
			fmt.Printf("Enter %s score for %s: ", subject, name)
			score := getScore()

			// Store the score in the map
			subjectScores[subjectCode] = score

			// Accumulate total score
			totalScore += float64(score)
		}

		// Calculate average score
		averageScore = totalScore / float64(len(subjects))

		// Determine the letter grade based on the average score using control flow
		switch {
		case averageScore >= 90:
			grade = 'A'
		case averageScore >= 80:
			grade = 'B'
		case averageScore >= 70:
			grade = 'C'
		case averageScore >= 60:
			grade = 'D'
		default:
			grade = 'F'
		}

		// Create a Student struct and store it in the map
		students[name] = Student{
			subscore: subjectScores,
			avg:      averageScore,
			grade:    grade,
		}
	}

	// Display the student information, subject scores, average, and grade
	for name, student := range students {
		fmt.Printf("\nStudent: %s\n", name)
		for subject, score := range student.subscore {
			fmt.Printf("%s Score: %d\n", subject, score)
		}
		fmt.Printf("Average Score: %.2f\n", student.avg)
		fmt.Printf("Grade: %c\n", student.grade)
	}
}

// getScore prompts the user for a score and returns the entered value
func getScore() int {
	var score int
	fmt.Scanln(&score)
	return score
}
