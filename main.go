package main

import (
	"fmt"
	"math/rand"
	"time"
)

const PROBABILITY_GOAL = 0.0001
const PROBABILITY_FIRST_TEAM_GOAL = 0.55
const STAMPS_NUMBER = 5000

type Score struct {
	Home int
	Away int
}
type ScoreStamp struct {
	Offset int
	Score  Score
}

func main() {
	fmt.Println(getScore(fillScores(), 5000))
}
func fillScores() *[]ScoreStamp {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))

	scores := make([]ScoreStamp, 0, 10)

	for i := 0; i < STAMPS_NUMBER; i++ {
		scoreChanged := random.Float32() < PROBABILITY_GOAL
		home := 0
		away := 0
		if scoreChanged {
			if random.Float32() < PROBABILITY_FIRST_TEAM_GOAL {
				home = 1
				away = 0
			} else {
				home = 0
				away = 1
			}
		}
		var prevScore Score
		if len(scores) == 0 {
			prevScore = Score{
				Home: 0,
				Away: 0,
			}
		} else {
			prevScore = scores[i-1].Score
		}
		newScore := Score{
			Home: prevScore.Home + home,
			Away: prevScore.Away + away,
		}
		scores = append(scores, ScoreStamp{Offset: i, Score: newScore})
	}
	return &scores
}

func getScore(scores *[]ScoreStamp, offset int) Score {
	return (*scores)[offset-1].Score
}
