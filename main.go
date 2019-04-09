package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {

	var b bytes.Buffer

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		b.WriteString(scanner.Text())
	}

	s := b.String()
	s = strings.Replace(s, `\n`, "\n", -1)
	s = strings.Replace(s, `\t`, "\t", -1)

	fmt.Printf("%s", s)
}
