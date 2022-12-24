package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	smallDirectoryLimit = 100000
)

var files = map[string]int{}
var dirs = map[string]int{}

func main() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()
	scan := bufio.NewScanner(file)
	scan.Split(bufio.ScanLines)

	path := NewPath("/")
	for scan.Scan() {
		line := scan.Text()
		if line[0] == '$' {
			args := strings.Split(line, " ")[1:]
			if args[0] == "cd" {
				if args[1] == ".." {
					path = path.Pop()
				} else if args[1][0] == '/' {
					path = NewPath(args[1])
				} else {
					path = path.Push(args[1])
				}
			} else if args[0] == "ls" {
				continue
			}
		} else {
			// output from previous command, which right now is always ls
			args := strings.Split(line, " ")
			if args[0] == "dir" {
				// do nothing
				continue
			} else {
				size, err := strconv.Atoi(args[0])
				if err != nil {
					panic(err)
				}
				filename := args[1]
				filepath := path.Push(filename)
				files[filepath.String()] += size

				// update directories
				for {
					filepath = filepath.Pop()
					dirs[filepath.String()] += size

					if filepath.String() == "/" {
						break
					}
				}
			}
		}
	}

	dirSizes := 0
	for _, size := range dirs {
		if size < smallDirectoryLimit {
			dirSizes += size
		}
	}

	fmt.Println("size: " + strconv.Itoa(dirSizes))
}
