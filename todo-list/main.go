package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

type Task struct {
    Title string
    Done  bool
}

var tasks []Task

func main() {
    reader := bufio.NewReader(os.Stdin)
    for {
        fmt.Println("1. Add Task")
        fmt.Println("2. List Tasks")
        fmt.Println("3. Delete Task")
        fmt.Println("4. Mark Task as Done")
        fmt.Println("5. Exit")
        fmt.Print("Enter choice: ")
        choiceStr, _ := reader.ReadString('\n')
        choiceStr = strings.TrimSpace(choiceStr)
        choice, err := strconv.Atoi(choiceStr)
        if err != nil {
            fmt.Println("Invalid choice")
            continue
        }

        switch choice {
        case 1:
            addTask(reader)
        case 2:
            if len(tasks) == 0 {
                fmt.Println("No tasks available.")
            } else {
                listTasks()
            }
        case 3:
            if len(tasks) == 0 {
                fmt.Println("No tasks available to delete.")
            } else {
                deleteTask(reader)
            }
        case 4:
            if len(tasks) == 0 {
                fmt.Println("No tasks available to mark as done.")
            } else {
                markTaskDone(reader)
            }
        case 5:
            return
        default:
            fmt.Println("Invalid choice")
        }
    }
}

func addTask(reader *bufio.Reader) {
    fmt.Print("Enter task title: ")
    title, _ := reader.ReadString('\n')
    title = strings.TrimSpace(title)
    tasks = append(tasks, Task{Title: title, Done: false})
    fmt.Println("Task added successfully.")
}

func listTasks() {
    for i, task := range tasks {
        status := "Pending"
        if task.Done {
            status = "Done"
        }
        fmt.Printf("%d. %s [%s]\n", i+1, task.Title, status)
    }
}

func deleteTask(reader *bufio.Reader) {
    fmt.Print("Enter task number to delete: ")
    indexStr, _ := reader.ReadString('\n')
    indexStr = strings.TrimSpace(indexStr)
    index, err := strconv.Atoi(indexStr)
    if err != nil || index < 1 || index > len(tasks) {
        fmt.Println("Invalid task number")
        return
    }
    tasks = append(tasks[:index-1], tasks[index:]...)
    fmt.Println("Task deleted successfully.")
}

func markTaskDone(reader *bufio.Reader) {
    fmt.Print("Enter task number to mark as done: ")
    indexStr, _ := reader.ReadString('\n')
    indexStr = strings.TrimSpace(indexStr)
    index, err := strconv.Atoi(indexStr)
    if err != nil || index < 1 || index > len(tasks) {
        fmt.Println("Invalid task number")
        return
    }
    tasks[index-1].Done = true
    fmt.Println("Task marked as done successfully.")
}
