package main

import (
	"bufio"
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"strings"
	"time"
)

func main() {
	find := os.Args[1]

	var clr bool
	last := time.Now()

	clear := func() {
		if clr {
			_, _ = fmt.Fprintf(os.Stderr, "\033[2K")
			clr = false
		}
	}

	out := func(str, end string) {
		width, _, err := terminal.GetSize(1)
		if err == nil && len(str) > width {
			str = str[:width - 1] + "…"
		}
		_, _ = fmt.Fprintf(os.Stderr, "%s%s", str, end)
	}

	var nr int64

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		nr++
		line := scanner.Text()
		if strings.Contains(line, find) {
			clear()
			fmt.Println(line)
		}

		now := time.Now()
		if now.After(last.Add(time.Second)) {
			last = now
			clear()
			out(fmt.Sprintf("%d: %s", nr, line), "\r")
			clr = true
		}
	}

	if err := scanner.Err(); err != nil {
		clear()
		panic(err)
	}
}
