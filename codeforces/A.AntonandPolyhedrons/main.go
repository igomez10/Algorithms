// A. Anton and Polyhedrons
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	figures := map[rune]uint64{
		'T': uint64(4),
		'C': uint64(6),
		'O': uint64(8),
		'D': uint64(12),
		'I': uint64(20),
	}

	text := ""
	fmt.Scanln(&text)
	numFigures, _ := strconv.Atoi(text)
	reader := bufio.NewReader(os.Stdin)

	counter := uint64(0)
	for i := 0; i < numFigures; i++ {
		r, _, _ := reader.ReadRune()
		//		fmt.Println("read", string(r))
		reader.ReadLine()
		counter += figures[r]
	}

	fmt.Println(counter)
}
