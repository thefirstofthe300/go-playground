package xor

import (
	"encoding/hex"
	"fmt"
)

func XORStrings(s1 string, s2 string) (string, error) {
	decoded1, err := hex.DecodeString(s1)
	if err != nil {
		return "", fmt.Errorf("unable to decode string: %s", err)
	}

	decoded2, err := hex.DecodeString(s2)
	if err != nil {
		return "", fmt.Errorf("unable to decode string: %s", err)
	}

	var decoded3 []byte

	for i := 0; i < len(decoded1); i++ {
		decoded3 = append(decoded3, decoded1[i]^decoded2[i])
	}

	return string(decoded3), nil
}
