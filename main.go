package main

import (
	"fmt"

	"github.com/Matt-Gleich/eagleye/pkg/config"
)

func main() {
	pname, config := config.LoadConfiguration()
	fmt.Println(pname)
	fmt.Println(config)
}
