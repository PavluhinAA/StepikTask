package main

import (
	"fmt"
	"strings"
)

type Stringer interface {
	String() string
}

type battery struct {
	zeus string
}

func (b *battery) String() string {
	var res string
	for i := 0; i < strings.Count(b.zeus, "0"); i++ {
		res += "\t"
	}
	for i := 0; i < strings.Count(b.zeus, "1"); i++ {
		res += "X"
	}
	return res
}
func main() {
	var zeus string
	fmt.Scan(&zeus)
	batteryForTest := battery{zeus: zeus}

}
