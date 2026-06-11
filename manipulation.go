package main

import (
    "os"
	"fmt"
)

func appendToFile(filepath, destfile string) error {
    srcContent, err := os.ReadFile(filepath)
    if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("File to append does not exist!")
			return err
		}
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

	return nil
}
