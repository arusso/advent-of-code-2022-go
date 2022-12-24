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
	fmt.Println(startOfPacket(scan.Text()))
}

const StartOfPacketLength = 4

// startOfPacket marker looks for the first four unique characters, returning
// the index of the last character of the marker.
func startOfPacket(str string) int {
	for i := StartOfPacketLength - 1; i < len(str); i++ {
		duplicate := false
		counts := map[byte]int{}
		for j := 0; j < StartOfPacketLength; j++ {
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
