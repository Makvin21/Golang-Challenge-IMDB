package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("data.tsv")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	data := make([]byte, 64)

	for {
		n, err := file.Read(data)
		if err == io.EOF { // если конец файла
			break // выходим из цикла
		}
		fmt.Print(string(data[:n]))
	}
}
