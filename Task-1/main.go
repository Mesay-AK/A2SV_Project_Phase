package main

import (
	"fmt"
)

// PromptUser prompts the user to input subjects and grades
func PromptUser() map[string]float64 {
	var num int
	fmt.Println("Please Enter the Number of Subjects:")
	for j :=0; j<3; j++{
		_, err := fmt.Scan(&num)

		// Checking input type and validity
		if err != nil {
			fmt.Println("Please enter a number.")
			continue
		}
		if num <= 0 {
			fmt.Println("Please try again.")
			continue
		}
		break
	}

	store := make(map[string]float64)

	for i := 1; i <= num; i++ {
		var subject string
		fmt.Printf("Enter name of subject %d: ", i)
		for j:=0; j<3; j++{
			_, err := fmt.Scan(&subject)

			// Checking input type and validity
			if err != nil || subject == "" {
				fmt.Println("Please enter a subject name.")
				continue
			}
			break
		}

		var grade float64
		fmt.Printf("Enter grade for %s: ", subject)
		for {
			_, err := fmt.Scan(&grade)

			// Checking input type and validity
			if err != nil {
				fmt.Println("Please enter a valid grade.")
				continue
			}
			if grade < 0.00 || grade > 100.00 {
				fmt.Println("Grade must be between 0 and 100.Try again.")
				continue
			}
			break
		}

		store[subject] = grade
	}

	return store
}

// Calculating the average grade
func CalculateGrade(name string, stored map[string]float64) float64 {
	var total float64
	var count float64
	fmt.Println("Name:", name)

	for subject, grade := range stored {
		fmt.Printf("%s: %.2f\n", subject, grade)
		total += grade
		count++
	}

	if count == 0 {
		return 0
	}
	return total / count
}

func main() {
	fmt.Println("WELCOME TO YOUR GRADE CALCULATOR!")

	var name string
	fmt.Println("Please Enter Your Name:")
	fmt.Scan(&name)

	grades := PromptUser()
	average := CalculateGrade(name, grades)

	fmt.Printf("Your Average Grade is %.2f.", average)
}
