package conversion

import (
	"fmt"
	"strconv"
	"unsafe"
	"strings"
	"github.com/aratan/srv4eth/saldo"

)

func Show() float32 {
	s := strconv.Itoa(saldo.Show())
	a := strings.Replace(s,"0x", "",16)
	n, err := strconv.ParseUint(a, 16, 32)

	if err != nil {
		panic(err)
	}

	n2 := uint32(n)
	f := *(*float32)(unsafe.Pointer(&n2))
	fmt.Println(f)
	return f
}
