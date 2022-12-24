package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open(os.Args[1])
	scan := bufio.NewScanner(file)
	scan.Split(bufio.ScanLines)

	stacks := 0
	var ship Ship
	loadSection := true
	for scan.Scan() {
		if stacks == 0 {
			// stacks are 4 cols (incl spaces) minus the last one which is only
			// 3, so just add one and divide by four to get stack count
			stacks = (len(scan.Text()) + 1) / 4
			ship = make([]*list.List, stacks)
		}

		if scan.Text() == "" {
			loadSection = false
			continue
		}

		if loadSection {
			load := parseContainerLine(scan.Text())
			// 0-indexed stack
			for stack, val := range load {
				if ship[stack] == nil {
					ship[stack] = &list.List{}
				}
				ship[stack].PushBack(val)
			}
		} else {
			s := strings.Split(scan.Text(), " ")
			stack, err := strconv.Atoi(s[1])
			if err != nil {
				panic(err)
			}
			from, err := strconv.Atoi(s[3])
			if err != nil {
				panic(err)
			}
			to, err := strconv.Atoi(s[5])
			if err != nil {
				panic(err)
			}

			ship.Move(stack, from, to)
		}
	}

	fmt.Println(ship)

	message := ""
	for _, stack := range ship {
		top := stack.Front()
		message += string(top.Value.(byte))
	}
	fmt.Println(message)
}

type Ship []*list.List

func (s Ship) String() string {
	build := &strings.Builder{}
	for stack := 1; stack <= len(s); stack++ {
		build.WriteString(fmt.Sprintf(" %d  ", stack))
	}
	build.WriteString(fmt.Sprintln())
	curr := make([]*list.Element, len(s))
	for i, stack := range s {
		curr[i] = stack.Back()
	}

	// print until we're all nil
	for {
		allNil := true
		for i, crate := range curr {
			if crate == nil {
				build.WriteString("    ")
				continue
			}
			allNil = false
			val := string(crate.Value.(byte))
			build.WriteString(fmt.Sprintf("[%s] ", val))
			curr[i] = crate.Prev()
		}
		build.WriteString(fmt.Sprintln())
		if allNil {
			break
		}
	}
	return build.String()
}

// Move a number of crates from one stack to another. Crate number must be
// 1-indexed like the input.
func (s Ship) Move(num, from, to int) {
	realFrom := from - 1
	realTo := to - 1
	if s[realFrom].Len() < num {
		panic(fmt.Sprintf("stack not large enough, %d < %d", s[realFrom].Len(), num))
	}

	for i := 0; i < num; i++ {
		elem := s[realFrom].Front()
		s[realTo].PushFront(elem.Value)
		s[realFrom].Remove(elem)
	}
}

func parseContainerLine(line string) map[int]byte {
	result := map[int]byte{}
	for i := 1; i < len(line); i += 4 {
		if line[i] == 32 {
			continue
		} else if line[i] == 49 {
			// a number? must be the container number
			return result
		}
		stack := (i - 1) / 4
		result[stack] = line[i]
	}
	return result
}
