package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	// if len(os.Args) > 1 {

	// 	data := &structure.ArtData{} // объявление структуры данных для сбора и сохранения
	// 	data.Args = os.Args[1]
	// 	input.ReadFile(data)    // запускаем функцию и полученнные данные сохраняем в data
	// 	input.CheckSymbol(data) //check symbols in range ascii of 32 to 126
	// 	input.SplitText(data)   // передали аргументы чтобы манипулировать ими

	// }
	if len(os.Args) != 2 {
		fmt.Println("No arguments")
		return
	}
	argMain := os.Args[1]

	argsArr := strings.Split(argMain, "\\n")
	file, _ := ioutil.ReadFile("fonts/standard.txt")
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
				if el < 32 || el > 126 {
					fmt.Println("ERROR")
					return
				}
				n := (el-32)*9 + 1
				fmt.Print(arr[int(n)+x])
			}
			fmt.Println()
		}
	}

}
