package main

// https://adventofcode.com/2022/day/1

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

var top = [3]int{}

func main() {
	if len(os.Args) != 2 {
		handle(errors.New("unexpected number of arguments"))
	}
	filename := os.Args[1]
	file, err := os.Open(filename)
	handle(err)

	scan := bufio.NewScanner(file)
	scan.Split(bufio.ScanLines)

	elfCalories := []int{0}
	elf := 0
	for scan.Scan() {
		if scan.Text() == "" {
			updateTop(elfCalories[elf])
			elf++
			elfCalories = append(elfCalories, 0)
			continue
		}

		cals, err := strconv.Atoi(scan.Text())
		handle(err)

		elfCalories[elf] += cals
	}
	updateTop(elfCalories[elf])

	total := 0
	for val := range top {
		total += top[val]
	}

	fmt.Printf("Calories carried by top %d elves: %d\n", len(top), total)
}

func updateTop(value int) {
	if value <= top[len(top)-1] {
		return
	}

	for i := len(top) - 1; i > 0; i-- {
		if value < top[i-1] {
			top[i] = value
			return
		}
		top[i] = top[i-1]
	}

	top[0] = value
}

func handle(err error) {
	if err != nil {
		fmt.Printf("usage: %s <filename>\n", os.Args[0])
		fmt.Printf("\nerror: %s\n", err)
		os.Exit(1)
	}
}
