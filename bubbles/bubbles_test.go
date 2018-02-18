package main

import (
	"reflect"
	"testing"
)

func TestPopBubble(t *testing.T) {
	tt := []struct {
		have        Board
		coordinates [2]int
		want        Board
	}{
		{
			have: Board{
				[2]int{0, 0}: "P",
				[2]int{0, 1}: "P",
				[2]int{0, 2}: "O",
				[2]int{1, 0}: "O",
				[2]int{1, 1}: "O",
				[2]int{1, 2}: "P",
				[2]int{2, 0}: "P",
				[2]int{2, 1}: "O",
				[2]int{2, 2}: "P",
			},
			coordinates: [2]int{1, 1},
			want: Board{
				[2]int{0, 0}: "P",
				[2]int{0, 1}: "P",
				[2]int{0, 2}: "O",
				[2]int{1, 0}: "X",
				[2]int{1, 1}: "X",
				[2]int{1, 2}: "P",
				[2]int{2, 0}: "P",
				[2]int{2, 1}: "X",
				[2]int{2, 2}: "P",
			},
		},
		{
			have: Board{
				[2]int{0, 0}: "P",
				[2]int{0, 1}: "P",
				[2]int{0, 2}: "O",
				[2]int{0, 3}: "O",
				[2]int{1, 0}: "O",
				[2]int{1, 1}: "O",
				[2]int{1, 2}: "P",
				[2]int{1, 3}: "P",
				[2]int{2, 0}: "P",
				[2]int{2, 1}: "O",
				[2]int{2, 2}: "P",
				[2]int{2, 3}: "P",
				[2]int{3, 0}: "O",
				[2]int{3, 1}: "O",
				[2]int{3, 2}: "O",
				[2]int{3, 3}: "P",
			},
			coordinates: [2]int{2, 3},
			want: Board{
				[2]int{0, 0}: "P",
				[2]int{0, 1}: "P",
				[2]int{0, 2}: "O",
				[2]int{0, 3}: "O",
				[2]int{1, 0}: "O",
				[2]int{1, 1}: "O",
				[2]int{1, 2}: "P",
				[2]int{1, 3}: "X",
				[2]int{2, 0}: "P",
				[2]int{2, 1}: "O",
				[2]int{2, 2}: "X",
				[2]int{2, 3}: "X",
				[2]int{3, 0}: "O",
				[2]int{3, 1}: "O",
				[2]int{3, 2}: "O",
				[2]int{3, 3}: "X",
			},
		},
	}

	for _, tc := range tt {
		popped, err := popBubble(tc.have, tc.coordinates[0], tc.coordinates[1])
		if err != nil {
			t.Errorf("error thrown during popping: %s", err)
		}
		if !reflect.DeepEqual(tc.want, popped) {
			t.Errorf("have: %v; want: %v", popped, tc.want)
		}
	}
}

func TestGenBubble(t *testing.T) {
	tt := []struct {
		source int64
		want   string
	}{
		{
			source: int64(1),
			want:   "P",
		},
		{
			source: int64(2),
			want:   "O",
		},
	}

	for _, tc := range tt {
		for i := 0; i < 10; i++ {
			if genBubble(tc.source) != tc.want {
				t.Errorf("have source: %v; want: %s", tc.source, tc.want)
			}
		}
	}
}

func TestGravitate(t *testing.T) {
	tt := []struct {
		have Board
		seed int64
		want Board
	}{
		{
			have: Board{
				[2]int{0, 0}: "P",
				[2]int{0, 1}: "P",
				[2]int{0, 2}: "O",
				[2]int{1, 0}: "X",
				[2]int{1, 1}: "X",
				[2]int{1, 2}: "P",
				[2]int{2, 0}: "P",
				[2]int{2, 1}: "X",
				[2]int{2, 2}: "P",
			},
			seed: int64(1),
			want: Board{
				[2]int{0, 0}: "P",
				[2]int{0, 1}: "P",
				[2]int{0, 2}: "O",
				[2]int{1, 0}: "P",
				[2]int{1, 1}: "P",
				[2]int{1, 2}: "P",
				[2]int{2, 0}: "P",
				[2]int{2, 1}: "P",
				[2]int{2, 2}: "P",
			},
		},
		{
			have: Board{
				[2]int{0, 0}: "P",
				[2]int{0, 1}: "P",
				[2]int{0, 2}: "O",
				[2]int{0, 3}: "O",
				[2]int{1, 0}: "O",
				[2]int{1, 1}: "O",
				[2]int{1, 2}: "P",
				[2]int{1, 3}: "X",
				[2]int{2, 0}: "P",
				[2]int{2, 1}: "O",
				[2]int{2, 2}: "X",
				[2]int{2, 3}: "X",
				[2]int{3, 0}: "O",
				[2]int{3, 1}: "O",
				[2]int{3, 2}: "O",
				[2]int{3, 3}: "X",
			},
			seed: int64(2),
			want: Board{
				[2]int{0, 0}: "P",
				[2]int{0, 1}: "P",
				[2]int{0, 2}: "O",
				[2]int{0, 3}: "O",
				[2]int{1, 0}: "O",
				[2]int{1, 1}: "O",
				[2]int{1, 2}: "P",
				[2]int{1, 3}: "O",
				[2]int{2, 0}: "P",
				[2]int{2, 1}: "O",
				[2]int{2, 2}: "P",
				[2]int{2, 3}: "O",
				[2]int{3, 0}: "O",
				[2]int{3, 1}: "O",
				[2]int{3, 2}: "O",
				[2]int{3, 3}: "O",
			},
		},
	}

	for _, tc := range tt {
		printBoard(tc.have)
		if !reflect.DeepEqual(gravitate(tc.have, tc.seed), tc.want) {
			printBoard(tc.have)
			printBoard(tc.want)
			t.Errorf("have board %v; have seed: %d; want board %v", tc.have, tc.seed, tc.want)
		}
	}
}
