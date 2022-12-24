package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open(os.Args[1])
	scan := bufio.NewScanner(file)
	scan.Split(bufio.ScanLines)

	totalOverlaps := 0
	for scan.Scan() {
		line := scan.Text()
		r1, r2 := GetRanges(line)
		if r1.Contains(r2) || r2.Contains(r1) {
			totalOverlaps++
		}
	}
	fmt.Println(totalOverlaps)
}

func GetRanges(s string) (Range, Range) {
	spl := strings.Split(s, ",")
	r1 := NewRange(spl[0])
	r2 := NewRange(spl[1])
	return r1, r2
}

type Range struct {
	Start int
	End   int
}

func (r Range) Contains(o Range) bool {
	return o.Start >= r.Start && o.End <= r.End
}

func NewRange(in string) Range {
	sp := strings.Split(in, "-")
	r := Range{}
	r.Start, _ = strconv.Atoi(sp[0])
	r.End, _ = strconv.Atoi(sp[1])
	return r
}
