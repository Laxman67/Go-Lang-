package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("index.html")
	if err != nil {
		fmt.Println(err)
		return
	}

	fileInfo, err := f.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}

	buf := make([]byte, fileInfo.Size())
	d, err := f.Read(buf)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(buf))
	fmt.Println(d)

	f.Close()

	// fmt.Println(fileInfo.Name())
	// fmt.Println(fileInfo.Size())
	// fmt.Println(fileInfo.IsDir())

	// fmt.Println(fileInfo.ModTime())

	// fmt.Println(fileInfo.Mode())

}
