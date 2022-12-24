package main

import "testing"

func TestStartOfPacket(t *testing.T) {
	tests := map[string]int{
		"bvwbjplbgvbhsrlpgdmjqwftvncz":      5,
		"nppdvjthqldpwncqszvftbrmjlhg":      6,
		"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg": 10,
		"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw":  11,
	}

	for input, output := range tests {
		t.Run("Checking "+input, func(t *testing.T) {
			sop := startOfPacket(input)
			if sop != output {
				t.Errorf("should have gotten start of packet == %d, got %d", output, sop)
			}
		})
	}
}
