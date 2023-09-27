package scenes

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
	"pong/views"
	"pong/models"
	)

func Run() {

	win, err := pixelgl.NewWindow(pixelgl.WindowConfig{
		Title:  "Ping Pong",
		Bounds: pixel.R(0, 0, views.Width, views.Height),
		VSync:  true,
	})
	if err != nil {
		panic(err)
	}

	// Crea y dibuja el fondo
	bgImg := image.NewRGBA(image.Rect(0, 0, int(views.Width), int(views.Height)))
	blue := color.RGBA{70, 70, 255, 255}
	for y := bgImg.Rect.Min.Y; y < bgImg.Rect.Max.Y; y++ {
		for x := bgImg.Rect.Min.X; x < bgImg.Rect.Max.X; x++ {
			bgImg.Set(x, y, blue)
		}
	}

	// Dibuja la red
	white := color.RGBA{255, 255, 255, 255}
	for x := int(views.Width/2) - 1; x <= int(views.Width/2) + 1; x++ {
		for y := bgImg.Rect.Min.Y; y < bgImg.Rect.Max.Y; y++ {
			bgImg.Set(x, y, white)
		}
	}
	bgPic := pixel.PictureDataFromImage(bgImg)
	bgSprite := pixel.NewSprite(bgPic, bgPic.Bounds())

	playerScorePos := pixel.V(views.Width/4, views.Height-40)
	player2ScorePos := pixel.V(3*views.Width/4, views.Height-40)
	basicAtlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	txt := text.New(pixel.V(100, 500), basicAtlas)

	paddleImg := image.NewRGBA(image.Rect(0, 0, int(views.PaddleWidth), int(views.PaddleHeight)))
	col := color.RGBA{255, 255, 255, 255}
	for y := paddleImg.Rect.Min.Y; y < paddleImg.Rect.Max.Y; y++ {
		for x := paddleImg.Rect.Min.X; x < paddleImg.Rect.Max.X; x++ {
			paddleImg.Set(x, y, col)
		}
	}
	paddlePic := pixel.PictureDataFromImage(paddleImg)
	paddleSprite := pixel.NewSprite(paddlePic, paddlePic.Bounds())

	ballImg := image.NewRGBA(image.Rect(0, 0, int(views.BallSize*2), int(views.BallSize*2)))
	
	center := float64(int(views.BallSize)) // coordenadas del centro del círculo
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
		models.UpdatePlayer(win)
		wg.Done()
	}()
	go func() {
		models.Updateplayer2(win)
		wg.Done()
	}()
	go func() {
		models.UpdateBall()
		wg.Done()
	}()

	for !win.Closed() {
		win.Clear(colornames.Black)

		bgSprite.Draw(win, pixel.IM.Moved(win.Bounds().Center())) // borrar

		paddleSprite.Draw(win, pixel.IM.Moved(views.PlayerPaddleRect.Center()))
		paddleSprite.Draw(win, pixel.IM.Moved(views.Player2PaddleRect.Center()))
		ballSprite.Draw(win, pixel.IM.Moved(views.BallPos))


		if views.PlayerScore >= 5 || views.Player2Score >= 5 {
            win.Clear(colornames.Black)

            winnerTxt := text.New(pixel.V(views.Width/2-200, views.Height/2), text.NewAtlas(basicfont.Face7x13, text.ASCII))
            if views.PlayerScore >= 5 {
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
		fmt.Fprintf(txt, "%d", views.PlayerScore)
		txt.Dot = player2ScorePos
		fmt.Fprintf(txt, "%d", views.Player2Score)
		txt.Draw(win, pixel.IM)

		win.Update()
	}
}