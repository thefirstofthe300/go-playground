package b64tohex

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func main() {
	hello, err := HexToB64("48656c6c6f20776f726c6421")
	if err != nil {
		fmt.Errorf("unable to convert hex to base64: %s", err)
	}
	fmt.Printf(hello)
}

func HexToB64(h string) (string, error) {
	plain, err := hex.DecodeString(h)
	if err != nil {
		return "", fmt.Errorf("unable to decode hex string: %s", err)
	}
	e := base64.StdEncoding
	b64 := e.EncodeToString(plain)
	return b64, nil
}
