package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func IsValid(str byte) {
	// for _, x := range str {
	if str < ' ' || str > '~' {
		fmt.Println("no correct symbols")
		os.Exit(0)
	}
	// }
}

func main() {
	// if len(os.Args) > 1 {

	// 	data := &structure.ArtData{} // объявление структуры данных для сбора и сохранения
	// 	data.Args = os.Args[1]
	// 	input.ReadFile(data)    // запускаем функцию и полученнные данные сохраняем в data
	// 	input.CheckSymbol(data) //check symbols in range ascii of 32 to 126
	// 	input.SplitText(data)   // передали аргументы чтобы манипулировать ими

	// }
	args := os.Args[1]
	fmt.Println(args, len(args))
	file, _ := ioutil.ReadFile("fonts/standard.txt")
	IsValid(args[0])
	arr := []string{}
	for _, el := range strings.Split(string(file), string('\n')) {
		arr = append(arr, el)
	}

	for x := 0; x < 8; x++ {

		for _, el := range args {
			n := (el-32)*9 + 1
			fmt.Print(arr[int(n)+x])
		}
		fmt.Println()
	}

}
