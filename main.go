// package main
package main

import (
	"fmt"
	"os"

	"conversor/rates"
)

func main() {
	var ratesMap any
	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Println("Invalid numbers of arguments")
		fmt.Println("Usage: ./conversor <valor_em_blr> <moeda_destino>")
	}
	if m, err := rates.UploadRates(); err == nil {
		ratesMap = m["rates"]
	}
	fmt.Println(ratesMap)
}
