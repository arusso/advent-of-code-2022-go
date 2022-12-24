package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	smallDirectoryLimit    = 100000
	updateSpaceRequirement = 30000000
	diskSize               = 70000000
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

	free := diskSize - dirs["/"]
	fmt.Printf("disk size: %d\n", diskSize)
	fmt.Printf("space free: %d\n", free)
	fmt.Printf("space used: %d\n", dirs["/"])

	needed := (free - updateSpaceRequirement) * -1
	if needed < 0 {
		fmt.Println("we have enough space for the update")
		os.Exit(0)
	}

	fmt.Printf("space needed: %d\n", needed)
	smallestDir := "/"
	for dir, size := range dirs {
		//fmt.Printf(" -> inspecting dir of size %d\n", size)
		if size < dirs[smallestDir] && size >= needed {
			smallestDir = dir
		}
	}

	fmt.Println("found dir: " + smallestDir)
	fmt.Println("space freed: " + strconv.Itoa(dirs[smallestDir]))
}
