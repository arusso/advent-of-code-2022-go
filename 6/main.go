package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open(os.Args[1])
	scan := bufio.NewScanner(file)
	scan.Split(bufio.ScanLines)

	// input is a single line
	scan.Scan()
	fmt.Println(startOfMessage(scan.Text()))
}

const StartOfPacketLength = 4
const StartOfMessageLength = 14

func afterUnique(str string, uniq int) int {
	for i := uniq - 1; i < len(str); i++ {
		duplicate := false
		counts := map[byte]int{}
		for j := 0; j < uniq; j++ {
			counts[str[i-j]]++
		}

		// this could be more efficent...
		for _, count := range counts {
			if count > 1 {
				// this was a duplcate, break here
				duplicate = true
				break
			}
		}

		if !duplicate {
			return i + 1
		}
	}

	return -1
}

func startOfPacket(str string) int {
	return afterUnique(str, StartOfPacketLength)
}

func startOfMessage(str string) int {
	return afterUnique(str, StartOfMessageLength)
}
