package main

import (
	"fmt"

	"github.com/dnnyjns/migraine/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		fmt.Println(err)
	}
}
