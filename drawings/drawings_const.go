/*
	@author: Robert Pratt <robert.pratt[at]homelab.farm>
	@package: drawings
	@description: Struct definitions for /api/v1/draw-results and children.
	@created: 2020-12-05
*/
package drawings

import "time"

const drawResultsUrlPath = "/api/v1/draw-results"

type Drawing struct {
	EstimatedJackpot []GameJackpot `json:"estimatedJackpot"`
	WinningNumbers   []GameDrawing `json:"winningNumbers"`
}

type GameJackpot struct {
	DrawDateFor            time.Time `json:"drawDateFor"`
	DrawDateFrom           time.Time `json:"drawDateFrom"`
	DrawNumberFor          int32     `json:"drawNumberFor"`
	DrawNumberFrom         int32     `json:"drawNumberFrom"`
	EstimatedCashOptionUSD int64     `json:"estimatedCashOptionUSD"`
	EstimatedJackpotUSD    int64     `json:"estimatedJackpotUSD"`
	GameIdentifier         string    `json:"gameIdentifier"`
	Status                 string    `json:"status"`
}

type GameDrawing struct {
	DrawDate       time.Time `json:"drawDate"`
	DrawNumber     int32     `json:"drawNumber"`
	DrawSequence   int32     `json:"drawSequence"`
	Extras         Extra     `json:"extras"`
	GameIdentifier Game      `json:"gameIdentifier"`
	Location       string    `json:"location"`
	Status         string    `json:"status"`
	VideoLink      string    `json:"videoLink"`
	Winners        int32     `json:"winners"`
	WinningNumbers []int32   `json:"winningNumbers"`
}

type Game string

const (
	POWERBALL         Game = "powerball"
	MEGABUCKS_DOUBLER Game = "megabucks_doubler"
	MASS_CASH         Game = "mass_cash"
	MEGA_MILLIONS     Game = "mega_millions"
	LUCKY_FOR_LIFE    Game = "lucky_for_life"
	THE_NUMBERS_GAME  Game = "the_numbers_game"
)

type Extra struct {
	PowerBall int32 `json:"powerball"`
	PowerPlay int32 `json:"powerplay"`
	StDoubler int32 `json:"stDoubler"`
	MegaBall  int32 `json:"megaball"`
	MegaPlier int32 `json:"megaplier"`
	LuckyBall int32 `json:"luckyball"`
}
