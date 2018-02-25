package main

import (
	"fmt"
	"testing"
)

func TestXORStrings(t *testing.T) {
	tt := []struct {
		have [2]string
		want string
	}{
		{
			have: [2]string{
				"1c0111001f010100061a024b53535009181c",
				"686974207468652062756c6c277320657965",
			},
			want: "746865206b696420646f6e277420706c6179",
		},
	}

	for _, tc := range tt {
		if got, err := XORStrings(tc.have[0], tc.have[1]); got != tc.want {
			if err != nil {
				fmt.Errorf("XORStrings threw an error: %s", err)
			} else {
				fmt.Errorf("have %s; want %s", got, tc.want)
			}
		}
	}
}
