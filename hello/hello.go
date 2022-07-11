package main

import (
    "fmt"

    "go_projects/greetings"
)

func main() {
    message := greetings.Hello("Gladys")
    fmt.Println(message)
}