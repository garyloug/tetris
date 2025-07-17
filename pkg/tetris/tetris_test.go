package tetris

import (
	"testing"
)

func TestInit(t *testing.T) {
	testConfig := Config{
		SpawnX: 5,
		SpawnY: 0,
		StyleO: BlockStyles{Block0: "O0", Block1: "O1", Block2: "O2", Block3: "O3"},
		StyleI: BlockStyles{Block0: "I0", Block1: "I1", Block2: "I2", Block3: "I3"},
		StyleS: BlockStyles{Block0: "S0", Block1: "S1", Block2: "S2", Block3: "S3"},
		StyleZ: BlockStyles{Block0: "Z0", Block1: "Z1", Block2: "Z2", Block3: "Z3"},
		StyleL: BlockStyles{Block0: "L0", Block1: "L1", Block2: "L2", Block3: "L3"},
		StyleJ: BlockStyles{Block0: "J0", Block1: "J1", Block2: "J2", Block3: "J3"},
		StyleT: BlockStyles{Block0: "T0", Block1: "T1", Block2: "T2", Block3: "T3"},
	}

	Init(testConfig)

	if config.SpawnX != testConfig.SpawnX {
		t.Errorf("Init() SpawnX = %v, want %v", config.SpawnX, testConfig.SpawnX)
	}
	if config.SpawnY != testConfig.SpawnY {
		t.Errorf("Init() SpawnY = %v, want %v", config.SpawnY, testConfig.SpawnY)
	}
	if config.StyleO.Block0 != testConfig.StyleO.Block0 {
		t.Errorf("Init() StyleO.Block0 = %v, want %v", config.StyleO.Block0, testConfig.StyleO.Block0)
	}
}

func TestNewRandomTetro(t *testing.T) {
	Init(Config{
		SpawnX: 5,
		SpawnY: 0,
		StyleO: BlockStyles{Block0: "O0", Block1: "O1", Block2: "O2", Block3: "O3"},
		StyleI: BlockStyles{Block0: "I0", Block1: "I1", Block2: "I2", Block3: "I3"},
		StyleS: BlockStyles{Block0: "S0", Block1: "S1", Block2: "S2", Block3: "S3"},
		StyleZ: BlockStyles{Block0: "Z0", Block1: "Z1", Block2: "Z2", Block3: "Z3"},
		StyleL: BlockStyles{Block0: "L0", Block1: "L1", Block2: "L2", Block3: "L3"},
		StyleJ: BlockStyles{Block0: "J0", Block1: "J1", Block2: "J2", Block3: "J3"},
		StyleT: BlockStyles{Block0: "T0", Block1: "T1", Block2: "T2", Block3: "T3"},
	})

	tetro := NewRandomTetro()
	if tetro == nil {
		t.Fatal("NewRandomTetro() returned nil")
	}

	blocks := tetro.Blocks()
	if len(blocks) != 4 {
		t.Errorf("NewRandomTetro() should return tetro with 4 blocks, got %d", len(blocks))
	}

	for i, block := range blocks {
		x, y := block.Coordinates()
		if block.Style() == nil {
			t.Errorf("Block %d has nil style", i)
		}
		if x < -3 || x > 8 || y < -3 || y > 3 {
			t.Errorf("Block %d has coordinates (%d, %d) outside of expected range", i, x, y)
		}
	}
}

func TestTetro_MoveRight(t *testing.T) {
	Init(Config{
		SpawnX: 5,
		SpawnY: 0,
		StyleO: BlockStyles{Block0: "O0", Block1: "O1", Block2: "O2", Block3: "O3"},
		StyleI: BlockStyles{Block0: "I0", Block1: "I1", Block2: "I2", Block3: "I3"},
		StyleS: BlockStyles{Block0: "S0", Block1: "S1", Block2: "S2", Block3: "S3"},
		StyleZ: BlockStyles{Block0: "Z0", Block1: "Z1", Block2: "Z2", Block3: "Z3"},
		StyleL: BlockStyles{Block0: "L0", Block1: "L1", Block2: "L2", Block3: "L3"},
		StyleJ: BlockStyles{Block0: "J0", Block1: "J1", Block2: "J2", Block3: "J3"},
		StyleT: BlockStyles{Block0: "T0", Block1: "T1", Block2: "T2", Block3: "T3"},
	})

	o := newO()
	originalBlocks := o.Blocks()

	o.MoveRight()
	newBlocks := o.Blocks()

	for i := 0; i < 4; i++ {
		origX, origY := originalBlocks[i].Coordinates()
		newX, newY := newBlocks[i].Coordinates()

		if newX != origX+1 {
			t.Errorf("Block %d x-coordinate should increase by 1: got %d, want %d", i, newX, origX+1)
		}
		if newY != origY {
			t.Errorf("Block %d y-coordinate should not change: got %d, want %d", i, newY, origY)
		}
	}
}

func TestTetro_MoveLeft(t *testing.T) {
	Init(Config{
		SpawnX: 5,
		SpawnY: 0,
		StyleO: BlockStyles{Block0: "O0", Block1: "O1", Block2: "O2", Block3: "O3"},
		StyleI: BlockStyles{Block0: "I0", Block1: "I1", Block2: "I2", Block3: "I3"},
		StyleS: BlockStyles{Block0: "S0", Block1: "S1", Block2: "S2", Block3: "S3"},
		StyleZ: BlockStyles{Block0: "Z0", Block1: "Z1", Block2: "Z2", Block3: "Z3"},
		StyleL: BlockStyles{Block0: "L0", Block1: "L1", Block2: "L2", Block3: "L3"},
		StyleJ: BlockStyles{Block0: "J0", Block1: "J1", Block2: "J2", Block3: "J3"},
		StyleT: BlockStyles{Block0: "T0", Block1: "T1", Block2: "T2", Block3: "T3"},
	})

	o := newO()
	originalBlocks := o.Blocks()

	o.MoveLeft()
	newBlocks := o.Blocks()

	for i := 0; i < 4; i++ {
		origX, origY := originalBlocks[i].Coordinates()
		newX, newY := newBlocks[i].Coordinates()

		if newX != origX-1 {
			t.Errorf("Block %d x-coordinate should decrease by 1: got %d, want %d", i, newX, origX-1)
		}
		if newY != origY {
			t.Errorf("Block %d y-coordinate should not change: got %d, want %d", i, newY, origY)
		}
	}
}

func TestTetro_MoveDown(t *testing.T) {
	Init(Config{
		SpawnX: 5,
		SpawnY: 0,
		StyleO: BlockStyles{Block0: "O0", Block1: "O1", Block2: "O2", Block3: "O3"},
		StyleI: BlockStyles{Block0: "I0", Block1: "I1", Block2: "I2", Block3: "I3"},
		StyleS: BlockStyles{Block0: "S0", Block1: "S1", Block2: "S2", Block3: "S3"},
		StyleZ: BlockStyles{Block0: "Z0", Block1: "Z1", Block2: "Z2", Block3: "Z3"},
		StyleL: BlockStyles{Block0: "L0", Block1: "L1", Block2: "L2", Block3: "L3"},
		StyleJ: BlockStyles{Block0: "J0", Block1: "J1", Block2: "J2", Block3: "J3"},
		StyleT: BlockStyles{Block0: "T0", Block1: "T1", Block2: "T2", Block3: "T3"},
	})

	o := newO()
	originalBlocks := o.Blocks()

	o.MoveDown()
	newBlocks := o.Blocks()

	for i := 0; i < 4; i++ {
		origX, origY := originalBlocks[i].Coordinates()
		newX, newY := newBlocks[i].Coordinates()

		if newX != origX {
			t.Errorf("Block %d x-coordinate should not change: got %d, want %d", i, newX, origX)
		}
		if newY != origY+1 {
			t.Errorf("Block %d y-coordinate should increase by 1: got %d, want %d", i, newY, origY+1)
		}
	}
}

func TestTetro_CanMoveDown(t *testing.T) {
	Init(Config{
		SpawnX: 5,
		SpawnY: 0,
		StyleO: BlockStyles{Block0: "O0", Block1: "O1", Block2: "O2", Block3: "O3"},
		StyleI: BlockStyles{Block0: "I0", Block1: "I1", Block2: "I2", Block3: "I3"},
		StyleS: BlockStyles{Block0: "S0", Block1: "S1", Block2: "S2", Block3: "S3"},
		StyleZ: BlockStyles{Block0: "Z0", Block1: "Z1", Block2: "Z2", Block3: "Z3"},
		StyleL: BlockStyles{Block0: "L0", Block1: "L1", Block2: "L2", Block3: "L3"},
		StyleJ: BlockStyles{Block0: "J0", Block1: "J1", Block2: "J2", Block3: "J3"},
		StyleT: BlockStyles{Block0: "T0", Block1: "T1", Block2: "T2", Block3: "T3"},
	})

	tests := []struct {
		name             string
		boardHeight      int
		boardWidth       int
		stationaryBlocks []Block
		expected         bool
		setupTetro       func() Tetro
	}{
		{
			name:             "can move down in empty board",
			boardHeight:      20,
			boardWidth:       10,
			stationaryBlocks: []Block{},
			expected:         true,
			setupTetro:       func() Tetro { return newO() },
		},
		{
			name:             "cannot move down at bottom",
			boardHeight:      3,
			boardWidth:       10,
			stationaryBlocks: []Block{},
			expected:         false,
			setupTetro: func() Tetro {
				o := newO()
				o.MoveDown()
				return o
			},
		},
		{
			name:        "cannot move down due to stationary block",
			boardHeight: 20,
			boardWidth:  10,
			stationaryBlocks: []Block{
				{x: 5, y: 2, style: "block"},
			},
			expected:   false,
			setupTetro: func() Tetro { return newO() },
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tetro := tt.setupTetro()
			result := tetro.CanMoveDown(tt.boardHeight, tt.boardWidth, tt.stationaryBlocks)
			if result != tt.expected {
				t.Errorf("CanMoveDown() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestTetro_CanMoveRight(t *testing.T) {
	Init(Config{
		SpawnX: 5,
		SpawnY: 0,
		StyleO: BlockStyles{Block0: "O0", Block1: "O1", Block2: "O2", Block3: "O3"},
		StyleI: BlockStyles{Block0: "I0", Block1: "I1", Block2: "I2", Block3: "I3"},
		StyleS: BlockStyles{Block0: "S0", Block1: "S1", Block2: "S2", Block3: "S3"},
		StyleZ: BlockStyles{Block0: "Z0", Block1: "Z1", Block2: "Z2", Block3: "Z3"},
		StyleL: BlockStyles{Block0: "L0", Block1: "L1", Block2: "L2", Block3: "L3"},
		StyleJ: BlockStyles{Block0: "J0", Block1: "J1", Block2: "J2", Block3: "J3"},
		StyleT: BlockStyles{Block0: "T0", Block1: "T1", Block2: "T2", Block3: "T3"},
	})

	tests := []struct {
		name             string
		boardHeight      int
		boardWidth       int
		stationaryBlocks []Block
		expected         bool
		setupTetro       func() Tetro
	}{
		{
			name:             "can move right in empty board",
			boardHeight:      20,
			boardWidth:       10,
			stationaryBlocks: []Block{},
			expected:         true,
			setupTetro:       func() Tetro { return newO() },
		},
		{
			name:             "cannot move right at right edge",
			boardHeight:      20,
			boardWidth:       7,
			stationaryBlocks: []Block{},
			expected:         false,
			setupTetro: func() Tetro {
				o := newO()
				o.MoveRight()
				return o
			},
		},
		{
			name:        "cannot move right due to stationary block",
			boardHeight: 20,
			boardWidth:  10,
			stationaryBlocks: []Block{
				{x: 7, y: 0, style: "block"},
			},
			expected:   false,
			setupTetro: func() Tetro { return newO() },
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tetro := tt.setupTetro()
			result := tetro.CanMoveRight(tt.boardHeight, tt.boardWidth, tt.stationaryBlocks)
			if result != tt.expected {
				t.Errorf("CanMoveRight() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestTetro_CanMoveLeft(t *testing.T) {
	Init(Config{
		SpawnX: 5,
		SpawnY: 0,
		StyleO: BlockStyles{Block0: "O0", Block1: "O1", Block2: "O2", Block3: "O3"},
		StyleI: BlockStyles{Block0: "I0", Block1: "I1", Block2: "I2", Block3: "I3"},
		StyleS: BlockStyles{Block0: "S0", Block1: "S1", Block2: "S2", Block3: "S3"},
		StyleZ: BlockStyles{Block0: "Z0", Block1: "Z1", Block2: "Z2", Block3: "Z3"},
		StyleL: BlockStyles{Block0: "L0", Block1: "L1", Block2: "L2", Block3: "L3"},
		StyleJ: BlockStyles{Block0: "J0", Block1: "J1", Block2: "J2", Block3: "J3"},
		StyleT: BlockStyles{Block0: "T0", Block1: "T1", Block2: "T2", Block3: "T3"},
	})

	tests := []struct {
		name             string
		boardHeight      int
		boardWidth       int
		stationaryBlocks []Block
		expected         bool
		setupTetro       func() Tetro
	}{
		{
			name:             "can move left in empty board",
			boardHeight:      20,
			boardWidth:       10,
			stationaryBlocks: []Block{},
			expected:         true,
			setupTetro:       func() Tetro { return newO() },
		},
		{
			name:             "cannot move left at left edge",
			boardHeight:      20,
			boardWidth:       10,
			stationaryBlocks: []Block{},
			expected:         false,
			setupTetro: func() Tetro {

				Init(Config{
					SpawnX: 0,
					SpawnY: 0,
					StyleO: BlockStyles{Block0: "O0", Block1: "O1", Block2: "O2", Block3: "O3"},
					StyleI: BlockStyles{Block0: "I0", Block1: "I1", Block2: "I2", Block3: "I3"},
					StyleS: BlockStyles{Block0: "S0", Block1: "S1", Block2: "S2", Block3: "S3"},
					StyleZ: BlockStyles{Block0: "Z0", Block1: "Z1", Block2: "Z2", Block3: "Z3"},
					StyleL: BlockStyles{Block0: "L0", Block1: "L1", Block2: "L2", Block3: "L3"},
					StyleJ: BlockStyles{Block0: "J0", Block1: "J1", Block2: "J2", Block3: "J3"},
					StyleT: BlockStyles{Block0: "T0", Block1: "T1", Block2: "T2", Block3: "T3"},
				})
				return newO()
			},
		},
		{
			name:        "cannot move left due to stationary block",
			boardHeight: 20,
			boardWidth:  10,
			stationaryBlocks: []Block{
				{x: 4, y: 0, style: "block"},
			},
			expected:   false,
			setupTetro: func() Tetro { return newO() },
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tetro := tt.setupTetro()
			result := tetro.CanMoveLeft(tt.boardHeight, tt.boardWidth, tt.stationaryBlocks)
			if result != tt.expected {
				t.Errorf("CanMoveLeft() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestClone(t *testing.T) {
	original := tetro{
		x:        5,
		y:        10,
		rotation: 2,
		block0:   Block{x: 5, y: 10, style: "style0"},
		block1:   Block{x: 6, y: 10, style: "style1"},
		block2:   Block{x: 5, y: 11, style: "style2"},
		block3:   Block{x: 6, y: 11, style: "style3"},
	}

	cloned := clone(original)

	if cloned.x != original.x {
		t.Errorf("Clone x = %v, want %v", cloned.x, original.x)
	}
	if cloned.y != original.y {
		t.Errorf("Clone y = %v, want %v", cloned.y, original.y)
	}
	if cloned.rotation != original.rotation {
		t.Errorf("Clone rotation = %v, want %v", cloned.rotation, original.rotation)
	}

	blocks := []struct{ orig, cloned *Block }{
		{&original.block0, &cloned.block0},
		{&original.block1, &cloned.block1},
		{&original.block2, &cloned.block2},
		{&original.block3, &cloned.block3},
	}

	for i, pair := range blocks {
		origX, origY := pair.orig.Coordinates()
		clonedX, clonedY := pair.cloned.Coordinates()

		if clonedX != origX || clonedY != origY {
			t.Errorf("Block %d coordinates not copied correctly: got (%d,%d), want (%d,%d)",
				i, clonedX, clonedY, origX, origY)
		}

		if pair.cloned.Style() != pair.orig.Style() {
			t.Errorf("Block %d style not copied correctly: got %v, want %v",
				i, pair.cloned.Style(), pair.orig.Style())
		}

		pair.cloned.MoveDown()
		origXAfter, origYAfter := pair.orig.Coordinates()
		if origXAfter != origX || origYAfter != origY {
			t.Errorf("Modifying clone affected original block %d", i)
		}
	}
}
