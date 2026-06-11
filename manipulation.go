package main

import (
    "os"
)

func appendToFile(filepath, destfile string) {
    srcContent, err := os.ReadFile(filepath)
    if err != nil {
        panic(err)
    }

    destFile, err := os.OpenFile(destfile, os.O_APPEND|os.O_WRONLY, 0644)
    if err != nil {
        panic(err)
    }
    defer destFile.Close()

    _, err = destFile.Write(srcContent)
    if err != nil {
        panic(err)
    }
}
