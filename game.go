package main

import (
	"bufio"
	"fmt"
	tm "github.com/buger/goterm"
	"github.com/ggwhite/bowlling-game/bowlling"
	"os"
	"strconv"
)

func main() {
	br := bufio.NewReader(os.Stdin)
	game := bowlling.New()

	// for i := 0; i < 16; i++ {
	// 	game.Roll(1)
	// }

	exit := false

	tm.Clear()

	for !exit {
		tm.MoveCursor(1, 1)
		fmt.Fprintln(tm.Screen, "Bowlling Game Started")

		game.ShowScore(tm.Screen)
		tm.Flush()

		if game.Over() {
			exit = true
			continue
		}

		left := game.LeftPins()
		fmt.Fprintf(tm.Screen, "Left Pins: %2d\n", left)
		fmt.Fprintln(tm.Screen, "Actions:")
		fmt.Fprintln(tm.Screen, "    1. [0 - 10] - Roll ball.")
		fmt.Fprintln(tm.Screen, "    2. x - Roll ball is 10 pins.")
		fmt.Fprintln(tm.Screen, "    4. q - Leave Game.")

		tm.Flush()
		key, _ := br.ReadString('\n')
		key = key[:len(key)-1]

		nbr, err := strconv.ParseInt(key, 0, 64)

		switch {
		case err == nil:
			switch {
			case int(nbr) < 0:
				fmt.Fprintf(tm.Screen, "%30s\n", "Pins Needs Positive")
			case int(nbr) > left:
				fmt.Fprintf(tm.Screen, "%27s %2d\n", "Roll too many, Pins is left", int(left))
			default:
				fmt.Fprintf(tm.Screen, "%27s %2d\n", "Roll", int(nbr))
				game.Roll(int(nbr))
			}
		case key == "q":
			exit = true
		case key == "x":
			if 10 > left {
				fmt.Fprintf(tm.Screen, "%27s %2d\n", "Roll too many, Pins is left", int(left))
			} else {
				game.Roll(10)
			}
		default:
			fmt.Fprintf(tm.Screen, "%30s\n", "Please input the Number")
		}
		tm.Flush()
	}
	fmt.Fprintln(os.Stdout, "Bowlling Game Over")

}
