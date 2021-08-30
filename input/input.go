package input

import (
	"ascii-art/exit"
	"ascii-art/structure"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadFile(data *structure.ArtData) {
	f, err := os.Open("fonts/standard.txt")

	if err != nil {
		fmt.Println(err)
		return
	}
	scan(data, f)

}

// for open
func scan(data *structure.ArtData, f *os.File) {
	scanner := bufio.NewScanner(f) // передаем открытый файл

	for scanner.Scan() {
		data.CharacterArray = append(data.CharacterArray, scanner.Text()) // при сканировании сохр в массив данные из structure
	}
	if err := scanner.Err(); err != nil {
		exit.Exit("Scanner: was wrong")
	}
}

// SplitText - разбиение текста по \n
func SplitText(data *structure.ArtData) {
	data.Words = strings.Split(data.Args, "\\n") // create the text ...\n ....
}

func CheckSymbol(data *structure.ArtData) {
	for _, el := range data.Args {
		if el < ' ' && el > '~' {
			exit.Exit("no correct symbols")
		}
	}

}
