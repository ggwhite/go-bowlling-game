package bowlling

import (
	"testing"
)

func TestRollOne(t *testing.T) {
	game := New()
	expected := 1

	game.Roll(1)

	for i := 0; i < 19; i++ {
		game.Roll(0)
	}

	score := game.Score()
	if expected != score {
		t.Errorf("Score %d, want %d", score, expected)
	}
}

func TestAllOne(t *testing.T) {
	game := New()
	expected := 20

	for i := 0; i < 20; i++ {
		game.Roll(1)
	}

	score := game.Score()
	if expected != score {
		t.Errorf("Score %d, want %d", score, expected)
	}
}

func TestOverRoll(t *testing.T) {
	game := New()
	expected := 0

	for i := 0; i < 20; i++ {
		game.Roll(0)
	}

	score := game.Score()
	if expected != score {
		t.Errorf("Score %d, want %d", score, expected)
	}
}

func TestSpare(t *testing.T) {
	game := New()
	expected := 14

	game.Roll(2)
	game.Roll(8)
	game.Roll(1)
	game.Roll(2)

	for i := 0; i < 16; i++ {
		game.Roll(0)
	}

	score := game.Score()
	if expected != score {
		t.Errorf("Score %d, want %d", score, expected)
	}
}

func TestSpare2(t *testing.T) {
	game := New()
	expected := 14

	game.Roll(1)
	game.Roll(9)
	game.Roll(1)
	game.Roll(2)

	for i := 0; i < 16; i++ {
		game.Roll(0)
	}

	score := game.Score()
	if expected != score {
		t.Errorf("Score %d, want %d", score, expected)
	}
}

func TestStrike(t *testing.T) {
	game := New()
	expected := 16

	game.Roll(10)
	game.Roll(1)
	game.Roll(2)

	for i := 0; i < 16; i++ {
		game.Roll(0)
	}

	score := game.Score()
	if expected != score {
		t.Errorf("Score %d, want %d", score, expected)
	}
}

func TestLastRoundSpare(t *testing.T) {
	game := New()
	expected := 11

	for i := 0; i < 18; i++ {
		game.Roll(0)
	}

	game.Roll(0)
	game.Roll(10)
	game.Roll(1)

	score := game.Score()
	if expected != score {
		t.Errorf("Score %d, want %d", score, expected)
	}
}

func TestLastRoundStrike(t *testing.T) {
	game := New()
	expected := 14

	for i := 0; i < 18; i++ {
		game.Roll(0)
	}

	game.Roll(10)
	game.Roll(3)
	game.Roll(1)

	score := game.Score()
	if expected != score {
		t.Errorf("Score %d, want %d", score, expected)
	}
}

func TestLastRoundTurkey(t *testing.T) {
	game := New()
	expected := 30

	for i := 0; i < 18; i++ {
		game.Roll(0)
	}

	game.Roll(10)
	game.Roll(10)
	game.Roll(10)

	score := game.Score()
	if expected != score {
		t.Errorf("Score %d, want %d", score, expected)
	}
}

func TestDouble(t *testing.T) {
	game := New()
	expected := 53

	game.Roll(10)
	game.Roll(10)
	game.Roll(3)
	game.Roll(7)

	for i := 0; i < 16; i++ {
		game.Roll(0)
	}

	score := game.Score()
	if expected != score {
		t.Errorf("Score %d, want %d", score, expected)
	}
}

func TestTurkey(t *testing.T) {
	game := New()
	expected := 77

	game.Roll(10)
	game.Roll(10)
	game.Roll(10)
	game.Roll(3)
	game.Roll(4)

	for i := 0; i < 15; i++ {
		game.Roll(0)
	}

	score := game.Score()
	if expected != score {
		t.Errorf("Score %d, want %d", score, expected)
	}
}

func TestTwoDouble(t *testing.T) {
	game := New()
	expected := 104

	game.Roll(10)
	game.Roll(10)
	game.Roll(3)
	game.Roll(7)

	game.Roll(10)
	game.Roll(10)
	game.Roll(3)
	game.Roll(1)

	for i := 0; i < 8; i++ {
		game.Roll(0)
	}

	score := game.Score()
	if expected != score {
		t.Errorf("Score %d, want %d", score, expected)
	}
}

func TestTwoTurkey(t *testing.T) {
	game := New()
	expected := 151

	game.Roll(10)
	game.Roll(10)
	game.Roll(10)
	game.Roll(3)
	game.Roll(4)

	game.Roll(10)
	game.Roll(10)
	game.Roll(10)
	game.Roll(2)
	game.Roll(4)

	for i := 0; i < 4; i++ {
		game.Roll(0)
	}

	score := game.Score()
	if expected != score {
		t.Errorf("Score %d, want %d", score, expected)
	}
}

func TestAllStrike(t *testing.T) {
	game := New()
	expected := 300

	for i := 0; i < 12; i++ {
		game.Roll(10)
	}

	score := game.Score()
	if expected != score {
		t.Errorf("Score %d, want %d", score, expected)
	}
}

func TestTooManyRoll(t *testing.T) {
	game := New()
	expected := 300

	for i := 0; i < 13; i++ {
		game.Roll(10)
	}

	score := game.Score()
	if expected != score {
		t.Errorf("Score %d, want %d", score, expected)
	}
}

func TestLastRoundNoBonus(t *testing.T) {
	game := New()
	expected := 2

	for i := 0; i < 18; i++ {
		game.Roll(0)
	}
	game.Roll(1)
	game.Roll(1)

	for i := 0; i < 18; i++ {
		game.Roll(1)
	}

	score := game.Score()
	if expected != score {
		t.Errorf("Score %d, want %d", score, expected)
	}
}

func TestUnfinish(t *testing.T) {
	game := New()
	expected := 7

	game.Roll(5)
	game.Roll(2)
	game.Roll(5)

	score := game.Score()
	if expected != score {
		t.Errorf("Score %d, want %d", score, expected)
	}
}

func TestLast9(t *testing.T) {
	game := New()
	expected := 299

	for i := 0; i < 11; i++ {
		game.Roll(10)
	}
	game.Roll(9)

	score := game.Score()
	if expected != score {
		t.Errorf("Score %d, want %d", score, expected)
	}
}

func TestPinsText(t *testing.T) {
	var game *Game
	var expected []string
	var res []string

	game = New()
	expected = []string{"0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "-"}
	for i := 0; i < 20; i++ {
		game.Roll(0)
	}
	res = game.PinsText()
	if len(expected) != len(res) {
		t.Errorf("Pins length %d, want %d", len(res), len(expected))
	}
	for i, _ := range expected {
		if expected[i] != res[i] {
			t.Errorf("Pin Text %s, want %s, index: %d", res[i], expected[i], i)
		}
	}

	game = New()
	expected = []string{"x", "0", "x", "0", "x", "0", "x", "0", "x", "0", "x", "0", "x", "0", "x", "0", "x", "0", "x", "x", "x"}
	for i := 0; i < 12; i++ {
		game.Roll(10)
	}
	res = game.PinsText()
	if len(expected) != len(res) {
		t.Errorf("Pins length %d, want %d", len(res), len(expected))
	}
	for i, _ := range expected {
		if expected[i] != res[i] {
			t.Errorf("Pin Text %s, want %s, index: %d", res[i], expected[i], i)
		}
	}

	game = New()
	expected = []string{"0", "/", "1", "/", "2", "/", "3", "/", "4", "/", "5", "/", "6", "/", "7", "/", "8", "/", "9", "/", "x"}
	for i := 0; i < 10; i++ {
		game.Roll(i)
		game.Roll(10 - i)
	}
	game.Roll(10)
	res = game.PinsText()
	if len(expected) != len(res) {
		t.Errorf("Pins length %d, want %d", len(res), len(expected))
	}
	for i, _ := range expected {
		if expected[i] != res[i] {
			t.Errorf("Pin Text %s, want %s, index: %d", res[i], expected[i], i)
		}
	}

	game = New()
	expected = []string{"0", "/", "x", "0", "2", "/", "0", "0", "4", "/", "x", "0", "6", "/", "0", "0", "8", "/", "3", "2", "-"}
	game.Roll(0)
	game.Roll(10)
	game.Roll(10)
	game.Roll(2)
	game.Roll(8)
	game.Roll(0)
	game.Roll(0)
	game.Roll(4)
	game.Roll(6)
	game.Roll(10)
	game.Roll(6)
	game.Roll(4)
	game.Roll(0)
	game.Roll(0)
	game.Roll(8)
	game.Roll(2)
	game.Roll(3)
	game.Roll(2)

	game.Roll(1)
	game.Roll(2)
	res = game.PinsText()
	if len(expected) != len(res) {
		t.Errorf("Pins length %d, want %d", len(res), len(expected))
	}
	for i, _ := range expected {
		if expected[i] != res[i] {
			t.Errorf("Pin Text %s, want %s, index: %d", res[i], expected[i], i)
		}
	}
}

func TestPinsScore(t *testing.T) {
	var game *Game
	var expected []string
	var res []string

	game = New()

	// not roll yet
	expected = []string{"", "", "", "", "", "", "", "", "", ""}
	res = game.ScoreText()
	if len(expected) != len(res) {
		t.Errorf("Score Text length %d, want %d", len(res), len(expected))
	}
	for i, _ := range expected {
		if expected[i] != res[i] {
			t.Errorf("Pin Text %s, want %s, index: %d", res[i], expected[i], i)
		}
	}

	game.Roll(0)
	expected = []string{"", "", "", "", "", "", "", "", "", ""}
	res = game.ScoreText()
	if len(expected) != len(res) {
		t.Errorf("Score Text length %d, want %d", len(res), len(expected))
	}
	for i, _ := range expected {
		if expected[i] != res[i] {
			t.Errorf("Pin Text %s, want %s, index: %d", res[i], expected[i], i)
		}
	}

	game.Roll(10)
	expected = []string{"", "", "", "", "", "", "", "", "", ""}
	res = game.ScoreText()
	if len(expected) != len(res) {
		t.Errorf("Score Text length %d, want %d", len(res), len(expected))
	}
	for i, _ := range expected {
		if expected[i] != res[i] {
			t.Errorf("Pin Text %s, want %s, index: %d", res[i], expected[i], i)
		}
	}

	game.Roll(1)
	expected = []string{"11", "", "", "", "", "", "", "", "", ""}
	res = game.ScoreText()
	if len(expected) != len(res) {
		t.Errorf("Score Text length %d, want %d", len(res), len(expected))
	}
	for i, _ := range expected {
		if expected[i] != res[i] {
			t.Errorf("Pin Text %s, want %s, index: %d", res[i], expected[i], i)
		}
	}

	game.Roll(2)
	expected = []string{"11", "14", "", "", "", "", "", "", "", ""}
	res = game.ScoreText()
	if len(expected) != len(res) {
		t.Errorf("Score Text length %d, want %d", len(res), len(expected))
	}
	for i, _ := range expected {
		if expected[i] != res[i] {
			t.Errorf("Pin Text %s, want %s, index: %d", res[i], expected[i], i)
		}
	}

	game.Roll(9)
	expected = []string{"11", "14", "", "", "", "", "", "", "", ""}
	res = game.ScoreText()
	if len(expected) != len(res) {
		t.Errorf("Score Text length %d, want %d", len(res), len(expected))
	}
	for i, _ := range expected {
		if expected[i] != res[i] {
			t.Errorf("Pin Text %s, want %s, index: %d", res[i], expected[i], i)
		}
	}

	game.Roll(1)
	expected = []string{"11", "14", "", "", "", "", "", "", "", ""}
	res = game.ScoreText()
	if len(expected) != len(res) {
		t.Errorf("Score Text length %d, want %d", len(res), len(expected))
	}
	for i, _ := range expected {
		if expected[i] != res[i] {
			t.Errorf("Pin Text %s, want %s, index: %d", res[i], expected[i], i)
		}
	}

	game.Roll(10)
	expected = []string{"11", "14", "34", "", "", "", "", "", "", ""}
	res = game.ScoreText()
	if len(expected) != len(res) {
		t.Errorf("Score Text length %d, want %d", len(res), len(expected))
	}
	for i, _ := range expected {
		if expected[i] != res[i] {
			t.Errorf("Pin Text %s, want %s, index: %d", res[i], expected[i], i)
		}
	}

	game.Roll(0)
	expected = []string{"11", "14", "34", "", "", "", "", "", "", ""}
	res = game.ScoreText()
	if len(expected) != len(res) {
		t.Errorf("Score Text length %d, want %d", len(res), len(expected))
	}
	for i, _ := range expected {
		if expected[i] != res[i] {
			t.Errorf("Pin Text %s, want %s, index: %d", res[i], expected[i], i)
		}
	}

	game.Roll(0)
	expected = []string{"11", "14", "34", "44", "44", "", "", "", "", ""}
	res = game.ScoreText()
	if len(expected) != len(res) {
		t.Errorf("Score Text length %d, want %d", len(res), len(expected))
	}
	for i, _ := range expected {
		if expected[i] != res[i] {
			t.Errorf("Pin Text %s, want %s, index: %d", res[i], expected[i], i)
		}
	}

	game.Roll(8)
	expected = []string{"11", "14", "34", "44", "44", "", "", "", "", ""}
	res = game.ScoreText()
	if len(expected) != len(res) {
		t.Errorf("Score Text length %d, want %d", len(res), len(expected))
	}
	for i, _ := range expected {
		if expected[i] != res[i] {
			t.Errorf("Pin Text %s, want %s, index: %d", res[i], expected[i], i)
		}
	}

	game.Roll(2)
	expected = []string{"11", "14", "34", "44", "44", "", "", "", "", ""}
	res = game.ScoreText()
	if len(expected) != len(res) {
		t.Errorf("Score Text length %d, want %d", len(res), len(expected))
	}
	for i, _ := range expected {
		if expected[i] != res[i] {
			t.Errorf("Pin Text %s, want %s, index: %d", res[i], expected[i], i)
		}
	}

	game.Roll(0)
	expected = []string{"11", "14", "34", "44", "44", "54", "", "", "", ""}
	res = game.ScoreText()
	if len(expected) != len(res) {
		t.Errorf("Score Text length %d, want %d", len(res), len(expected))
	}
	for i, _ := range expected {
		if expected[i] != res[i] {
			t.Errorf("Pin Text %s, want %s, index: %d", res[i], expected[i], i)
		}
	}

	game.Roll(10)
	expected = []string{"11", "14", "34", "44", "44", "54", "", "", "", ""}
	res = game.ScoreText()
	if len(expected) != len(res) {
		t.Errorf("Score Text length %d, want %d", len(res), len(expected))
	}
	for i, _ := range expected {
		if expected[i] != res[i] {
			t.Errorf("Pin Text %s, want %s, index: %d", res[i], expected[i], i)
		}
	}

	game.Roll(10)
	expected = []string{"11", "14", "34", "44", "44", "54", "74", "", "", ""}
	res = game.ScoreText()
	if len(expected) != len(res) {
		t.Errorf("Score Text length %d, want %d", len(res), len(expected))
	}
	for i, _ := range expected {
		if expected[i] != res[i] {
			t.Errorf("Pin Text %s, want %s, index: %d", res[i], expected[i], i)
		}
	}

	game.Roll(10)
	expected = []string{"11", "14", "34", "44", "44", "54", "74", "", "", ""}
	res = game.ScoreText()
	if len(expected) != len(res) {
		t.Errorf("Score Text length %d, want %d", len(res), len(expected))
	}
	for i, _ := range expected {
		if expected[i] != res[i] {
			t.Errorf("Pin Text %s, want %s, index: %d", res[i], expected[i], i)
		}
	}

	game.Roll(10)
	expected = []string{"11", "14", "34", "44", "44", "54", "74", "104", "", ""}
	res = game.ScoreText()
	if len(expected) != len(res) {
		t.Errorf("Score Text length %d, want %d", len(res), len(expected))
	}
	for i, _ := range expected {
		if expected[i] != res[i] {
			t.Errorf("Pin Text %s, want %s, index: %d", res[i], expected[i], i)
		}
	}

	game.Roll(0)
	expected = []string{"11", "14", "34", "44", "44", "54", "74", "104", "124", "134"}
	res = game.ScoreText()
	if len(expected) != len(res) {
		t.Errorf("Score Text length %d, want %d", len(res), len(expected))
	}
	for i, _ := range expected {
		if expected[i] != res[i] {
			t.Errorf("Pin Text %s, want %s, index: %d", res[i], expected[i], i)
		}
	}

	game.Roll(6)
	expected = []string{"11", "14", "34", "44", "44", "54", "74", "104", "124", "140"}
	res = game.ScoreText()
	if len(expected) != len(res) {
		t.Errorf("Score Text length %d, want %d", len(res), len(expected))
	}
	for i, _ := range expected {
		if expected[i] != res[i] {
			t.Errorf("Pin Text %s, want %s, index: %d", res[i], expected[i], i)
		}
	}

}
