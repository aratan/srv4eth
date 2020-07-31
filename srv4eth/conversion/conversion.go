package conversion

import (
	"fmt"
	"strconv"
//	"unsafe"
	"strings"
	"github.com/aratan/srv4eth/saldo"

)

func Show() float32 {
	s,_ :=  fmt.Printf("%v",saldo.Show)

	// quita el 0x del hexadecial
	a := strings.Replace("0x123","0x","",16)
	// hex to  dec
	n, err := strconv.ParseUint(a, 16, 32)

	if err != nil {
		panic(err)
	}

	return n
}
