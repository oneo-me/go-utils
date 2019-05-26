package guid

import (
	crand "crypto/rand"
	"fmt"
	mrand "math/rand"
)

func randBytes(x []byte) {
	length := len(x)
	n, err := crand.Read(x)

	if n != length || err != nil {
		for length > 0 {
			length--
			x[length] = byte(mrand.Int31n(256))
		}
	}
}

// New 创建一个新的 GUID
func New(space string) string {
	var x [16]byte
	randBytes(x[:])
	x[6] = (x[6] & 0x0F) | 0x40
	x[8] = (x[8] & 0x3F) | 0x80

	return fmt.Sprintf("%02x%02x%02x%02x"+space+"%02x%02x"+space+"%02x%02x"+space+"%02x%02x"+space+"%02x%02x%02x%02x%02x%02x",
		x[0], x[1], x[2], x[3], x[4],
		x[5], x[6],
		x[7], x[8],
		x[9], x[10], x[11], x[12], x[13], x[14], x[15])
}
