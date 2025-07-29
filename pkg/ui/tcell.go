package ui

import (
	"fmt"
	"log"
	"strconv"

	"github.com/garyloug/tetris/pkg/tetris"
	tcellLib "github.com/gdamore/tcell/v2"
)

const (
	full        = '█'
	shade       = '▓'
	oColour     = tcellLib.ColorYellow
	iColour     = tcellLib.ColorMediumTurquoise
	sColour     = tcellLib.ColorLimeGreen
	zColour     = tcellLib.ColorRed
	lColour     = tcellLib.ColorOrange
	jColour     = tcellLib.ColorDarkBlue
	tColour     = tcellLib.ColorRebeccaPurple
	bgColour    = tcellLib.ColorDimGrey
	xMultiplier = 2 // block is 2x1 characters wide
	yMultiplier = 1 // block is 1 character tall
)

var (
	// "constant" styles for each tetro type
	oStyle = TcellStyle{
		style: tcellLib.StyleDefault.Background(oColour).Foreground(oColour),
		fill:  full,
	}
	iStyle = TcellStyle{
		style: tcellLib.StyleDefault.Background(iColour).Foreground(iColour),
		fill:  full,
	}
	sStyle = TcellStyle{
		style: tcellLib.StyleDefault.Background(sColour).Foreground(sColour),
		fill:  full,
	}
	zStyle = TcellStyle{
		style: tcellLib.StyleDefault.Background(zColour).Foreground(zColour),
		fill:  full,
	}
	lStyle = TcellStyle{
		style: tcellLib.StyleDefault.Background(lColour).Foreground(lColour),
		fill:  full,
	}
	jStyle = TcellStyle{
		style: tcellLib.StyleDefault.Background(jColour).Foreground(jColour),
		fill:  full,
	}
	tStyle = TcellStyle{
		style: tcellLib.StyleDefault.Background(tColour).Foreground(tColour),
		fill:  full,
	}
)

type TcellStyle struct {
	style tcellLib.Style
	fill  rune
}

type tcell struct {
	ui
	screen     tcellLib.Screen
	boardStyle tcellLib.Style
	quit       chan struct{}
	bgFill     rune
}

func newTcellUI() (UI, func()) {
	defStyle := tcellLib.StyleDefault.Background(tcellLib.ColorReset).Foreground(tcellLib.ColorReset)

	// init screen
	s, err := tcellLib.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := s.Init(); err != nil {
		log.Fatalf("%+v", err)
	}
	s.SetStyle(defStyle)

	s.Clear()

	// cleanup function
	cleanup := func() {
		possiblePanic := recover() // capture any panic so it can be printed
		s.Fini()                   // must be called exactly once
		if possiblePanic != nil {
			panic(possiblePanic)
		}
	}

	tc := &tcell{
		screen:     s,
		boardStyle: tcellLib.StyleDefault.Background(bgColour).Foreground(bgColour),
		quit:       make(chan struct{}),
		bgFill:     full,
	}

	tc.ui = ui{
		eventChan: make(chan KeyPress, 10),
		oStyles:   tetris.BlockStyles{Block0: oStyle, Block1: oStyle, Block2: oStyle, Block3: oStyle},
		iStyles:   tetris.BlockStyles{Block0: iStyle, Block1: iStyle, Block2: iStyle, Block3: iStyle},
		sStyles:   tetris.BlockStyles{Block0: sStyle, Block1: sStyle, Block2: sStyle, Block3: sStyle},
		zStyles:   tetris.BlockStyles{Block0: zStyle, Block1: zStyle, Block2: zStyle, Block3: zStyle},
		lStyles:   tetris.BlockStyles{Block0: lStyle, Block1: lStyle, Block2: lStyle, Block3: lStyle},
		jStyles:   tetris.BlockStyles{Block0: jStyle, Block1: jStyle, Block2: jStyle, Block3: jStyle},
		tStyles:   tetris.BlockStyles{Block0: tStyle, Block1: tStyle, Block2: tStyle, Block3: tStyle},
	}

	tc.screen.Show()

	return tc, cleanup
}

func newTcellDevUI() (UI, func()) {
	tcRaw, cleanup := newTcellUI()
	tc := tcRaw.(*tcell)
	tc.setDevStyles()
	tc.bgFill = '·'
	return tc, cleanup
}

func (tc *tcell) Init(boardH, boardW int) error {
	if boardH <= 0 || boardW <= 0 {
		return fmt.Errorf("invalid board dimensions")
	}
	tc.boardHeight = boardH
	tc.boardWidth = boardW
	return nil
}

func (tc *tcell) GetBlockStyles() (o, i, s, z, l, j, t tetris.BlockStyles) {
	return tc.oStyles, tc.iStyles, tc.sStyles, tc.zStyles, tc.lStyles, tc.jStyles, tc.tStyles
}

func (tc *tcell) Update(blocks []tetris.Block, queue []tetris.Tetro, score, level, linesCleared int, status Status) {
	tc.screen.Clear()

	// draw the board
	for y := 0; y < tc.boardHeight; y++ {
		for x := 0; x < tc.boardWidth; x++ {
			tc.screen.SetContent(x*xMultiplier, y*yMultiplier, tc.bgFill, nil, tc.boardStyle)
			tc.screen.SetContent(x*xMultiplier+1, y*yMultiplier, tc.bgFill, nil, tc.boardStyle)
		}
	}

	// draw the blocks
	for _, block := range blocks {
		x, y := block.Coordinates()
		style := block.Style().(TcellStyle)
		fill := style.fill
		tcellStyle := style.style

		tc.screen.SetContent(x*xMultiplier, y*yMultiplier, fill, nil, tcellStyle)
		tc.screen.SetContent(x*xMultiplier+1, y*yMultiplier, fill, nil, tcellStyle)
	}

	// draw the tetro queue
	xOffset := tc.boardWidth + 5
	yOffset := 1
	for _, tetro := range queue {
		minY, maxY := 0, 0
		first := true

		// find the height of the tetro
		for _, block := range tetro.Blocks() {
			_, y := block.Coordinates()
			if first {
				minY, maxY = y, y
				first = false
			} else {
				if y < minY {
					minY = y
				}
				if y > maxY {
					maxY = y
				}
			}
		}

		// draw the tetro, adjust Y to allow for height
		yAdjustment := yOffset - minY
		for _, block := range tetro.Blocks() {
			x, y := block.Coordinates()
			style := block.Style().(TcellStyle)
			fill := style.fill
			tcellStyle := style.style

			adjustedY := y + yAdjustment
			tc.screen.SetContent(xOffset+x*xMultiplier, adjustedY*yMultiplier, fill, nil, tcellStyle)
			tc.screen.SetContent(xOffset+x*xMultiplier+1, adjustedY*yMultiplier, fill, nil, tcellStyle)
		}

		// move yOffset by the actual height of this tetro + 1 for gap
		tetroHeight := maxY - minY + 1
		yOffset += tetroHeight + 1
	}

	// draw the score, level and lines cleared
	style := tcellLib.StyleDefault.Foreground(tcellLib.ColorWhite)

	scoreText := "Score: " + strconv.Itoa(score)
	for i, char := range scoreText {
		tc.screen.SetContent(tc.boardWidth*xMultiplier+15+i, 0, char, nil, style)
	}

	levelText := "Level: " + strconv.Itoa(level)
	for i, char := range levelText {
		tc.screen.SetContent(tc.boardWidth*xMultiplier+15+i, 1, char, nil, style)
	}

	clearedText := "Lines: " + strconv.Itoa(linesCleared)
	for i, char := range clearedText {
		tc.screen.SetContent(tc.boardWidth*xMultiplier+15+i, 2, char, nil, style)
	}

	// draw instructions
	instructions := []string{
		"← - Left",
		"→ - Right",
		"↓ - Down",
		"↑ - Rotate",
		"⎵ - Pause",
		"Esc - Quit",
	}
	for i, line := range instructions {
		for j, char := range line {
			tc.screen.SetContent(tc.boardWidth*xMultiplier+15+j, 5+i, char, nil, style)
		}
	}

	// draw status
	var statusText string
	switch status {
	case Running:
		statusText = "Status: Running"
	case Pause:
		statusText = "Status: Paused"
	case GameOver:
		statusText = "Status: Game Over"
	}
	for i, char := range statusText {
		tc.screen.SetContent(tc.boardWidth*xMultiplier+15+i, 12, char, nil, style)
	}

	// show the screen
	tc.screen.Show()
}

func (tc *tcell) Start() {
	go tc.run()
}

func (tc *tcell) Stop() {
	close(tc.quit)
	tc.screen.Clear()
	tc.screen.Show()
	// intentionally not calling tc.screen.Fini()
	// this is called in the cleanup function
}

func (tc *tcell) run() {
	for {
		// check for quit signal
		select {
		case <-tc.quit:
			return
		default:
		}

		// poll for events
		ev := tc.screen.PollEvent()
		if ev == nil {
			continue
		}

		// process event and send to event channel
		switch ev := ev.(type) {
		case *tcellLib.EventResize:
			tc.screen.Sync()
		case *tcellLib.EventKey:
			if ev.Key() == tcellLib.KeyEscape || ev.Key() == tcellLib.KeyCtrlC {
				tc.ui.eventChan <- KeyStop
			} else if ev.Key() == tcellLib.KeyRight {
				tc.ui.eventChan <- KeyRight
			} else if ev.Key() == tcellLib.KeyLeft {
				tc.ui.eventChan <- KeyLeft
			} else if ev.Key() == tcellLib.KeyDown {
				tc.ui.eventChan <- KeyDown
			} else if ev.Key() == tcellLib.KeyUp {
				tc.ui.eventChan <- KeyUp
			} else if ev.Rune() == ' ' {
				tc.ui.eventChan <- KeyPause
			}
		}
	}
}

func (tc *tcell) setDevStyles() {
	styles := []*tetris.BlockStyles{&tc.oStyles, &tc.iStyles, &tc.sStyles, &tc.zStyles, &tc.lStyles, &tc.jStyles, &tc.tStyles}
	for _, style := range styles {
		style.Block0 = TcellStyle{
			style: style.Block0.(TcellStyle).style,
			fill:  '0',
		}
		style.Block1 = TcellStyle{
			style: style.Block1.(TcellStyle).style,
			fill:  '1',
		}
		style.Block2 = TcellStyle{
			style: style.Block2.(TcellStyle).style,
			fill:  '2',
		}
		style.Block3 = TcellStyle{
			style: style.Block3.(TcellStyle).style,
			fill:  '3',
		}
	}
}
