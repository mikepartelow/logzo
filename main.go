package main

import (
	"fmt"

	"github.com/mikepartelow/logzo/pkg/logging"
	"golang.org/x/exp/slog"
)

func main() {
	slog.Info("first")
	fmt.Println("hi 0")

	logging.Init("./log.txt")
	slog.Info("second")
	fmt.Println("hi 1")

	slog.Info("third")
	fmt.Println("hi 2")

	_ = logging.Rotate()

	slog.Info("fourth")
	fmt.Println("hi 3")
}
