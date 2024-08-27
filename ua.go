package main

import (
    "fmt"
    "os"
    "strconv"
    "github.com/brianvoe/gofakeit/v6"
    "strings"
    "io/ioutil"
)

func main() {
    if len(os.Args) != 3 {
        fmt.Println("Usage: go run ua.go [save file] [amount]")
        return
    }

    saveFile := os.Args[1]
    amount, err := strconv.Atoi(os.Args[2])
    if err != nil {
        fmt.Println("Please provide a valid number for amount")
        return
    }

    var userAgents []string

    for i := 0; i < amount; i++ {
        userAgents = append(userAgents, gofakeit.UserAgent())
    }

    err = ioutil.WriteFile(saveFile, []byte(strings.Join(userAgents, "\n")), 0644)
    if err != nil {
        fmt.Println("Error writing to file:", err)
        return
    }

    fmt.Printf("%d user agents saved to %s\n", amount, saveFile)
}
