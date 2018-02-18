/*
Problem Statement: Bubble Pop

    Our users need to be able to Pop Bubbles in a fun, efficient, and programmatically sound matter.

Version the first:
    Ours users require a 3x3 grid, with bubbles that will pop. When they pop a bubble,
    they
    Naturally expect all of the like, adjacent bubbles to pop as well. Bubbles look
    like either an X, or an O.

Inputs:
P|O|P
O|O|P
P|P|O

(i,j)

Output:
Board with bubbles popped in accordance with the rules provided
P|X|P
X|X|P
P|P|O

change current coordinate to X
add 1 to first coordinate and check if out of bounds
if still in board
  check if bubble type equals current bubble type
  if true
    change current coordinate to X
subtract 1 from first coordinate and check if out of bounds
if still in board
  check if bubble type equals current bubble type
  if true
    change current coordinate to X
add 1 to second coordinate and check if out of bounds
if still in board
  check if bubble type equals current bubble type
  if true
    change current coordinate to X
subtract 1 second coordinate and check if out of bounds
if still in board
  check if bubble type equals current bubble type
  if true
    change current coordinate to X

Version the second:
    Our users need gravity; They're used to it in their daily lives and the need it in
    their bubble
    simulators. Implement gravity such that when bubbles pop, those that remain fall
    down - and
    new bubbles fall from the sky to replace those lost. In accordance with the
    prophecy.

Output:
Board with bubbles popped in accordance with the rules provided

X|X|P
P|X|P
P|P|O

Loop over the last line
if bubble is popped:
  check if bubble above exists:
  if not continue checking above
    if no bubbles exist above:
      generate one


O|O|P
P|P|P
P|P|O
(this is strictly an example, not perscriptive)
*/

package main

import "fmt"
import "math"
import "math/rand"
import "time"

type Board map[[2]int]string

func main() {

	board := Board{
		[2]int{0, 0}: "P",
		[2]int{0, 1}: "P",
		[2]int{0, 2}: "O",
		[2]int{1, 0}: "O",
		[2]int{1, 1}: "O",
		[2]int{1, 2}: "P",
		[2]int{2, 0}: "P",
		[2]int{2, 1}: "O",
		[2]int{2, 2}: "P",
	}

	x := 1
	y := 1

	printBoard(board)

	board, err := popBubble(board, x, y)
	if err != nil {
		fmt.Errorf("unable to pop bubble: %s", err)
	}

	seed := int64(time.Now().UnixNano())

	board = gravitate(board, seed)

	printBoard(board)
}

func popBubble(board Board, x, y int) (Board, error) {
	if val, ok := board[[2]int{x, y}]; ok {
		if board[[2]int{x + 1, y}] == val {
			board[[2]int{x + 1, y}] = "X"
		}
		if board[[2]int{x - 1, y}] == val {
			board[[2]int{x - 1, y}] = "X"
		}
		if board[[2]int{x, y + 1}] == val {
			board[[2]int{x, y + 1}] = "X"
		}
		if board[[2]int{x, y - 1}] == val {
			board[[2]int{x, y - 1}] = "X"
		}
		board[[2]int{x, y}] = "X"
	}
	return board, nil
}

func gravitate(b Board, s int64) Board {
	boardSize := int(math.Sqrt(float64(len(b))))
	toFall := 0

	for x := 0; x < boardSize; x++ {
		for y := 0; y < boardSize; y++ {
			// start counting how far the end bubbles will need to fall
			if b[[2]int{x, y}] == "X" {
				toFall++
			}
			// have existing bubbles fall past all popped bubbles
			if b[[2]int{x, y}] != "X" {
				b[[2]int{x, y}], b[[2]int{x, y - toFall}] = b[[2]int{x, y - toFall}], b[[2]int{x, y}]
			}
			// If we have reached the top row of the board, we need to start pushing bubbles down
			if y+1 == boardSize {
				for toFall > 0 || b[[2]int{x, y}] == "X" {
					b[[2]int{x, y - toFall}] = genBubble(s)
					toFall--
				}
				toFall = 0
			}
		}
	}
	return b
}

func genBubble(s int64) string {
	source := rand.NewSource(s)
	r := rand.New(source)
	seed := r.Intn(2)
	if seed == 1 {
		return "P"
	}
	return "O"
}

func printBoard(b Board) {
	boardSize := int(math.Sqrt(float64(len(b))))
	for y := boardSize - 1; y > -1; y-- {
		for x := 0; x < boardSize; x++ {
			fmt.Printf("%s", b[[2]int{x, y}])
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}
