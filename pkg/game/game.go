package game

import (
	"time"

	"github.com/garyloug/tetris/pkg/tetris"
	"github.com/garyloug/tetris/pkg/ui"
)

const (
	boardHeight    = 20
	boardWidth     = 10
	tetroSpawnX    = 5
	tetroSpawnY    = 0
	tetroQueueSize = 5
	userInputClock = 50 * time.Millisecond
	pauseForEffect = 100 * time.Millisecond
	moveDelay      = 500 * time.Millisecond
)

type Game struct {
	ui               ui.UI
	boardHeight      int
	boardWidth       int
	score            int
	level            int
	cleared          int
	pause            bool
	over             bool
	activeTetro      tetris.Tetro
	tetroQueue       []tetris.Tetro
	stationaryBlocks []tetris.Block
	speed            time.Duration
	moveTimer        time.Time
	moveDelay        time.Duration
	done             chan struct{}
}

func NewGame(ui ui.UI) Game {
	return Game{
		ui:          ui,
		boardHeight: boardHeight,
		boardWidth:  boardWidth,
		tetroQueue:  make([]tetris.Tetro, tetroQueueSize),
		moveDelay:   moveDelay,
		level:       0,
		score:       0,
		cleared:     0,
		done:        make(chan struct{}),
	}
}

func (g *Game) Start() <-chan struct{} {
	o, i, s, z, l, j, t := g.ui.GetBlockStyles()

	tetrisConfig := tetris.Config{
		SpawnX: tetroSpawnX,
		SpawnY: tetroSpawnY,
		StyleO: o,
		StyleI: i,
		StyleS: s,
		StyleZ: z,
		StyleL: l,
		StyleJ: j,
		StyleT: t,
	}

	tetris.Init(tetrisConfig)

	g.activeTetro = tetris.NewRandomTetro()
	g.stationaryBlocks = []tetris.Block{}

	for i := 0; i < tetroQueueSize; i++ {
		tetro := tetris.NewRandomTetro()
		g.tetroQueue[i] = tetro
	}

	if err := g.ui.Init(g.boardHeight, g.boardWidth); err != nil {
		// TODO implement log to file
		panic(err)
	}

	go g.userEvents()
	go g.run()

	g.resetMoveTimer()
	g.updateUI()
	g.ui.Start()

	return g.done
}

// run is the main game loop
func (g *Game) run() {
	for !g.over {

		// wait while game paused
		for g.pause {
			g.updateUI()
			time.Sleep(100 * time.Millisecond)
		}

		// move active tetro down in time with the game clock, speeds up with level
		time.Sleep(500*time.Millisecond - g.speed)
		g.moveTetroDown()

		// check for game over
		if g.reachedTop() {
			g.over = true
			break
		}

		// check for cleared lines and update score
		numCleared := g.clearCompletedLines()
		if numCleared > 0 {
			switch numCleared {
			case 1:
				g.score += 40 * (g.level + 1)
			case 2:
				g.score += 100 * (g.level + 1)
			case 3:
				g.score += 300 * (g.level + 1)
			case 4:
				g.score += 1200 * (g.level + 1)
			}
		}
		g.cleared += numCleared

		// update level based on cleared lines, cap at 20
		g.level = g.cleared / 10
		if g.level > 20 {
			g.level = 20
		}

		// update speed based on level
		g.speed = time.Duration(g.level*20) * time.Millisecond

		g.updateUI()
	}

	// final update with result
	g.updateUI()
}

// userEvents is a secondary loop to handle user inputs
func (g *Game) userEvents() {
	for {
		select {
		case event := <-g.ui.KeyPress():
			g.processKeyPress(event)
		case <-g.done:
			return
		}
	}
}

func (g *Game) processKeyPress(keypress ui.KeyPress) {
	switch keypress {
	case ui.KeyUp:
		if !g.over && !g.pause {
			g.rotateTetro()
		}
	case ui.KeyDown:
		if !g.over && !g.pause {
			g.moveTetroDown()
		}
	case ui.KeyLeft:
		if !g.over && !g.pause {
			g.moveTetroLeft()
		}
	case ui.KeyRight:
		if !g.over && !g.pause {
			g.moveTetroRight()
		}
	case ui.KeyPause:
		if !g.over {
			g.pause = !g.pause
		}
	case ui.KeyStop:
		g.ui.Stop()
		close(g.done)
	}
}

func (g *Game) moveTetroDown() {
	if g.activeTetro.CanMoveDown(g.boardHeight, g.boardWidth, g.stationaryBlocks) {
		g.activeTetro.MoveDown()
		g.resetMoveTimer()
	} else {
		// tetro can't move down, start/check timer
		if g.moveTimer.IsZero() {
			g.moveTimer = time.Now() // start timer
		} else if time.Since(g.moveTimer) >= g.moveDelay {
			g.nextTetro()      // timer expired, get next tetro
			g.resetMoveTimer() // reset for next tetro
		}
	}
	g.updateUI()
}

func (g *Game) moveTetroLeft() {
	if g.activeTetro.CanMoveLeft(g.boardHeight, g.boardWidth, g.stationaryBlocks) {
		g.activeTetro.MoveLeft()
		g.updateUI()

		if !g.activeTetro.CanMoveDown(g.boardHeight, g.boardWidth, g.stationaryBlocks) {
			g.resetMoveTimer()
		}
	}
}

func (g *Game) moveTetroRight() {
	if g.activeTetro.CanMoveRight(g.boardHeight, g.boardWidth, g.stationaryBlocks) {
		g.activeTetro.MoveRight()
		g.updateUI()

		if !g.activeTetro.CanMoveDown(g.boardHeight, g.boardWidth, g.stationaryBlocks) {
			g.resetMoveTimer()
		}
	}
}

func (g *Game) rotateTetro() {
	if g.activeTetro.CanRotate(g.boardHeight, g.boardWidth, g.stationaryBlocks) {
		g.activeTetro.Rotate()
		g.updateUI()

		if !g.activeTetro.CanMoveDown(g.boardHeight, g.boardWidth, g.stationaryBlocks) {
			g.resetMoveTimer()
		}
	}
}

func (g *Game) nextTetro() {
	// append active tetro blocks to stationary blocks
	g.stationaryBlocks = append(g.stationaryBlocks, g.activeTetro.Blocks()...)

	// take the next tetro from the queue
	g.activeTetro = g.tetroQueue[0]
	g.tetroQueue = g.tetroQueue[1:]

	// add a new random tetro to the end of the queue
	newTetro := tetris.NewRandomTetro()
	g.tetroQueue = append(g.tetroQueue, newTetro)
}

func (g *Game) reachedTop() bool {
	for _, block := range g.stationaryBlocks {
		_, y := block.Coordinates()
		if y < 1 {
			return true
		}
	}
	return false
}

func (g *Game) clearCompletedLines() (completedLines int) {
	for y := g.boardHeight - 1; y >= 0; y-- {
		lineComplete := true
		for x := 0; x < g.boardWidth; x++ {
			found := false
			for _, block := range g.stationaryBlocks {
				blockX, blockY := block.Coordinates()
				if blockX == x && blockY == y {
					found = true
					break
				}
			}
			if !found {
				lineComplete = false
				break
			}
		}
		if lineComplete {
			completedLines++
			g.removeLine(y)
			y++ // check the same line again after removing
		}
	}
	return completedLines
}

func (g *Game) removeLine(y int) {
	i := 0
	for _, block := range g.stationaryBlocks {
		_, blockY := block.Coordinates()
		if blockY != y {
			g.stationaryBlocks[i] = block
			i++
		}
	}
	g.stationaryBlocks = g.stationaryBlocks[:i]

	g.updateUI()
	time.Sleep(pauseForEffect)

	// lines above drop down
	for i := range g.stationaryBlocks {
		_, blockY := g.stationaryBlocks[i].Coordinates()
		if blockY < y {
			g.stationaryBlocks[i].MoveDown()
		}
	}

	g.updateUI()
}

func (g *Game) updateUI() {
	status := ui.Running
	if g.pause {
		status = ui.Pause
	} else if g.over {
		status = ui.GameOver
	}

	blocks := append(g.activeTetro.Blocks(), g.stationaryBlocks...)
	g.ui.Update(blocks, g.tetroQueue, g.score, g.level, g.cleared, status)
}

func (g *Game) resetMoveTimer() {
	g.moveTimer = time.Now()
}
