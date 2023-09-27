package views

import (

		"github.com/faiface/pixel"
		
	)

const (
	Width, Height             		= 800.0, 600.0
	PaddleWidth, PaddleHeight 		= 20.0, 120.0
	BallSize                  		= 20.0
)

var (
	PaddleSpeed                     = 10
	BallSpeed                		= 10.0 
	PlayerPos, Player2Pos         	= Height / 2, Height / 2
	BallPos                  		= pixel.V(Width/2, Height/2)
	BallDir                  		= pixel.V(1, 1).Unit()
	PlayerPaddleRect         		= pixel.R(30, PlayerPos-PaddleHeight/2, 30+PaddleWidth, PlayerPos+PaddleHeight/2)
	Player2PaddleRect             	= pixel.R(Width-30-PaddleWidth, Player2Pos-PaddleHeight/2, Width-30, Player2Pos+PaddleHeight/2)
	PlayerScore, Player2Score     	int
)