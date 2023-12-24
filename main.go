package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func printTree(root string, prefix string) {
	files, err := os.ReadDir(root)
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
		if info, err := file.Info(); err == nil && info.IsDir() {
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
