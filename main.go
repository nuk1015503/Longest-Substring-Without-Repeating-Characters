package main

import "fmt"

const input = "anviaj"

func main() {
	answser := lengthOfLongestSubstring(input)
	fmt.Println(answser)
}
func lengthOfLongestSubstring(s string) int {
	max := 0
	for i := 0; i < len(s); i++ {
		tmpByte := make(map[byte]bool)
		for j := i; j < len(s); j++ {
			if tmpByte[s[j]] {
				break
			}
			tmpByte[s[j]] = true
		}
		if len(tmpByte) > max {
			max = len(tmpByte)
		}
	}

	return max
}
