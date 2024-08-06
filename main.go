package main

import (
    "fmt"
    "os/exec"
    "os"
    "time"
)

var BINDINGS = make(map[string]string)

func main() {
    BINDINGS["help"] = "shows this message."
    BINDINGS["newfile"] = "creates a new file."
    BINDINGS["delfile"] = "deletes file."
    BINDINGS["q"] = "exits the application."
    BINDINGS["clear"] = "clears the screen."

    var command string

    fmt.Println("---------------Explorer---------------")
    for true {
        fmt.Printf("~ ")
        fmt.Scan(&command)

        switch command {
        case "newfile":
            createFile()
        case "delfile":
            deleteFile()
        case "help":
            help()
        case "q":
            fmt.Println("Exiting program, bye :)")
            time.Sleep(1 * time.Second)
            return

        case "clear":
            clearShell()

        default:
            fmt.Printf("'%v' is not a valid command! use [help] to see a list of commands.\n", command)
        }

    }
}



func createFile()  {
    var name string
    fmt.Printf("Enter file name: ")
    fmt.Scan(&name)

    err, _ := os.Create(name)
    if err  != nil {
        fmt.Printf("File: %v created successfully!\n", name)
        return
    }

    fmt.Printf("file error: Unable to create file, reason: %v\n", err)
}


func deleteFile()  {
    var path string
    fmt.Printf("Enter path of file to delete: ")
    fmt.Scan(&path)

    out  := os.Remove(path)
    if out == nil {
        fmt.Println("File deleted successfully!")
        return
    }

    fmt.Printf("file error: Unable to delete file, reason: %v\n", out)
}


func help(){
    fmt.Println("HELP")
    for index, value := range BINDINGS {
        fmt.Println(index, " ", value)
    }
}

func clearShell(){
    var cmd *exec.Cmd
    cmd = exec.Command("cmd", "/c", "cls")

    output, err := cmd.Output()
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println("not an error: ", output)

    cmd.Stdout = os.Stdout
    cmd.Run()
}
