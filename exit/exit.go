package exit

import "os"

// Exit для остановки и выхода из программы
func Exit(s string) {
	os.Stdout.WriteString(s)
	os.Exit(0)
}
