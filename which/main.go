package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide more argument")
		return
	}
	file := arguments[1]
	t := os.Getenv("PATH")
	paths := filepath.SplitList(t)
	for _, path := range paths {
		fullPath := filepath.Join(path, file)
		fileInfo, err := os.Stat(fullPath)
		if err == nil {
			mode := fileInfo.Mode()
			if mode.IsRegular() {
				// fmt.Println(mode)
				if mode&0111 != 0 {
					fmt.Println(fullPath)
					return
				}
			}

		}
	}

}
