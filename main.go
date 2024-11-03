package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
	"todo/task"
)

func main() {
	fileName := "bin/database.csv"
	err := createFile(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	var action int
	fmt.Print("What do you want to do?\n")
	fmt.Print("1. Create task\n")
	fmt.Print("2. Read tasks\n")
	fmt.Scanln(&action)

	if(action == 1){
		fmt.Println("Create task")
		newTask := askForInput();
		task.CreateTask(newTask)
	} else if(action == 2){
		fmt.Println("Read tasks")
	}

}

func askForInput() task.Task {
	var name string
	var description string
	var completed bool
	inputReader := bufio.NewReader(os.Stdin)

	fmt.Print("Name: ")
	name, _ = inputReader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Print("Description: ")
	description, _ = inputReader.ReadString('\n')
	description = strings.TrimSpace(description)

	fmt.Print("Completed: ")
	fmt.Scanln(&completed)

	return task.Task{
		Name: name,
		Description: description,
		Completed: completed,
	}
}

func createFile(fileName string) error {
	_, err := os.Stat(fileName)
	// If file doesn't exist, create it
	if os.IsNotExist(err) {
		file, err := os.Create(fileName)
		if err != nil {
			fmt.Println(err)
			return err
		}
		// The defers apperently must be called after the function returns
		defer file.Close() // Defer the closing of the file

		// Set the headers.
		writer := csv.NewWriter(file)
		defer writer.Flush()

		headers := []string{"Name", "Description", "Completed"}
		writer.Write(headers)
	
	}else if(err != nil){
		return err
	}else{
		fmt.Println("Reading file...")
	}
	return nil
}