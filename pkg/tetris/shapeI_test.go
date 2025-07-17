package tetris

import (
	"testing"
)

func TestNewI(t *testing.T) {
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

	i := newI()

	if i == nil {
		t.Fatal("newI() returned nil")
	}

	blocks := i.Blocks()
	if len(blocks) != 4 {
		t.Errorf("newI() should have 4 blocks, got %d", len(blocks))
	}

	expectedInitialCoords := []struct{ x, y int }{
		{5, 0}, // block0: spawn position
		{6, 0}, // block1: x + 1
		{7, 0}, // block2: x + 2
		{4, 0}, // block3: x - 1
	}

	for i, expected := range expectedInitialCoords {
		x, y := blocks[i].Coordinates()
		if x != expected.x || y != expected.y {
			t.Errorf("Initial I shape: Block %d at (%d,%d), want (%d,%d)", i, x, y, expected.x, expected.y)
		}
	}

	expectedStyles := []any{config.StyleI.Block0, config.StyleI.Block1, config.StyleI.Block2, config.StyleI.Block3}
	for i, expected := range expectedStyles {
		if blocks[i].Style() != expected {
			t.Errorf("Block %d style = %v, want %v", i, blocks[i].Style(), expected)
		}
	}
}

func TestI_Rotate(t *testing.T) {
	Init(Config{
		SpawnX: 5,
		SpawnY: 5,
		StyleO: BlockStyles{Block0: "O0", Block1: "O1", Block2: "O2", Block3: "O3"},
		StyleI: BlockStyles{Block0: "I0", Block1: "I1", Block2: "I2", Block3: "I3"},
		StyleS: BlockStyles{Block0: "S0", Block1: "S1", Block2: "S2", Block3: "S3"},
		StyleZ: BlockStyles{Block0: "Z0", Block1: "Z1", Block2: "Z2", Block3: "Z3"},
		StyleL: BlockStyles{Block0: "L0", Block1: "L1", Block2: "L2", Block3: "L3"},
		StyleJ: BlockStyles{Block0: "J0", Block1: "J1", Block2: "J2", Block3: "J3"},
		StyleT: BlockStyles{Block0: "T0", Block1: "T1", Block2: "T2", Block3: "T3"},
	})

	tests := []struct {
		name     string
		rotation int
	}{
		{"rotation 0", 0},
		{"rotation 1", 1},
		{"rotation 2", 2},
		{"rotation 3", 3},
		{"rotation 4", 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := newI()

			for r := 0; r < tt.rotation; r++ {
				i.Rotate()
			}

			blocks := i.Blocks()
			if len(blocks) != 4 {
				t.Errorf("I piece should always have 4 blocks, got %d", len(blocks))
			}

			expectedCoords := [][]struct{ x, y int }{
				{{5, 5}, {6, 5}, {7, 5}, {4, 5}}, // initial state
				{{5, 5}, {5, 4}, {5, 3}, {5, 6}}, // 1 rotation
				{{5, 5}, {4, 5}, {3, 5}, {6, 5}}, // 2 rotations
				{{5, 5}, {5, 6}, {5, 7}, {5, 4}}, // 3 rotations
				{{5, 5}, {6, 5}, {7, 5}, {4, 5}}, // 4 rotations (back to initial)
			}

			expected := expectedCoords[tt.rotation]
			for i, block := range blocks {
				x, y := block.Coordinates()
				if x != expected[i].x || y != expected[i].y {
					t.Errorf("Rotation %d, Block %d: got (%d,%d), want (%d,%d)", tt.rotation, i, x, y, expected[i].x, expected[i].y)
				}
			}
		})
	}
}

func TestI_CanRotate(t *testing.T) {
	Init(Config{
		SpawnX: 5,
		SpawnY: 5,
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
		setupI           func() Tetro
	}{
		{
			name:             "can rotate in empty board",
			boardHeight:      20,
			boardWidth:       10,
			stationaryBlocks: []Block{},
			expected:         true,
			setupI:           func() Tetro { return newI() },
		},
		{
			name:             "can rotate at left edge (implementation allows)",
			boardHeight:      20,
			boardWidth:       10,
			stationaryBlocks: []Block{},
			expected:         true,
			setupI: func() Tetro {
				Init(Config{
					SpawnX: 1,
					SpawnY: 5,
					StyleI: BlockStyles{Block0: "I0", Block1: "I1", Block2: "I2", Block3: "I3"},
					StyleO: BlockStyles{Block0: "O0", Block1: "O1", Block2: "O2", Block3: "O3"},
					StyleS: BlockStyles{Block0: "S0", Block1: "S1", Block2: "S2", Block3: "S3"},
					StyleZ: BlockStyles{Block0: "Z0", Block1: "Z1", Block2: "Z2", Block3: "Z3"},
					StyleL: BlockStyles{Block0: "L0", Block1: "L1", Block2: "L2", Block3: "L3"},
					StyleJ: BlockStyles{Block0: "J0", Block1: "J1", Block2: "J2", Block3: "J3"},
					StyleT: BlockStyles{Block0: "T0", Block1: "T1", Block2: "T2", Block3: "T3"},
				})
				return newI()
			},
		},
		{
			name:             "can rotate at top edge (implementation allows)",
			boardHeight:      20,
			boardWidth:       10,
			stationaryBlocks: []Block{},
			expected:         true,
			setupI: func() Tetro {
				Init(Config{
					SpawnX: 5,
					SpawnY: 1,
					StyleI: BlockStyles{Block0: "I0", Block1: "I1", Block2: "I2", Block3: "I3"},
					StyleO: BlockStyles{Block0: "O0", Block1: "O1", Block2: "O2", Block3: "O3"},
					StyleS: BlockStyles{Block0: "S0", Block1: "S1", Block2: "S2", Block3: "S3"},
					StyleZ: BlockStyles{Block0: "Z0", Block1: "Z1", Block2: "Z2", Block3: "Z3"},
					StyleL: BlockStyles{Block0: "L0", Block1: "L1", Block2: "L2", Block3: "L3"},
					StyleJ: BlockStyles{Block0: "J0", Block1: "J1", Block2: "J2", Block3: "J3"},
					StyleT: BlockStyles{Block0: "T0", Block1: "T1", Block2: "T2", Block3: "T3"},
				})
				return newI()
			},
		},
		{
			name:        "can rotate with stationary blocks that don't interfere",
			boardHeight: 20,
			boardWidth:  10,
			stationaryBlocks: []Block{
				{x: 0, y: 0, style: "block"},
			},
			expected: true,
			setupI:   func() Tetro { return newI() },
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := tt.setupI()
			result := i.CanRotate(tt.boardHeight, tt.boardWidth, tt.stationaryBlocks)
			if result != tt.expected {
				t.Errorf("CanRotate() = %v, want %v", result, tt.expected)
			}
		})
	}
}
