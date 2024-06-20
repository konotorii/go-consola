package consola

import (
	"fmt"
	"github.com/TwiN/go-color"
	"time"
)

func Log(msg ...any) {
	fmt.Println(color.Ize(color.White, time.Now().Format(time.RFC3339)), "[LOG]", msg)
}

func Info(msg ...any) {
	fmt.Println(color.Ize(color.White, time.Now().Format(time.RFC3339)), "[INFO]", msg)
}

func Debug(msg ...any) {
	fmt.Println(color.Ize(color.White, time.Now().Format(time.RFC3339)), color.Ize(color.Gray, "[DEBUG]"), msg)
}

func Success(msg ...any) {
	fmt.Println(color.Ize(color.White, time.Now().Format(time.RFC3339)), color.Ize(color.Green, "[SUCCESS]"), msg)
}

func Warning(msg ...any) {
	fmt.Println(color.Ize(color.White, time.Now().Format(time.RFC3339)), color.Ize(color.Yellow, "[WARNING]"), msg)
}

func Error(msg ...any) {
	fmt.Println(color.Ize(color.White, time.Now().Format(time.RFC3339)), color.Ize(color.Red, "[ERROR]"), msg)
}
