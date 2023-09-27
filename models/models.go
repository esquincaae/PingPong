package models

import (
		"time"
	
		"github.com/faiface/pixel"
		"github.com/faiface/pixel/pixelgl"
		"pong/views"
	)

	
func UpdatePlayer(win *pixelgl.Window) {
	for {
		if win.Pressed(pixelgl.KeyW) && views.PlayerPaddleRect.Max.Y < views.Height {
			views.PlayerPos += 10.0
			views.PlayerPaddleRect = views.PlayerPaddleRect.Moved(pixel.V(0, float64(views.PaddleSpeed)))
		}
		if win.Pressed(pixelgl.KeyS) && views.PlayerPaddleRect.Min.Y > 0 {
			views.PlayerPos -= 10.0
			views.PlayerPaddleRect = views.PlayerPaddleRect.Moved(pixel.V(0, -float64(views.PaddleSpeed)))
		}
		time.Sleep(time.Millisecond * 16)
	}
}

func Updateplayer2(win *pixelgl.Window) {
	for {
		if win.Pressed(pixelgl.KeyUp) && views.Player2PaddleRect.Max.Y < views.Height {
			views.Player2Pos += 10.0
			views.Player2PaddleRect = views.Player2PaddleRect.Moved(pixel.V(0, float64(views.PaddleSpeed)))
		}
		if win.Pressed(pixelgl.KeyDown) && views.Player2PaddleRect.Min.Y > 0 {
			views.Player2Pos -= 10.0
			views.Player2PaddleRect = views.Player2PaddleRect.Moved(pixel.V(0, -float64(views.PaddleSpeed)))
		}
		time.Sleep(time.Millisecond * 16)
	}
}

func UpdateBall() {
	for {
		views.BallPos = views.BallPos.Add(views.BallDir.Scaled(views.BallSpeed))

		// Check collision with the top and bottom of the screen
		if views.BallPos.Y-views.BallSize < 0 || views.BallPos.Y+views.BallSize > views.Height {
			views.BallDir.Y = -views.BallDir.Y
		}

		// Collision with player paddle
		if views.BallPos.X-views.BallSize < views.PlayerPaddleRect.Max.X &&
			views.BallPos.Y > views.PlayerPaddleRect.Min.Y &&
			views.BallPos.Y < views.PlayerPaddleRect.Max.Y {
				views.BallDir.X = -views.BallDir.X
		}

		// Collision with player2 paddle
		if views.BallPos.X+views.BallSize > views.Player2PaddleRect.Min.X &&
			views.BallPos.Y > views.Player2PaddleRect.Min.Y &&
			views.BallPos.Y < views.Player2PaddleRect.Max.Y {
			views.BallDir.X = -views.BallDir.X
		}

		if views.BallPos.X-views.BallSize < 0 {
			views.Player2Score++
			views.BallPos = pixel.V(views.Width/2, views.Height/2)
			views.BallDir = pixel.V(1, 1).Unit()
		} else if views.BallPos.X+views.BallSize > views.Width {
			views.PlayerScore++
			views.BallPos = pixel.V(views.Width/2, views.Height/2)
			views.BallDir = pixel.V(-1, 1).Unit()
		} 

		time.Sleep(time.Millisecond * 16)
	}
}