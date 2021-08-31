package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const (
	FILE_LEN = 6623
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
	argsArr := strings.Split(strings.ReplaceAll(argMain, "\\n", "\n"), "\n")
	file, err := ioutil.ReadFile("fonts/standard.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if len(file) != FILE_LEN {
		fmt.Println("File is corrupted")
		return
	}

	arr := []string{}
	for _, el := range strings.Split(string(file), string('\n')) {
		arr = append(arr, el)
	}

	for _, arg := range argsArr {
		if arg == "" {
			fmt.Println()
			continue
		}
		for x := 0; x < 8; x++ {
			for _, el := range arg {
				n := (el-32)*9 + 1
				fmt.Print(arr[int(n)+x])
			}
			fmt.Println()
		}
		fmt.Println()
	}

}

func isValid(s string) bool {
	for _, ch := range s {
		if ch < ' ' && ch != 10 || ch > '~' {
			return false
		}
	}
	return true
}
