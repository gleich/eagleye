package utils

import (
	"fmt"

	"github.com/fatih/color"
)

func Success(msg string) {
	color.New(color.FgGreen).Printf("SUCCESS")
	fmt.Printf("[0000] %v", msg)
}
