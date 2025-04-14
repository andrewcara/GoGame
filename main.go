package main

import (
	"HeadSoccer/Sprites"
	headsoccer_constants "HeadSoccer/constants"
	player "HeadSoccer/input"
	linalg "HeadSoccer/math/helper"
	"image/color"

	"HeadSoccer/initialization"
	"HeadSoccer/math/physics"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

// Gravity is positve since "down" on the screen is positive and up is negative
var gravity = linalg.Vector{X: 0, Y: 100.81}
var id = uuid.New()
var gameFont font.Face

var imageGameBG = Sprites.CreateImage("background.png")

const (
	physicsTickRate   = 1.0 / 100
	screenWidth       = 600
	screenHeight      = 300
	squareWidth       = 15
	maxSteps          = 3
	moveInputVelocity = 15
	moveForce         = 500.0 // Base movement force
	maxSpeed          = 200.0 // Maximum speed cap
	ground            = screenHeight - headsoccer_constants.GroundHeight
)

type Game struct {
	world          physics.PhysicsWorld
	Collision      bool
	accumulator    float64
	lastUpdateTime time.Time
	pressedKeys    []ebiten.Key
	player1_score  int
	player2_score  int
}

func initFont() font.Face {
	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}

	const dpi = 72
	fontFace, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    24,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
	return fontFace
}

func NewGame() *Game {
	// Initialize the font
	gameFont = initFont()

	// Initialize the world
	world := initialization.Setup(screenWidth, screenHeight, gravity)

	game := &Game{
		world:          world,
		Collision:      false,
		lastUpdateTime: time.Now(),
		accumulator:    0,
		player1_score:  0,
		player2_score:  0,
	}

	// Add Physics Objects to Game
	return game
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) HanlePlayerInput(player_moves player.PlayerMoves) {
	player := &g.world.Objects[player_moves.PlayerID]
	boundary := (*player).Shape.GetBoundaryPoints()
	onGround := boundary.MaxY >= float64(ground)-1
	g.pressedKeys = inpututil.AppendPressedKeys(g.pressedKeys[:0])
	for _, key := range g.pressedKeys {
		switch key.String() {
		case player_moves.Down:
			move := linalg.Vector{X: 0, Y: moveInputVelocity}
			(*player).SetVelocity(move.Add((*player).GetVelocity()))
		case player_moves.Up:
			if onGround {
				move := linalg.Vector{X: 0, Y: -100}
				(*player).SetVelocity(move.Add((*player).GetVelocity()))
			}
		case player_moves.Right:
			if onGround {
				move := linalg.Vector{X: moveInputVelocity, Y: 0}
				(*player).SetVelocity(move.Add((*player).GetVelocity()))
			}
		case player_moves.Left:
			if onGround {
				move := linalg.Vector{X: -moveInputVelocity, Y: 0}
				(*player).SetVelocity(move.Add((*player).GetVelocity()))
			}
		}
	}
}

func (g *Game) HandleUserInput() {
	player1_moves := player.PlayerMoves{Up: "ArrowUp", Down: "ArrowDown", Left: "ArrowLeft", Right: "ArrowRight", PlayerID: 0}
	player2_moves := player.PlayerMoves{Up: "W", Down: "S", Left: "A", Right: "D", PlayerID: 1}
	g.HanlePlayerInput(player1_moves)
	g.HanlePlayerInput(player2_moves)
}

func (g *Game) Update() error {
	g.HandleUserInput()
	currentTime := time.Now()
	frameTime := currentTime.Sub(g.lastUpdateTime).Seconds()
	g.lastUpdateTime = currentTime

	if frameTime > 0.25 {
		frameTime = 0.25
	}

	g.accumulator += frameTime
	steps := 0

	for g.accumulator >= physicsTickRate && steps < maxSteps {
		g.UpdatePhysics(physicsTickRate)
		g.accumulator -= physicsTickRate
		steps++
	}

	return nil
}

// Current implementation is that there must be two objects
func (g *Game) UpdatePhysics(timeDelta float64) {
	for i := 0; i < len(g.world.Objects); i++ {
		for j := i + 1; j < len(g.world.Objects); j++ {
			obj1 := g.world.Objects[i]
			obj2 := g.world.Objects[j]

			if physics.CollisionOccurs(obj1, obj2) {
				g.Collision = true
			} else {
				g.Collision = false
			}
			if !obj1.IsStatic {
				obj1.UpdateKinematics(screenWidth, ground, timeDelta, gravity)
			}
			if !obj2.IsStatic {
				obj2.UpdateKinematics(screenWidth, ground, timeDelta, gravity)
			}
		}
	}
}

func (g *Game) DrawBackground(screen *ebiten.Image) {
	w, h := imageGameBG.Bounds().Dx(), imageGameBG.Bounds().Dy()
	scaleW := screenWidth / float64(w)
	scaleH := screenHeight / float64(h)
	scale := scaleW
	if scale < scaleH {
		scale = scaleH
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-float64(w)/2, -float64(h)/2)
	op.GeoM.Scale(scale, scale)
	op.GeoM.Translate(screenWidth/2, screenHeight/2)
	op.Filter = ebiten.FilterLinear
	screen.DrawImage(imageGameBG, op)
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.DrawBackground(screen)
	// Draw the objects
	for _, obj := range g.world.Objects {
		obj.Shape.DrawShape(screen)
	}

	// Display score text
	scoreText := "Score: " + string(rune('0'+g.player1_score)) + " - " + string(rune('0'+g.player2_score))
	text.Draw(screen, scoreText, gameFont, screenWidth/2-100, 30, color.RGBA{255, 255, 255, 255})
	// Draw game title
	text.Draw(screen, "Head Soccer", gameFont, screenWidth/2-100, 60, color.RGBA{255, 255, 255, 255})
}

func main() {
	// Create the game
	game := NewGame()

	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Head Soccer")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
