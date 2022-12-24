package main

// https://adventofcode.com/2022/day/2

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type choice int

const (
	rock choice = iota
	paper
	scissors
)

const win = 6
const loss = 0
const draw = 3

// Against will play your choice against theirs, and return the number of points
// scored.
func (c choice) Against(their choice) int {
	if c == their {
		// draw
		return draw + choicePoints[c]
	} else if (c == rock && their == paper) ||
		(c == scissors && their == rock) ||
		(c == paper && their == scissors) {
		// loss
		return loss + choicePoints[c]
	}

	// win
	return win + choicePoints[c]

}

var choicePoints = map[choice]int{
	rock:     1,
	paper:    2,
	scissors: 3,
}

type round struct {
	Theirs choice
	Yours  choice
}

func (r *round) Score() int {
	return r.Yours.Against(r.Theirs)
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	scan := bufio.NewScanner(file)
	scan.Split(bufio.ScanLines)

	score := 0
	rounds := 0
	for scan.Scan() {
		choices := strings.Split(scan.Text(), " ")
		theirs := parseTheirs(choices[0])
		mine := selectOurs(theirs, choices[1])
		score += mine.Against(theirs)
		rounds++
	}
	fmt.Printf("score: %d\n", score)
	fmt.Printf("rounds: %d\n", rounds)
}

func parseTheirs(c string) choice {
	if c == "A" {
		return rock
	} else if c == "B" {
		return paper
	} else if c == "C" {
		return scissors
	}
	panic(fmt.Sprintf("c cannot be %s", c))
}

func selectOurs(theirs choice, c string) choice {
	switch c {
	case "X":
		if theirs == rock {
			return scissors
		} else if theirs == scissors {
			return paper
		} else if theirs == paper {
			return rock
		}
	case "Y":
		return theirs
	case "Z":
		if theirs == rock {
			return paper
		} else if theirs == paper {
			return scissors
		} else if theirs == scissors {
			return rock
		}
	}

	panic("should have made a selection by now")
}
