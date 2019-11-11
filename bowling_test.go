package main

import "testing"

func TestGutterGame(t *testing.T) {
	game := Game{}
	rollMany(20, 0, &game)
	score := game.GetScore()
	checkResult(score, 0, t)
}

func TestRollAllOnes(t *testing.T) {
	game := Game{}
	rollMany(20, 1, &game)
	score := game.GetScore()
	checkResult(score, 20, t)
}

func TestOneSpare(t *testing.T) {
	game := Game{}
	rollSpare(&game)
	game.Roll(3)
	score := game.GetScore()
	checkResult(score, 16, t)
}

func TestOneStrike(t *testing.T) {
	game := Game{}
	rollStrike(&game)
	game.Roll(4)
	game.Roll(5)
	score := game.GetScore()
	checkResult(score, 28, t)
}

func TestPerfectGame( t *testing.T) {
	game := Game{}
	rollMany(12, 10, &game)
	score := game.GetScore()
	checkResult(score, 300, t)
}

func rollMany(numTimes int, pins int, game *Game) {
	for i := 0; i < numTimes; i++ {
		game.Roll(pins)
	}
}

func rollSpare(game *Game) {
	game.Roll(5)
	game.Roll(5)
}

func rollStrike(game *Game) {
	game.Roll(10)
}

func checkResult(actual int, expected int, t *testing.T) {
	if actual != expected {
		t.Fail()
		t.Logf("Expected %d, got %d", expected, actual)
	}
}