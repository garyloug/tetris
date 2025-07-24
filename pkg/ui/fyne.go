//go:build !console

package ui

import (
	"embed"
	"fmt"
	"image/color"
	"sync"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/garyloug/tetris/pkg/tetris"
)

//go:embed assets/*.png
var assets embed.FS

const (
	appName    = "Tetris"
	blockSize  = 60 // px
	queueWidth = 4  // blocks (max width of any tetro)
)

var (
	fyneOImage *canvas.Image
	fyneIImage *canvas.Image
	fyneSImage *canvas.Image
	fyneZImage *canvas.Image
	fyneLImage *canvas.Image
	fyneJImage *canvas.Image
	fyneTImage *canvas.Image
)

type fyneUI struct {
	ui
	app         fyne.App
	window      fyne.Window
	boardGrid   *fyne.Container
	queueGrid   *fyne.Container
	scoreLabel  *widget.Label
	levelLabel  *widget.Label
	linesLabel  *widget.Label
	statusLabel *widget.Label
	boardImages [][]*canvas.Image
	queueImages [][]*canvas.Image

	// Key throttling
	lastKeyTime time.Time
	keyMutex    sync.Mutex
}

type FyneStyle struct {
	image *canvas.Image
}

func newFyneUI() (UI, func(), error) {
	cleanup := func() {
		// TODO
	}

	app := app.New()
	window := app.NewWindow(appName)

	f := &fyneUI{
		app:    app,
		window: window,
	}

	assetResource := func(name string) *fyne.StaticResource {
		data, err := assets.ReadFile("assets/" + name)
		if err != nil {
			panic("Failed to load asset: " + name + " - " + err.Error())
		}
		return fyne.NewStaticResource(name, data)
	}

	fyneOImage = canvas.NewImageFromResource(assetResource("yellow.png"))
	fyneOImage.FillMode = canvas.ImageFillStretch
	fyneIImage = canvas.NewImageFromResource(assetResource("light-blue.png"))
	fyneIImage.FillMode = canvas.ImageFillStretch
	fyneSImage = canvas.NewImageFromResource(assetResource("green.png"))
	fyneSImage.FillMode = canvas.ImageFillStretch
	fyneZImage = canvas.NewImageFromResource(assetResource("red.png"))
	fyneZImage.FillMode = canvas.ImageFillStretch
	fyneLImage = canvas.NewImageFromResource(assetResource("orange.png"))
	fyneLImage.FillMode = canvas.ImageFillStretch
	fyneJImage = canvas.NewImageFromResource(assetResource("blue.png"))
	fyneJImage.FillMode = canvas.ImageFillStretch
	fyneTImage = canvas.NewImageFromResource(assetResource("purple.png"))
	fyneTImage.FillMode = canvas.ImageFillStretch

	f.ui = ui{
		eventChan: make(chan KeyPress, 5), // TODO 5
		oStyles:   tetris.BlockStyles{Block0: FyneStyle{fyneOImage}, Block1: FyneStyle{fyneOImage}, Block2: FyneStyle{fyneOImage}, Block3: FyneStyle{fyneOImage}},
		iStyles:   tetris.BlockStyles{Block0: FyneStyle{fyneIImage}, Block1: FyneStyle{fyneIImage}, Block2: FyneStyle{fyneIImage}, Block3: FyneStyle{fyneIImage}},
		sStyles:   tetris.BlockStyles{Block0: FyneStyle{fyneSImage}, Block1: FyneStyle{fyneSImage}, Block2: FyneStyle{fyneSImage}, Block3: FyneStyle{fyneSImage}},
		zStyles:   tetris.BlockStyles{Block0: FyneStyle{fyneZImage}, Block1: FyneStyle{fyneZImage}, Block2: FyneStyle{fyneZImage}, Block3: FyneStyle{fyneZImage}},
		lStyles:   tetris.BlockStyles{Block0: FyneStyle{fyneLImage}, Block1: FyneStyle{fyneLImage}, Block2: FyneStyle{fyneLImage}, Block3: FyneStyle{fyneLImage}},
		jStyles:   tetris.BlockStyles{Block0: FyneStyle{fyneJImage}, Block1: FyneStyle{fyneJImage}, Block2: FyneStyle{fyneJImage}, Block3: FyneStyle{fyneJImage}},
		tStyles:   tetris.BlockStyles{Block0: FyneStyle{fyneTImage}, Block1: FyneStyle{fyneTImage}, Block2: FyneStyle{fyneTImage}, Block3: FyneStyle{fyneTImage}},
	}

	return f, cleanup, nil
}

func (f *fyneUI) Init(boardH, boardW int) error {
	if boardH <= 0 || boardW <= 0 {
		return fmt.Errorf("invalid board dimensions")
	}
	f.boardHeight = boardH
	f.boardWidth = boardW

	boardPixelWidth := boardW * blockSize
	boardPixelHeight := boardH * blockSize
	queuePixelWidth := queueWidth * blockSize
	windowWidth := boardPixelWidth + queuePixelWidth + 200 // board + queue + right panel
	windowHeight := boardPixelHeight + 100                 // 100px for padding/title bar

	f.window.Resize(fyne.NewSize(float32(windowWidth), float32(windowHeight)))
	f.window.SetFixedSize(true) // non-resizable

	// Board grid - use border layout to contain the background and images
	f.boardGrid = container.NewBorder(nil, nil, nil, nil,
		canvas.NewRectangle(color.Transparent), // placeholder for proper sizing
	)
	f.boardGrid.Resize(fyne.NewSize(float32(boardW*blockSize), float32(boardH*blockSize)))

	// Gradient background for main board
	boardBackground := canvas.NewLinearGradient(
		color.RGBA{R: 20, G: 20, B: 30, A: 255},
		color.RGBA{R: 40, G: 40, B: 50, A: 255},
		0,
	)
	boardBackground.Resize(fyne.NewSize(float32(boardW*blockSize), float32(boardH*blockSize)))
	boardBackground.Move(fyne.NewPos(0, 0))

	// Create container for board with background and images
	boardContent := container.NewWithoutLayout()
	boardContent.Add(boardBackground)

	// Create a grid of images for main board
	f.boardImages = make([][]*canvas.Image, boardH)
	for y := 0; y < boardH; y++ {
		f.boardImages[y] = make([]*canvas.Image, boardW)
		for x := 0; x < boardW; x++ {
			img := canvas.NewImageFromFile("")
			img.Resize(fyne.NewSize(blockSize, blockSize))
			img.Move(fyne.NewPos(float32(x*blockSize), float32(y*blockSize)))
			img.FillMode = canvas.ImageFillStretch
			img.Hide()
			f.boardImages[y][x] = img
			boardContent.Add(img)
		}
	}

	// Replace placeholder with actual content
	f.boardGrid = container.NewBorder(nil, nil, nil, nil, boardContent)

	// Create container for queue images
	queueContent := container.NewWithoutLayout()

	// Create a grid of images for queue board
	f.queueImages = make([][]*canvas.Image, boardH)
	for y := 0; y < boardH; y++ {
		f.queueImages[y] = make([]*canvas.Image, queueWidth)
		for x := 0; x < queueWidth; x++ {
			img := canvas.NewImageFromFile("")
			img.Resize(fyne.NewSize(blockSize, blockSize))
			img.Move(fyne.NewPos(float32(x*blockSize), float32(y*blockSize)))
			img.FillMode = canvas.ImageFillStretch
			img.Hide()
			f.queueImages[y][x] = img
			queueContent.Add(img)
		}
	}

	// Create fixed-size spacers to force proper layout
	boardSpacer := canvas.NewRectangle(color.Transparent)
	boardSpacer.Resize(fyne.NewSize(float32(boardW*blockSize), float32(boardH*blockSize)))

	queueSpacer := canvas.NewRectangle(color.Transparent)
	queueSpacer.Resize(fyne.NewSize(float32(queueWidth*blockSize), float32(boardH*blockSize)))

	// Board grid - overlay content on spacer
	f.boardGrid = container.NewStack(boardSpacer, boardContent)

	// Queue grid - overlay content on spacer
	f.queueGrid = container.NewStack(queueSpacer, queueContent)

	// Key bindings
	f.window.Canvas().SetOnTypedKey(func(key *fyne.KeyEvent) {
		f.keyMutex.Lock()
		defer f.keyMutex.Unlock()

		now := time.Now()

		throttleTime := 10 * time.Millisecond // TODO

		if now.Sub(f.lastKeyTime) < throttleTime {
			return // too soon
		}

		f.lastKeyTime = now

		var keyPress KeyPress
		switch key.Name {
		case fyne.KeyLeft:
			keyPress = KeyLeft
		case fyne.KeyRight:
			keyPress = KeyRight
		case fyne.KeyDown:
			keyPress = KeyDown
		case fyne.KeyUp:
			keyPress = KeyUp
		case fyne.KeySpace:
			keyPress = KeyPause
		case fyne.KeyEscape:
			keyPress = KeyStop
		default:
			return // unknown key
		}

		// non-blocking send
		select {
		case f.ui.eventChan <- keyPress:
		default:
			// channel full
		}
	})

	// UI elements
	f.scoreLabel = widget.NewLabel("Score: 0")
	f.levelLabel = widget.NewLabel("Level: 1")
	f.linesLabel = widget.NewLabel("Lines: 0")
	f.statusLabel = widget.NewLabel("Status: Running")

	// Instructions
	instructions := widget.NewLabel(`Controls:
← - Left
→ - Right
↓ - Down
↑ - Rotate
⎵ - Pause
Esc - Quit`)

	// Layout
	rightPanel := container.NewVBox(
		f.scoreLabel,
		f.levelLabel,
		f.linesLabel,
		widget.NewSeparator(),
		instructions,
		widget.NewSeparator(),
		f.statusLabel,
	)

	// Layout - simple horizontal arrangement: main board | queue | text
	content := container.NewHBox(
		f.boardGrid,
		f.queueGrid,
		rightPanel,
	)

	f.window.SetContent(content)

	return nil
}

func (f *fyneUI) GetBlockStyles() (o, i, s, z, l, j, t tetris.BlockStyles) {
	return f.oStyles, f.iStyles, f.sStyles, f.zStyles, f.lStyles, f.jStyles, f.tStyles
}

func (f *fyneUI) Start() {

	// go f.run()

	f.window.ShowAndRun()
}

func (f *fyneUI) Stop() {

}

// TODO comment on changes to this function
func (f *fyneUI) Update(blocks []tetris.Block, queue []tetris.Tetro, score, level, linesCleared int, status Status) {
	fyne.Do(func() {
		// Create a map of positions that should have blocks on main board
		activePositions := make(map[[2]int]tetris.Block)
		for _, block := range blocks {
			x, y := block.Coordinates()
			if x >= 0 && x < f.boardWidth && y >= 0 && y < f.boardHeight {
				activePositions[[2]int{x, y}] = block
			}
		}

		// Update main board - only update positions that have changed
		for y := 0; y < f.boardHeight; y++ {
			for x := 0; x < f.boardWidth; x++ {
				img := f.boardImages[y][x]
				pos := [2]int{x, y}

				if block, hasBlock := activePositions[pos]; hasBlock {
					// should have a block
					style := block.Style().(FyneStyle)
					if style.image != nil {
						img.File = style.image.File
						img.Resource = style.image.Resource
						img.Show()
					}
				} else {
					// should be empty
					img.Hide()
					img.File = ""
					img.Resource = nil
				}
			}
		}

		// Create a map of positions that should have blocks on queue board
		queuePositions := make(map[[2]int]tetris.Block)
		currentY := 0
		for _, tetro := range queue {
			tetroBlocks := tetro.Blocks()

			// Find the bounding box of the tetromino
			if len(tetroBlocks) > 0 {
				minX, maxX := tetroBlocks[0].Coordinates()
				minY, maxY := minX, maxX
				for _, block := range tetroBlocks {
					x, y := block.Coordinates()
					if x < minX {
						minX = x
					}
					if x > maxX {
						maxX = x
					}
					if y < minY {
						minY = y
					}
					if y > maxY {
						maxY = y
					}
				}

				// Place tetromino blocks in queue board
				for _, block := range tetroBlocks {
					x, y := block.Coordinates()
					// Normalize to start at (0,0) and offset by current Y position
					queueX := x - minX
					queueY := (y - minY) + currentY

					if queueX >= 0 && queueX < queueWidth && queueY >= 0 && queueY < f.boardHeight {
						queuePositions[[2]int{queueX, queueY}] = block
					}
				}

				// Move to next position for next tetromino (add some spacing)
				currentY += (maxY - minY) + 2
				if currentY >= f.boardHeight {
					break // No more room
				}
			}
		}

		// Update queue board - only update positions that have changed
		for y := 0; y < f.boardHeight; y++ {
			for x := 0; x < queueWidth; x++ {
				img := f.queueImages[y][x]
				pos := [2]int{x, y}

				if block, hasBlock := queuePositions[pos]; hasBlock {
					// should have a block
					style := block.Style().(FyneStyle)
					if style.image != nil {
						img.File = style.image.File
						img.Resource = style.image.Resource
						img.Show()
					}
				} else {
					// should be empty
					img.Hide()
					img.File = ""
					img.Resource = nil
				}
			}
		}

		f.boardGrid.Refresh()
		f.queueGrid.Refresh()
	})
}
