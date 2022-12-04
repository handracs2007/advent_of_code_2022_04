package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Read the input file.
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("unable to open input file: %s", err)
	}
	defer f.Close()

	// Create the buffered reader to read the file content.
	r := bufio.NewReader(f)

	// Initial counter variable for both part1 and part2.
	toConsider1, toConsider2 := 0, 0
	for {
		l, err := r.ReadString('\n')
		if err == io.EOF {
			break
		}

		l = strings.TrimSpace(l)   // Clear the line from any whitespace
		r := strings.Split(l, ",") // Separate to get the first and second range.

		r1 := strings.Split(r[0], "-") // Separate the first range.
		r11, _ := strconv.Atoi(r1[0])  // Get the first number of first range. We ignore the error here as the input is always valid.
		r12, _ := strconv.Atoi(r1[1])  // Get the second number of first range. We ignore the error here as the input is always valid.

		r2 := strings.Split(r[1], "-") // Separate the second range.
		r21, _ := strconv.Atoi(r2[0])  // Get the first number of second range. We ignore the error here as the input is always valid.
		r22, _ := strconv.Atoi(r2[1])  // Get the first number of second range. We ignore the error here as the input is always valid.

		// For the first case, we need to find the range that is totally within another range.
		if (r11 >= r21 && r11 <= r22 && r12 >= r21 && r12 <= r22) || (r21 >= r11 && r21 <= r12 && r22 >= r11 && r22 <= r12) {
			toConsider1++
		}

		// For the second case, we just need to ensure that the minimum or maximum value of a range is within the range of another range.
		if (r11 >= r21 && r11 <= r22) || (r12 >= r21 && r12 <= r22) || (r21 >= r11 && r21 <= r12) || (r22 >= r11 && r22 <= r12) {
			toConsider2++
		}
	}

	fmt.Println(toConsider1)
	fmt.Println(toConsider2)
}
