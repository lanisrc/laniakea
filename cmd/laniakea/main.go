package main

import (
	"fmt"
	"io"
	"os"
)

const version = "dev"

func main() {
	os.Exit(Run(os.Args[1:], os.Stdout, os.Stderr))
}

func Run(args []string, out, err io.Writer) int {
	fmt.Fprintln(err, "usage: laniakea <command>")
	return 1
}
