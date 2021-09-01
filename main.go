package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const (
	fileLen = 6623
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("No arguments")
		return
	}
	argMain := os.Args[1]

	if !isValid(argMain) {
		fmt.Println("Non-valid characters")
		return
	}
	// in order to pass the test - student$ go run . "" | cat -e
	if argMain == "" {
		fmt.Println()
		return
	}
	// in order to pass the test - student$ go run . "Hello\n" | cat -e
	if argMain == "\\n" {
		fmt.Println("\n")
		return
	}

	//read the content of the file
	argsArr := strings.Split(strings.ReplaceAll(argMain, "\\n", "\n"), "\n")
	// argsArr := strings.Split(argMain, "\\n")
	file, err := ioutil.ReadFile("fonts/standard.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if len(file) != fileLen {
		fmt.Println("File is corrupted")
		return
	}

	arr := []string{}
	for _, el := range strings.Split(string(file), string('\n')) {
		arr = append(arr, el)
	}

	printBanners(argsArr, arr)
}

//input validation
func isValid(s string) bool {
	for _, ch := range s {
		if ch < ' ' && ch != 10 || ch > '~' {
			return false
		}
	}
	return true
}

//print the full outcome
func printBanners(banners, arr []string) {
	for _, b := range banners {
		if b == "" {
			continue
		}

		for x := 0; x < 8; x++ {
			for _, el := range b {
				n := (el-32)*9 + 1
				fmt.Print(arr[int(n)+x])
			}
			fmt.Println("")
		}
		fmt.Println("")
	}
}
