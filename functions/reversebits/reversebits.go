package main

import "fmt"

// ReverseBits reverses the bits of a given byte
func ReverseBits(oct byte) byte {
	var result byte
	for i := 0; i < 8; i++ {
		// Shift result to the left to make space for the next bit
		result <<= 1
		// Extract the least significant bit from oct and add it to result
		result |= oct & 1
		// Shift oct to the right to process the next bit
		oct >>= 1
	}
	return result
}

func main() {
	fmt.Println(ReverseBits(1))
}