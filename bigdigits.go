package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

var bigDigits = [][]string{
	{"  000  ",
		" 0   0 ",
		"0     0",
		"0     0",
		"0     0",
		" 0   0 ",
		"  000  "},
	{"  111  ",
		"11111  ",
		"  111  ",
		"  111  ",
		"  111  ",
		"  111  ",
		"1111111"},
	{"  2222 ",
		"22   22",
		"    22 ",
		"   22  ",
		"  22   ",
		" 22    ",
		"2222222"},
	{" 33333 ",
		"33   33",
		"    33 ",
		"   33  ",
		"    33 ",
		"33   33",
		" 33333 "},
	{"    44 ",
		"   444 ",
		"  4 44 ",
		" 4  44 ",
		"4444444",
		"    44 ",
		"    44 "},
	{"5555555",
		"55     ",
		"55     ",
		" 55555 ",
		"     55",
		"55   55",
		" 55555 "},
	{"   66  ",
		"  66   ",
		" 66    ",
		"666666 ",
		"66   66",
		" 66  66",
		"  666  "},
	{"7777777",
		"     77",
		"    77 ",
		"    77 ",
		"   77  ",
		"   77  ",
		"   77  "},
	{"  888  ",
		" 8   8 ",
		" 8   8 ",
		"  888  ",
		" 8   8 ",
		" 8   8 ",
		"  888  "},
	{"  9999 ",
		" 9    9",
		" 9    9",
		"  99999",
		"    99 ",
		"   99  ",
		"  99   "},
}

func main() {

	// If only one argument is passed (Args[0] holds program's name)
	if len(os.Args) == 1 {
		fmt.Printf("usage: %s <whole-number>\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	// Set the stringOfDigits to be the second argument of the command-line
	stringOfDigits := os.Args[1]

	// For a variable row = 0, loop over the items in the first slot of bigDigits
	// for row = 0; row < len(bigDigits[0]); row++
	for row := range bigDigits[0] {

		// Set line to be an empty string
		line := ""

		// Iterate over each digit in stringOfDigits
		for column := range stringOfDigits {

			// stringOfDigits[0], for example, returns a single character represented by a byte (since that's what strings are)
			digit := stringOfDigits[column] - '0'

			// If the digit is between 0 and 9...
			// Go converts between uint8 and byte
			if 0 <= digit && digit <= 9 {
				// Why concatenate the space?
				// Because it adds a space at the end of each row (and therefore between each big digit)
				line += bigDigits[digit][row] + " "
			} else {
				log.Fatal("invalid whole number")
			}
		}
		fmt.Println(line)
	}
}
