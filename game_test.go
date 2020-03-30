package bowlinggame

import (
	"testing"
)

func TestGutterGame(t *testing.T) {
	g := NewGame()
	rollNTimesWithPins(g, 20, 0)
	checkGameScore(g, t, 0)
}

func TestAllOnes(t *testing.T) {
	g := NewGame()
	rollNTimesWithPins(g, 20, 1)
	checkGameScore(g, t, 20)
}

func TestOneSpare(t *testing.T) {
	g := NewGame()
	rollSpare(g) // 11
	rollFrame(g, 1, 0)
	rollNTimesWithPins(g, 16, 0)
	checkGameScore(g, t, 12)
}

func TestThreeSpares(t *testing.T) {
	g := NewGame()
	rollFrame(g, 3, 3)
	rollSpare(g) //15
	rollSpare(g) //15
	rollSpare(g) //11
	rollFrame(g, 1, 1)
	rollNTimesWithPins(g, 10, 0)
	checkGameScore(g, t, 49)
}

func TestOneStrike(t *testing.T) {
	g := NewGame()
	rollStrike(g)
	rollFrame(g, 3, 3)
	rollNTimesWithPins(g, 16, 0)
	checkGameScore(g, t, 22)
}

func TestThreeStrikes(t *testing.T) {
	g := NewGame()
	rollFrame(g, 1, 1)
	rollStrike(g) //30
	rollStrike(g) //23
	rollStrike(g) //16
	rollFrame(g, 3, 3)
	rollNTimesWithPins(g, 10, 0)
	checkGameScore(g, t, 77)
}

func TestSpareThenStrike(t *testing.T) {
	g := NewGame()
	rollFrame(g, 1, 1)
	rollSpare(g)  //20
	rollStrike(g) //16
	rollFrame(g, 3, 3)
	rollNTimesWithPins(g, 12, 0)
	checkGameScore(g, t, 44)
}

func TestStrikeThenSpare(t *testing.T) {
	g := NewGame()
	rollFrame(g, 1, 1)
	rollStrike(g) //20
	rollSpare(g)  //13
	rollFrame(g, 3, 3)
	rollNTimesWithPins(g, 12, 0)
	checkGameScore(g, t, 41)
}

func TestPerfectGame(t *testing.T) {
	g := NewGame()
	rollNTimesWithPins(g, 12, 10)
	checkGameScore(g, t, 300)
}

func rollFrame(g *Game, num, num2 int) {
	g.Roll(num)
	g.Roll(num2)
}

func rollStrike(g *Game) {
	g.Roll(10)
}

func rollSpare(g *Game) {
	g.Roll(5)
	g.Roll(5)
}

func checkGameScore(game *Game, t *testing.T, expectedScore int) {
	if game.Score() != expectedScore {
		t.Errorf("expected score to be %v: got %v", expectedScore, game.Score())
	}
}

func NewGame() *Game {
	frames := make([]*Frame, 12)
	for i, _ := range frames {
		frames[i] = &Frame{whichRoll: FirstRoll}
	}
	return &Game{
		currentFrameIndex: 0,
		frames:            frames,
	}
}

func rollNTimesWithPins(game *Game, frames int, pins int) {
	for i := 0; i < frames; i++ {
		game.Roll(pins)
	}
}
