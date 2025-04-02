package main

import (
	"fmt"
	"github.com/0xweb-3/cobra_example/cmd"
	"os"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
