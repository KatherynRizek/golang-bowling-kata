package main

type Game struct {
	rolls [21]int
	currentRoll int
}

func (game *Game) GetScore() int {
	score := 0
	rollCounter := 0
	for frameCounter := 0; frameCounter < 10; frameCounter += 1 {
		if isStrike(game.rolls[rollCounter]) {
			score += 10 + game.rolls[rollCounter + 1] + game.rolls[rollCounter + 2]
			rollCounter += 1
		} else if isSpare(game.rolls[rollCounter], game.rolls[rollCounter + 1]) {
			score += 10 + game.rolls[rollCounter + 2]
			rollCounter += 2
		} else {
			score += game.rolls[rollCounter] + game.rolls[rollCounter + 1]
			rollCounter += 2
		}
	}
	return score
}

func (game *Game) Roll(pins int) {
	game.rolls[game.currentRoll] = pins
	game.currentRoll++
}

func isSpare(throw1, throw2 int) bool {
	status:= false
	frameScore:= throw1 + throw2
	if frameScore == 10 {
		status = true
	}
	return status
}

func isStrike(throw1 int) bool {
	status:= false
	if throw1 == 10 {
		status = true
	}
	return status
}