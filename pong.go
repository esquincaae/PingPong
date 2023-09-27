package main

import (
	"pong/scenes"
/*	"fmt"
	"image"
	"image/color"
	"sync"
	"time"
	"os"
*/
	//"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
/*	"github.com/faiface/pixel/text"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font/basicfont"
	
*/)


/* modelos
func updatePlayer(win *pixelgl.Window) {
	for {
		if win.Pressed(pixelgl.KeyW) && playerPaddleRect.Max.Y < height {
			playerPos += 10.0
			playerPaddleRect = playerPaddleRect.Moved(pixel.V(0, float64(paddleSpeed)))
		}
		if win.Pressed(pixelgl.KeyS) && playerPaddleRect.Min.Y > 0 {
			playerPos -= 10.0
			playerPaddleRect = playerPaddleRect.Moved(pixel.V(0, -float64(paddleSpeed)))
		}
		time.Sleep(time.Millisecond * 16)
	}
}

func updateplayer2(win *pixelgl.Window) {
	for {
		if win.Pressed(pixelgl.KeyUp) && player2PaddleRect.Max.Y < height {
			player2Pos += 10.0
			player2PaddleRect = player2PaddleRect.Moved(pixel.V(0, float64(paddleSpeed)))
		}
		if win.Pressed(pixelgl.KeyDown) && player2PaddleRect.Min.Y > 0 {
			player2Pos -= 10.0
			player2PaddleRect = player2PaddleRect.Moved(pixel.V(0, -float64(paddleSpeed)))
		}
		time.Sleep(time.Millisecond * 16)
	}
}

func updateBall() {
	for {
		ballPos = ballPos.Add(ballDir.Scaled(ballSpeed))

		// Check collision with the top and bottom of the screen
		if ballPos.Y-ballSize < 0 || ballPos.Y+ballSize > height {
			ballDir.Y = -ballDir.Y
		}

		// Collision with player paddle
		if ballPos.X-ballSize < playerPaddleRect.Max.X &&
			ballPos.Y > playerPaddleRect.Min.Y &&
			ballPos.Y < playerPaddleRect.Max.Y {
			ballDir.X = -ballDir.X
		}

		// Collision with player2 paddle
		if ballPos.X+ballSize > player2PaddleRect.Min.X &&
			ballPos.Y > player2PaddleRect.Min.Y &&
			ballPos.Y < player2PaddleRect.Max.Y {
			ballDir.X = -ballDir.X
		}

		if ballPos.X-ballSize < 0 {
			player2Score++
			ballPos = pixel.V(width/2, height/2)
			ballDir = pixel.V(1, 1).Unit()
		} else if ballPos.X+ballSize > width {
			playerScore++
			ballPos = pixel.V(width/2, height/2)
			ballDir = pixel.V(-1, 1).Unit()
		} 

		time.Sleep(time.Millisecond * 16)
	}
}
*/


/* ------  scenes --------
func run() {

	win, err := pixelgl.NewWindow(pixelgl.WindowConfig{
		Title:  "Ping Pong",
		Bounds: pixel.R(0, 0, width, height),
		VSync:  true,
	})
	if err != nil {
		panic(err)
	}

	// Crea y dibuja el fondo
	bgImg := image.NewRGBA(image.Rect(0, 0, int(width), int(height)))
	blue := color.RGBA{70, 70, 255, 255}
	for y := bgImg.Rect.Min.Y; y < bgImg.Rect.Max.Y; y++ {
		for x := bgImg.Rect.Min.X; x < bgImg.Rect.Max.X; x++ {
			bgImg.Set(x, y, blue)
		}
	}

	// Dibuja la red
	white := color.RGBA{255, 255, 255, 255}
	for x := int(width/2) - 1; x <= int(width/2) + 1; x++ {
		for y := bgImg.Rect.Min.Y; y < bgImg.Rect.Max.Y; y++ {
			bgImg.Set(x, y, white)
		}
	}
	bgPic := pixel.PictureDataFromImage(bgImg)
	bgSprite := pixel.NewSprite(bgPic, bgPic.Bounds())

	playerScorePos := pixel.V(width/4, height-40)
	player2ScorePos := pixel.V(3*width/4, height-40)
	basicAtlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	txt := text.New(pixel.V(100, 500), basicAtlas)

	paddleImg := image.NewRGBA(image.Rect(0, 0, int(paddleWidth), int(paddleHeight)))
	col := color.RGBA{255, 255, 255, 255}
	for y := paddleImg.Rect.Min.Y; y < paddleImg.Rect.Max.Y; y++ {
		for x := paddleImg.Rect.Min.X; x < paddleImg.Rect.Max.X; x++ {
			paddleImg.Set(x, y, col)
		}
	}
	paddlePic := pixel.PictureDataFromImage(paddleImg)
	paddleSprite := pixel.NewSprite(paddlePic, paddlePic.Bounds())

	ballImg := image.NewRGBA(image.Rect(0, 0, int(ballSize*2), int(ballSize*2)))
	
	center := float64(int(ballSize)) // coordenadas del centro del círculo
	for y := ballImg.Rect.Min.Y; y < ballImg.Rect.Max.Y; y++ {
		for x := ballImg.Rect.Min.X; x < ballImg.Rect.Max.X; x++ {
			dx := float64(x) - center
			dy := float64(y) - center
			if dx*dx+dy*dy <= center*center { // Fórmula del círculo
				ballImg.Set(x, y, col)
			}
		}
	}

	ballPic := pixel.PictureDataFromImage(ballImg)
	ballSprite := pixel.NewSprite(ballPic, ballPic.Bounds())

	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		updatePlayer(win)
		wg.Done()
	}()
	go func() {
		updateplayer2(win)
		wg.Done()
	}()
	go func() {
		updateBall()
		wg.Done()
	}()

	for !win.Closed() {
		win.Clear(colornames.Black)

		bgSprite.Draw(win, pixel.IM.Moved(win.Bounds().Center())) // borrar

		paddleSprite.Draw(win, pixel.IM.Moved(playerPaddleRect.Center()))
		paddleSprite.Draw(win, pixel.IM.Moved(player2PaddleRect.Center()))
		ballSprite.Draw(win, pixel.IM.Moved(ballPos))


		if playerScore >= 5 || player2Score >= 5 {
            win.Clear(colornames.Black)

            winnerTxt := text.New(pixel.V(width/2-200, height/2), text.NewAtlas(basicfont.Face7x13, text.ASCII))
            if playerScore >= 5 {
                fmt.Fprintf(winnerTxt, "Ganador: Jugador 1")
            } else {
                fmt.Fprintf(winnerTxt, "Ganador: Jugador 2")
            }

            winnerTxt.Draw(win, pixel.IM.Scaled(winnerTxt.Orig, 3))
            win.Update()

            // Mantener el mensaje en pantalla durante 5 segundos
            time.Sleep(5 * time.Second)
            
            // Cierra la ventana y sale del programa
            win.Destroy()
            os.Exit(0)
        }

		// Render scores
		txt.Clear()
		txt.Dot = playerScorePos
		fmt.Fprintf(txt, "%d", playerScore)
		txt.Dot = player2ScorePos
		fmt.Fprintf(txt, "%d", player2Score)
		txt.Draw(win, pixel.IM)

		win.Update()
	}
}
------- scenes -------*/

func main() {
	pixelgl.Run(scenes.Run)
}
