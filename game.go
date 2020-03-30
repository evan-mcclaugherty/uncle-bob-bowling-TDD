package bowlinggame

const (
	FirstRoll = iota
	SecondRoll
)

type Game struct {
	currentFrameIndex int
	frames            []*Frame
}

type Frame struct {
	whichRoll   WhichRoll
	firstScore  int
	secondScore int
	isStrike    bool
	isSpare     bool
}

type WhichRoll int

func (g *Game) Roll(pins int) {
	currentFrame := g.currentFrame()

	if currentFrame.whichRoll == FirstRoll {
		currentFrame.firstScore = pins
		if pins == 10 {
			currentFrame.isStrike = true
			g.currentFrameIndex++
		}
		currentFrame.whichRoll = SecondRoll
	} else {
		currentFrame.secondScore = pins
		if currentFrame.firstScore+currentFrame.secondScore == 10 {
			currentFrame.isSpare = true
		}
		g.currentFrameIndex++
	}
}

func (g *Game) currentFrame() *Frame {
	return g.frames[g.currentFrameIndex]
}

func (g *Game) Score() int {
	score := 0
	for indexOfFrame, frame := range g.frames {
		if indexOfFrame == 10 {
			break
		}
		if frame.isSpare {
			score += g.calculateSpare(indexOfFrame)
		} else if frame.isStrike {
			score += g.calculateStrike(indexOfFrame)
		} else {
			score += frame.firstScore + frame.secondScore
		}
	}
	return score
}

func (g *Game) calculateSpare(indexOfFrame int) int {
	return 10 + g.nextFrame(indexOfFrame).firstScore
}

func (g *Game) nextFrame(indexOfFrame int) *Frame {
	return g.frames[indexOfFrame+1]
}

func (g *Game) calculateStrike(indexOfFrame int) int {
	score := 10
	nextFrame := g.nextFrame(indexOfFrame)
	if nextFrame.isStrike {
		score += 10
		nextFrame = g.nextFrame(indexOfFrame + 1)
		if nextFrame.isStrike {
			score += 10
		} else {
			score += nextFrame.firstScore
		}
	} else {
		score += nextFrame.firstScore + nextFrame.secondScore
	}
	return score
}
