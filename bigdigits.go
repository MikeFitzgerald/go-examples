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
	{"  99999",
		" 9    9",
		" 9    9",
		"  99999",
		"    99 ",
		"   99  ",
		"  99   "},
}

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("usage: %s <whole-number>\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	stringOfDigits := os.Args[1]
	for row := range bigDigits[0] {
		line := ""
		for column := range stringOfDigits {
			digit := stringOfDigits[column] - '0'
			if 0 <= digit && digit <= 9 {
				line += bigDigits[digit][row] + " "
			} else {
				log.Fatal("invalid whole number")
			}
		}
		fmt.Println(line)
	}
}
