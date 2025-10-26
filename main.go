// package main
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"conversor/rates"
)

func main() {
	ratesMap := make(map[string]any)
	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Println("Invalid numbers of arguments")
		fmt.Println("Usage: ./conversor <valor_em_blr> <moeda_destino>")
		return
	}
	blr, err := strconv.ParseFloat(strings.TrimSpace(args[0]), 32)
	if err != nil {
		fmt.Println("BRL", args[0], "value is not a valid number")
		return
	}
	if m, err := rates.UploadRates(); err == nil {
		ratesMap = m["rates"].(map[string]any)
	}
	rateCode := strings.ToUpper(args[1])
	rate, ok := ratesMap[rateCode].(float32)
	if !ok {
		fmt.Println(rateCode, "is a invalid rate code", rate, ratesMap["USD"])
		return
	}
	fmt.Println(rate * float32(blr))
}
