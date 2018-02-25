package b64tohex

import "testing"

func TestHexToB64(t *testing.T) {
	tt := []struct {
		have string
		want string
	}{
		{
			have: "48656c6c6f20476f7068657221",
			want: "SGVsbG8gR29waGVyIQ==",
		},
		{
			have: "476f70686572206a757374206475672068696d73656c66206120686f6c652e",
			want: "R29waGVyIGp1c3QgZHVnIGhpbXNlbGYgYSBob2xlLg==",
		},
		{
			have: "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d",
			want: "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t",
		},
	}

	for _, tc := range tt {
		if b64, _ := HexToB64(tc.have); b64 != tc.want {
			t.Errorf("have %s; want %s", b64, tc.want)
		}
	}
}
