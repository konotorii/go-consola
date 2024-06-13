package consola

import (
	"fmt"
	"github.com/TwiN/go-color"
	"time"
)

func log(msg ...any) {
	fmt.Println(color.Ize(color.White, time.Now().Format(time.RFC3339)), "[LOG]", msg)
}

func info(msg ...any) {
	fmt.Println(color.Ize(color.White, time.Now().Format(time.RFC3339)), "[INFO]", msg)
}

func debug(msg ...any) {
	fmt.Println(color.Ize(color.White, time.Now().Format(time.RFC3339)), color.Ize(color.Gray, "[DEBUG]"), msg)
}

func success(msg ...any) {
	fmt.Println(color.Ize(color.White, time.Now().Format(time.RFC3339)), color.Ize(color.Green, "[SUCCESS]"), msg)
}

func warning(msg ...any) {
	fmt.Println(color.Ize(color.White, time.Now().Format(time.RFC3339)), color.Ize(color.Yellow, "[WARNING]"), msg)
}

func error(msg ...any) {
	fmt.Println(color.Ize(color.White, time.Now().Format(time.RFC3339)), color.Ize(color.Red, "[ERROR]"), msg)
}
