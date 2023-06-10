// Om Ganeshaya Namah!!
package main

import (
	"fmt"
	"os"

	"github.com/ManianVSS/karta.go.poc/xlang"
)

func main() {
	fmt.Println("Xlang parser example!!")

	if len(os.Args) < 2 {
		panic("No file name provided to run")
	}
	xlang.Main(os.Args[1])
}
