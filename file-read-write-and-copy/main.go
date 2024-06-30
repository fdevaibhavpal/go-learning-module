package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "path/filepath"
    "strings"
)

func main() {
    reader := bufio.NewReader(os.Stdin)

    for {
        fmt.Println("Choose an option:")
        fmt.Println("1. Create File")
        fmt.Println("2. Copy File")
        fmt.Println("3. Read File Content")
        fmt.Println("4. Write to File")
        fmt.Println("5. Delete File")
        fmt.Println("6. Exit")
        fmt.Print("Enter choice: ")
        choice, _ := reader.ReadString('\n')
        choice = strings.TrimSpace(choice)

        switch choice {
        case "1":
            createFile(reader)
        case "2":
            copyFile(reader)
        case "3":
            readFileContent(reader)
        case "4":
            writeFileContent(reader)
        case "5":
            deleteFile(reader)
        case "6":
            fmt.Println("Exiting.")
            return
        default:
            fmt.Println("Invalid choice.")
        }
    }
}

func createFile(reader *bufio.Reader) {
    fmt.Print("Enter file path to create: ")
    filePath, _ := reader.ReadString('\n')
    filePath = strings.TrimSpace(filePath)


    if _, err := os.Stat(filePath); err == nil {
        fmt.Println("File already exists. Exiting.")
        return
    }

    file, err := os.Create(filePath)
    if err != nil {
        fmt.Println("Error creating file:", err)
        return
    }
    defer file.Close()

    fmt.Println("File created successfully!")


    fmt.Print("Enter content to write: ")
    content, _ := reader.ReadString('\n')
    content = strings.TrimSpace(content)

    formattedContent := formatContent(content)

    _, err = file.WriteString(formattedContent + "\n")
    if err != nil {
        fmt.Println("Error writing to file:", err)
        return
    }

    fmt.Println("Content written successfully!")
}

func copyFile(reader *bufio.Reader) {
    var sourceFile, destinationFile string
    fmt.Print("Enter source file path: ")
    sourceFile, _ = reader.ReadString('\n')
    sourceFile = strings.TrimSpace(sourceFile)
    fmt.Print("Enter destination file path: ")
    destinationFile, _ = reader.ReadString('\n')
    destinationFile = strings.TrimSpace(destinationFile)

    src, err := os.Open(sourceFile)
    if err != nil {
        fmt.Println("Error opening source file:", err)
        return
    }
    defer src.Close()


    destDir := filepath.Dir(destinationFile)
    if _, err := os.Stat(destDir); os.IsNotExist(err) {
        fmt.Print("Destination directory does not exist. Do you want to create it? (y/n): ")
        response, _ := reader.ReadString('\n')
        response = strings.TrimSpace(strings.ToLower(response))
        if response == "y" {
            err = os.MkdirAll(destDir, os.ModePerm)
            if err != nil {
                fmt.Println("Error creating destination directory:", err)
                return
            }
        } else if response == "n" {
            fmt.Println("Destination directory not created. Exiting.")
            return
        } else {
            fmt.Println("Invalid input. Exiting.")
            return
        }
    }

    dst, err := os.Create(destinationFile)
    if err != nil {
        fmt.Println("Error creating destination file:", err)
        return
    }
    defer dst.Close()

    _, err = io.Copy(dst, src)
    if err != nil {
        fmt.Println("Error copying file:", err)
        return
    }

    fmt.Println("File copied successfully!")
}

func readFileContent(reader *bufio.Reader) {
    fmt.Print("Enter file path to read: ")
    filePath, _ := reader.ReadString('\n')
    filePath = strings.TrimSpace(filePath)

    content, err := os.ReadFile(filePath)
    if err != nil {
        fmt.Println("Error reading file:", err)
        return
    }

    fmt.Println("File content:")
    fmt.Println(string(content))
}

func writeFileContent(reader *bufio.Reader) {
    fmt.Print("Enter file path to write: ")
    filePath, _ := reader.ReadString('\n')
    filePath = strings.TrimSpace(filePath)

   
    if _, err := os.Stat(filePath); os.IsNotExist(err) {
        fmt.Println("File does not exist. Exiting.")
        return
    }

    fmt.Print("Enter content to write: ")
    content, _ := reader.ReadString('\n')
    content = strings.TrimSpace(content)

    fmt.Print("Do you want to overwrite the file or append to it? (o/a): ")
    mode, _ := reader.ReadString('\n')
    mode = strings.TrimSpace(strings.ToLower(mode))

    var file *os.File
    var err error

    if mode == "o" {
        file, err = os.Create(filePath) 
    } else if mode == "a" {
        file, err = os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644) 
    } else {
        fmt.Println("Invalid choice. Exiting.")
        return
    }

    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close()


    formattedContent := formatContent(content)

    if mode == "a" {
        _, err = file.WriteString(formattedContent + "\n")
    } else {
        _, err = file.WriteString(formattedContent + "\n")
    }

    if err != nil {
        fmt.Println("Error writing to file:", err)
        return
    }

    fmt.Println("Content written successfully!")
}

func formatContent(content string) string {
    words := strings.Fields(content)
    var formattedContent string
    var line string

    for _, word := range words {
        if len(line)+len(word)+1 > 50 {
            formattedContent += line 
            line = word
        } else {
            if line == "" {
                line = word
            } else {
                line += " " + word
            }
        }
    }

    if line != "" {
        formattedContent += line
    }

    return formattedContent
}

func deleteFile(reader *bufio.Reader) {
    fmt.Print("Enter file path to delete: ")
    filePath, _ := reader.ReadString('\n')
    filePath = strings.TrimSpace(filePath)

    fmt.Print("Are you sure you want to delete this file? (y/n): ")
    response, _ := reader.ReadString('\n')
    response = strings.TrimSpace(strings.ToLower(response))

    if response == "y" {
        err := os.Remove(filePath)
        if err != nil {
            fmt.Println("Error deleting file:", err)
            return
        }
        fmt.Println("File deleted successfully!")
    } else {
        fmt.Println("File deletion cancelled.")
    }
}
