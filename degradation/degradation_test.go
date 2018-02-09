package degradation

import (
	"testing"
)

func TestDegrade(t *testing.T) {
	tt := []struct {
		haveStr string
		haveI1  int
		haveI2  int
		want    string
	}{
		{
			"filibustering", 9, 2, "niribustelifg",
		},
		{
			"renaissance", 1, 5, "rsianesance",
		},
	}

	for _, tc := range tt {
		if got := Degrade(tc.haveStr, tc.haveI1, tc.haveI2); got != tc.want {
			t.Errorf("Have: %s; want %s", got, tc.want)
		}
	}
}

func TestreverseString(t *testing.T) {
	tt := []struct {
		have string
		want string
	}{
		{
			have: "hello",
			want: "olleh",
		},
		{
			have: "wassup",
			want: "pussaw",
		},
	}
	for _, tc := range tt {
		if got := reverseString(tc.have); got != tc.want {
			t.Errorf("Have: %s; want %s", got, tc.want)
		}
	}
}
