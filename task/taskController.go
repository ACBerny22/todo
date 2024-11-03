package task

import (
	"encoding/csv"
	"fmt"
	"os"
)

type Task struct {
	Name string
	Description string
	Completed bool
}

func CreateTask(task Task) {
	fmt.Println("Name:", task.Name)
	fmt.Println("Description:", task.Description)
	fmt.Printf("Completed: %t \n", task.Completed)

	// Open the file with write permissions, creating it if it doesn't exist
	file, err := os.OpenFile("bin/database.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Convert task data to a slice of strings
	data := []string{task.Name, task.Description, fmt.Sprintf("%t", task.Completed)}
	
	// Write the record to the CSV file
	if err := writer.Write(data); err != nil {
		fmt.Println("Error writing to CSV:", err)
	}
}