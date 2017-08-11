package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	proto := `1233213
	123
	31
	31
	3123
	`

	reader := strings.NewReader(proto)
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
