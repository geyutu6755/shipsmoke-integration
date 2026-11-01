package main

import (
	"fmt"
	"os"
	"strconv"
)

const version = "0.3.1"

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: teapot [--version|--help|exit N]")
		return
	}
	switch os.Args[1] {
	case "--version":
		fmt.Printf("teapot version %s\n", version)
	case "--help":
		fmt.Println("teapot — shipsmoke integration test subject")
	case "exit":
		if len(os.Args) > 2 {
			n, _ := strconv.Atoi(os.Args[2])
			os.Exit(n)
		}
	default:
		fmt.Fprintf(os.Stderr, "teapot: unknown command %q\n", os.Args[1])
		os.Exit(2)
	}
}
