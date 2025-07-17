package tetris

import (
	"testing"
)

func TestNewO(t *testing.T) {
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

	if o == nil {
		t.Fatal("newO() returned nil")
	}

	blocks := o.Blocks()
	if len(blocks) != 4 {
		t.Errorf("newO() should have 4 blocks, got %d", len(blocks))
	}

	expectedInitialCoords := []struct{ x, y int }{
		{5, 0}, // block0: spawn position
		{5, 1}, // block1: y + 1
		{6, 1}, // block2: x + 1, y + 1
		{6, 0}, // block3: x + 1
	}

	for i, expected := range expectedInitialCoords {
		x, y := blocks[i].Coordinates()
		if x != expected.x || y != expected.y {
			t.Errorf("Initial O shape: Block %d at (%d,%d), want (%d,%d)", i, x, y, expected.x, expected.y)
		}
	}

	expectedStyles := []any{config.StyleO.Block0, config.StyleO.Block1, config.StyleO.Block2, config.StyleO.Block3}
	for i, expected := range expectedStyles {
		if blocks[i].Style() != expected {
			t.Errorf("Block %d style = %v, want %v", i, blocks[i].Style(), expected)
		}
	}
}

func TestO_Rotate(t *testing.T) {
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
			o := newO()
			for r := 0; r < tt.rotation; r++ {
				o.Rotate()
			}

			blocks := o.Blocks()
			if len(blocks) != 4 {
				t.Errorf("O piece should have 4 blocks after %d rotations, got %d", tt.rotation, len(blocks))
			}

			expectedCoords := [][]struct{ x, y int }{
				{{5, 0}, {5, 1}, {6, 1}, {6, 0}}, // initial state
				{{5, 1}, {6, 1}, {6, 0}, {5, 0}}, // 1 rotation
				{{6, 1}, {6, 0}, {5, 0}, {5, 1}}, // 2 rotations
				{{6, 0}, {5, 0}, {5, 1}, {6, 1}}, // 3 rotations
				{{5, 0}, {5, 1}, {6, 1}, {6, 0}}, // 4 rotations (back to original)
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

func TestO_CanRotate(t *testing.T) {
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
		setupO           func() Tetro
	}{
		{
			name:             "can rotate in empty board",
			boardHeight:      20,
			boardWidth:       10,
			stationaryBlocks: []Block{},
			expected:         true,
			setupO:           func() Tetro { return newO() },
		},
		{
			name:             "can rotate at edges (O shape fits in 2x2)",
			boardHeight:      20,
			boardWidth:       10,
			stationaryBlocks: []Block{},
			expected:         true,
			setupO: func() Tetro {
				Init(Config{
					SpawnX: 8,
					SpawnY: 18,
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
			name:        "can rotate with stationary blocks that don't interfere",
			boardHeight: 20,
			boardWidth:  10,
			stationaryBlocks: []Block{
				{x: 0, y: 0, style: "block"},
			},
			expected: true,
			setupO:   func() Tetro { return newO() },
		},
		{
			name:        "can always rotate - O piece stays in same 2x2 grid",
			boardHeight: 20,
			boardWidth:  10,
			stationaryBlocks: []Block{
				{x: 3, y: 5, style: "block"},
				{x: 7, y: 5, style: "block"},
				{x: 5, y: 3, style: "block"},
				{x: 5, y: 7, style: "block"},
			},
			expected: true,
			setupO:   func() Tetro { return newO() },
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := tt.setupO()
			result := o.CanRotate(tt.boardHeight, tt.boardWidth, tt.stationaryBlocks)
			if result != tt.expected {
				t.Errorf("CanRotate() = %v, want %v", result, tt.expected)
			}
		})
	}
}
