package main

import (
    "fmt"
    "os"
)

func readFile(name string) error {
    _, err := os.ReadFile(name)
    if err != nil {
        return fmt.Errorf("readFile(%s) failed: %w", name, err)
    }
    return nil
}

func main() {
    err := readFile("")
    fmt.Println(err)
}
