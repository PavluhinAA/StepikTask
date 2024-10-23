package main

import (
	"fmt"
	"strconv"
)

func main() {
	fn := func(input uint) uint {
		var res string
		for i := 0; i < len(strconv.Itoa(int(input))); i++ {
			if int(strconv.Itoa(int(input))[i])%2 == 0 {
				if string(strconv.Itoa(int(input))[i]) != "0" {
					res += string(strconv.Itoa(int(input))[i])
				}
			}
		}
		if res == "" {
			return 100
		}
		res2, _ := strconv.ParseInt(res, 10, 64)
		x := uint(res2)
		return x
	}
	fmt.Println(fn(12657650001313013))
}
