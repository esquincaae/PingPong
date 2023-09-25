package main

import (
	"fmt"
	"image"
	"image/color"
	"sync"
	"time"
	"os"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font/basicfont"
)

const (
	width, height             = 800.0, 600.0
	paddleWidth, paddleHeight = 15.0, 60.0
	ballSize                  = 15.0
)

var (
	ballSpeed                = 10.0 
	playerPos, aiPos         = height / 2, height / 2
	ballPos                  = pixel.V(width/2, height/2)
	ballDir                  = pixel.V(1, 1).Unit()
	playerPaddleRect         = pixel.R(30, playerPos-paddleHeight/2, 30+paddleWidth, playerPos+paddleHeight/2)
	aiPaddleRect             = pixel.R(width-30-paddleWidth, aiPos-paddleHeight/2, width-30, aiPos+paddleHeight/2)
	playerScore, aiScore     int
)

func updatePlayer(win *pixelgl.Window) {
	for {
		if win.Pressed(pixelgl.KeyW) && playerPaddleRect.Max.Y < height {
			playerPos += 10.0
			playerPaddleRect = playerPaddleRect.Moved(pixel.V(0, 10))
		}
		if win.Pressed(pixelgl.KeyS) && playerPaddleRect.Min.Y > 0 {
			playerPos -= 10.0
			playerPaddleRect = playerPaddleRect.Moved(pixel.V(0, -10))
		}
		time.Sleep(time.Millisecond * 16)
	}
}

func updateAI(win *pixelgl.Window) {
	for {
		if win.Pressed(pixelgl.KeyUp) && aiPaddleRect.Max.Y < height {
			aiPos += 10.0
			aiPaddleRect = aiPaddleRect.Moved(pixel.V(0, 10))
		}
		if win.Pressed(pixelgl.KeyDown) && aiPaddleRect.Min.Y > 0 {
			aiPos -= 10.0
			aiPaddleRect = aiPaddleRect.Moved(pixel.V(0, -10))
		}
		time.Sleep(time.Millisecond * 16)
	}
}

/*func updateAI() {
	for {
		if aiPos < ballPos.Y {
			aiPos += 5.0
			aiPaddleRect = aiPaddleRect.Moved(pixel.V(0, 5))
		} else {
			aiPos -= 5.0
			aiPaddleRect = aiPaddleRect.Moved(pixel.V(0, -5))
		}
		time.Sleep(time.Millisecond * 16)
	}
}*/

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

		// Collision with AI paddle
		if ballPos.X+ballSize > aiPaddleRect.Min.X &&
			ballPos.Y > aiPaddleRect.Min.Y &&
			ballPos.Y < aiPaddleRect.Max.Y {
			ballDir.X = -ballDir.X
		}

		if ballPos.X-ballSize < 0 {
			aiScore++
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

func run() {

	win, err := pixelgl.NewWindow(pixelgl.WindowConfig{
		Title:  "Ping Pong",
		Bounds: pixel.R(0, 0, width, height),
		VSync:  true,
	})
	if err != nil {
		panic(err)
	}

	 // Crear y dibujar el fondo
	 bgImg := image.NewRGBA(image.Rect(0, 0, int(width), int(height)))
	 green := color.RGBA{0, 128, 0, 255}
	 for y := bgImg.Rect.Min.Y; y < bgImg.Rect.Max.Y; y++ {
		 for x := bgImg.Rect.Min.X; x < bgImg.Rect.Max.X; x++ {
			 bgImg.Set(x, y, green)
		 }
	 }

	  // Dibujar una línea blanca para la red
	  white := color.RGBA{255, 255, 255, 255}
	  for x := int(width/2) - 1; x <= int(width/2) + 1; x++ {
		  for y := bgImg.Rect.Min.Y; y < bgImg.Rect.Max.Y; y++ {
			  bgImg.Set(x, y, white)
		  }
	  }
	 bgPic := pixel.PictureDataFromImage(bgImg)
	 bgSprite := pixel.NewSprite(bgPic, bgPic.Bounds())



	 startAnimationDuration := 3.0 // duración en segundos
	 startAnimationSteps := 60     // número de pasos en la animación
	 startAtlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	 startTxt := text.New(pixel.V(width/2, height/2), startAtlas)
 
	 for i := 0; i <= startAnimationSteps; i++ {
		 opacity := float64(i) / float64(startAnimationSteps)
		 if i > startAnimationSteps/2 {
			 opacity = 1 - opacity
		 }
 
		 win.Clear(color.RGBA{
			 R: uint8(float64(colornames.Black.R) * (1-opacity) + float64(colornames.White.R) * opacity),
			 G: uint8(float64(colornames.Black.G) * (1-opacity) + float64(colornames.White.G) * opacity),
			 B: uint8(float64(colornames.Black.B) * (1-opacity) + float64(colornames.White.B) * opacity),
			 A: 255,
		 })
 
		 startTxt.Clear()
		 startTxt.Dot = pixel.V(width/2-40, height/2)
		 fmt.Fprintf(startTxt, "Ping Pong")
		 startTxt.Draw(win, pixel.IM.Scaled(startTxt.Orig, 3))
 
		 win.Update()
		 time.Sleep(time.Duration(startAnimationDuration / float64(startAnimationSteps) * float64(time.Second)))
	 }

	// bgPic := pixel.PictureDataFromImage(bgImg)
	// bgSprite := pixel.NewSprite(bgPic, bgPic.Bounds()) //Borrar
	 
	playerScorePos := pixel.V(width/4, height-40)
	aiScorePos := pixel.V(3*width/4, height-40)
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

	/*ballImg := image.NewRGBA(image.Rect(0, 0, int(ballSize*2), int(ballSize*2)))
	for y := ballImg.Rect.Min.Y; y < ballImg.Rect.Max.Y; y++ {
		for x := ballImg.Rect.Min.X; x < ballImg.Rect.Max.X; x++ {
			ballImg.Set(x, y, col)
		}
	}*/
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
		updateAI(win)
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
		paddleSprite.Draw(win, pixel.IM.Moved(aiPaddleRect.Center()))
		ballSprite.Draw(win, pixel.IM.Moved(ballPos))


		if playerScore >= 5 || aiScore >= 5 {
            win.Clear(colornames.Black)

            winnerTxt := text.New(pixel.V(width/2-40, height/2), text.NewAtlas(basicfont.Face7x13, text.ASCII))
            if playerScore >= 5 {
                fmt.Fprintf(winnerTxt, "Ganador: Jugador 1")
            } else {
                fmt.Fprintf(winnerTxt, "Ganador: Jugador 2")
            }

            winnerTxt.Draw(win, pixel.IM.Scaled(winnerTxt.Orig, 3))
            win.Update()

            // Mantener el mensaje en pantalla durante 5 segundos
            time.Sleep(5 * time.Second)
            
            // Cerrar la ventana y salir del programa
            win.Destroy()
            os.Exit(0)
        }

		// Render scores
		txt.Clear()
		txt.Dot = playerScorePos
		fmt.Fprintf(txt, "%d", playerScore)
		txt.Dot = aiScorePos
		fmt.Fprintf(txt, "%d", aiScore)
		txt.Draw(win, pixel.IM)

		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
