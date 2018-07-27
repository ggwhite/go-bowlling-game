package bowlling

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

type Game struct {
	Pins []int
}

func New() *Game {
	return &Game{make([]int, 0, 21)}
}

func (game *Game) ShowScore(w io.Writer) {
	buf := bufio.NewWriter(w)
	buf.Reset(w)

	// show thead
	for i := 0; i < 9; i++ {
		fmt.Fprintf(buf, "| %5s %d ", "Round", i+1)
	}
	fmt.Fprintf(buf, "| %7s %d   |\n", "Round", 10)

	// show pins
	pins := game.PinsText()
	for _, pin := range pins {
		fmt.Fprintf(buf, "|  %s ", pin)
	}
	fmt.Fprintln(buf, "|")

	// show score
	score := game.ScoreText()
	for i := 0; i < 9; i++ {
		fmt.Fprintf(buf, "|%3s%3s%3s", "", score[i], "")
	}
	fmt.Fprintf(buf, "|%5s%4s%5s|\n", "", score[9], "")

	defer buf.Flush()
}

func (game *Game) Over() bool {
	if len(game.Pins) == 21 {
		return true
	}
	if len(game.Pins) == 20 && (game.Pins[18]+game.Pins[19]) < 10 {
		return true
	}
	return false
}

func (game *Game) Roll(pins int) {
	if game.Over() {
		return
	}
	game.Pins = append(game.Pins, pins)
	if pins == 10 && len(game.Pins)%2 == 1 && len(game.Pins) <= 18 {
		game.Pins = append(game.Pins, 0)
	}
}

func (game *Game) LeftPins() int {
	left := 10

	if len(game.Pins)%2 == 0 && len(game.Pins) < 19 {
		return left
	}

	if len(game.Pins) == 19 {
		if game.Pins[18] == 10 {
			return left
		}
	}

	if len(game.Pins) == 20 {
		if game.Pins[19] == 10 {
			return left
		}

		if game.Pins[19]+game.Pins[18] == 10 {
			return left
		}
	}

	left -= game.Pins[len(game.Pins)-1]

	return left
}

func (game *Game) Score() int {
	score := 0
	point := len(game.Pins) - 1

	round := int((point + 1) / 2)
	score = game.RoundScore(round)
	if score < 0 {
		score = game.RoundScore(round - 1)
	}

	return score
}

func (game *Game) RoundScore(round int) int {
	score := 0

	if round < 1 || round > 10 {
		return -99
	}

	point := len(game.Pins) - 1

	if round*2-1 > point {
		return -1
	}

	for idx := 0; idx < round*2; idx++ {
		if idx > point {
			continue
		}
		pin := game.Pins[idx]
		score += pin
		if idx%2 == 1 {
			last := game.Pins[idx-1]
			if 10 == pin+last && idx < 19 {
				if point > idx {
					next := game.Pins[idx+1]
					score += next
					if pin == 0 {
						if next == 10 && idx+2 < 19 {
							if point > idx+2 {
								score += game.Pins[idx+3]
							} else {
								return -1
							}
						} else {
							if point > idx+1 {
								score += game.Pins[idx+2]
							} else {
								return -1
							}
						}
					}
				} else {
					return -1
				}
			}
		}
	}
	if point == 20 && round == 10 {
		score += game.Pins[20]
	}

	return score
}

func (game *Game) PinsText() []string {
	pins := []string{}

	for i := 0; i < 21; i++ {
		if len(game.Pins) > i {
			if i == 19 || i == 20 {
				if game.Pins[i-1] == 10 && game.Pins[i] == 10 {
					pins = append(pins, "x")
					continue
				}
				if game.Pins[i-1] == 0 && game.Pins[i] == 10 {
					pins = append(pins, "/")
					continue
				}
			}
			if i%2 == 0 && game.Pins[i] == 10 {
				pins = append(pins, "x")
			} else if i%2 == 1 && game.Pins[i]+game.Pins[i-1] == 10 && game.Pins[i-1] != 10 {
				pins = append(pins, "/")
			} else {
				pins = append(pins, strconv.Itoa(game.Pins[i]))
			}
		} else {
			pins = append(pins, "-")
		}
	}

	return pins
}

func (game *Game) ScoreText() []string {
	score := []string{}

	for i := 0; i < 10; i++ {
		s := game.RoundScore(i + 1)
		if s > 0 {
			score = append(score, strconv.Itoa(s))
		} else {
			score = append(score, "")
		}
	}

	return score
}
