package grabtest

import "encoding/hex"


// ff:
// b:
// s:
func MustHexDecodeString(s string) (b []byte) {
	var err error
	b, err = hex.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return
}


// ff:
// s:
// b:
func MustHexEncodeString(b []byte) (s string) {
	return hex.EncodeToString(b)
}
