// Package vanilla is a good example... source: https://smartystreets.com/blog/2015/02/go-testing-part-1-vanillla
package vanilla

// Game contains the state of a bowling game.
type Game struct {
	rolls   []int
	current int
}

// NewGame allocates and starts a new game of bowling.
func NewGame() *Game {
	game := new(Game)
	game.rolls = make([]int, maxThrowsPerGame)
	return game
}

// Roll rolls the ball and knocks down the number of pins specified by pins.
func (selfg *Game) Roll(pins int) {
	selfg.rolls[selfg.current] = pins
	selfg.current++
}

// Score calculates and returns the player's current score.
func (selfg *Game) Score() (sum int) {
	for throw, frame := 0, 0; frame < framesPerGame; frame++ {
		if selfg.isStrike(throw) {
			sum += selfg.strikeBonusFor(throw)
			throw++
		} else if selfg.isSpare(throw) {
			sum += selfg.spareBonusFor(throw)
			throw += 2
		} else {
			sum += selfg.framePointsAt(throw)
			throw += 2
		}
	}
	return sum
}

// isStrike determines if a given throw is a strike or not. A strike is knocking
// down all pins in one throw.
func (selfg *Game) isStrike(throw int) bool {
	return selfg.rolls[throw] == allPins
}

// strikeBonusFor calculates and returns the strike bonus for a throw.
func (selfg *Game) strikeBonusFor(throw int) int {
	return allPins + selfg.framePointsAt(throw+1)
}

// isSpare determines if a given frame is a spare or not. A spare is knocking
// down all pins in one frame with two throws.
func (selfg *Game) isSpare(throw int) bool {
	return selfg.framePointsAt(throw) == allPins
}

// spareBonusFor calculates and returns the spare bonus for a throw.
func (selfg *Game) spareBonusFor(throw int) int {
	return allPins + selfg.rolls[throw+2]
}

// framePointsAt computes and returns the score in a frame specified by throw.
func (selfg *Game) framePointsAt(throw int) int {
	return selfg.rolls[throw] + selfg.rolls[throw+1]
}

// testing utilities:

func (selfg *Game) rollMany(times, pins int) {
	for x := 0; x < times; x++ {
		selfg.Roll(pins)
	}
}
func (selfg *Game) rollSpare() {
	selfg.Roll(5)
	selfg.Roll(5)
}
func (selfg *Game) rollStrike() {
	selfg.Roll(10)
}

const (
	// allPins is the number of pins allocated per fresh throw.
	allPins = 10

	// framesPerGame is the number of frames per bowling game.
	framesPerGame = 10

	// maxThrowsPerGame is the maximum number of throws possible in a single game.
	maxThrowsPerGame = 21
)
