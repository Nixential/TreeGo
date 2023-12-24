package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func printTree(root string, prefix string) {
	files, err := ioutil.ReadDir(root)
	if err != nil {
		fmt.Println(err)
		return
	}

	for i, file := range files {
		isLast := i == len(files)-1
		fmt.Print(prefix)
		if isLast {
			fmt.Print("└── ")
		} else {
			fmt.Print("├── ")
		}
		fmt.Println(file.Name())
		if file.IsDir() {
			newPrefix := prefix
			if isLast {
				newPrefix += "    "
			} else {
				newPrefix += "│   "
			}
			printTree(filepath.Join(root, file.Name()), newPrefix)
		}
	}
}

func main() {
	rootDir := "."
	if len(os.Args) > 1 {
		rootDir = os.Args[1]
	}
	printTree(rootDir, "")
}
