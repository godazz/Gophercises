package internals

import (
	"encoding/binary"
	"fmt"
	"os"
)

func Exitf(format string, a ...interface{}) {
	fmt.Println(format, a)
	os.Exit(1)
}

func Itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}
