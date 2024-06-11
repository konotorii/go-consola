package consola

import (
	"fmt"
)

func log(msg ...any) {
	fmt.Println(color.Ize(color.Red, ""))
}

func info() {
	fmt.Println("")
}

func debug() {
	fmt.Println("")
}

func warning() {
	fmt.Println("")
}

func error() {
	fmt.Println("")
}
