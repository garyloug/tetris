package tetris

import (
	"testing"
)

func TestNewZ(t *testing.T) {
	Init(Config{
		SpawnX: 5,
		SpawnY: 5,
		StyleZ: BlockStyles{Block0: "Z0", Block1: "Z1", Block2: "Z2", Block3: "Z3"},
		StyleI: BlockStyles{Block0: "I0", Block1: "I1", Block2: "I2", Block3: "I3"},
		StyleO: BlockStyles{Block0: "O0", Block1: "O1", Block2: "O2", Block3: "O3"},
		StyleS: BlockStyles{Block0: "S0", Block1: "S1", Block2: "S2", Block3: "S3"},
		StyleJ: BlockStyles{Block0: "J0", Block1: "J1", Block2: "J2", Block3: "J3"},
		StyleL: BlockStyles{Block0: "L0", Block1: "L1", Block2: "L2", Block3: "L3"},
		StyleT: BlockStyles{Block0: "T0", Block1: "T1", Block2: "T2", Block3: "T3"},
	})

	z := newZ()

	if z == nil {
		t.Fatal("newZ() returned nil")
	}

	blocks := z.Blocks()
	if len(blocks) != 4 {
		t.Errorf("newZ() should have 4 blocks, got %d", len(blocks))
	}

	expectedInitialCoords := []struct{ x, y int }{
		{5, 5}, // block0: spawn position
		{4, 5}, // block1: x - 1
		{5, 6}, // block2: y + 1
		{6, 6}, // block3: x + 1, y + 1
	}

	for i, expected := range expectedInitialCoords {
		x, y := blocks[i].Coordinates()
		if x != expected.x || y != expected.y {
			t.Errorf("Initial Z shape: Block %d at (%d,%d), want (%d,%d)", i, x, y, expected.x, expected.y)
		}
	}

	expectedStyles := []any{config.StyleZ.Block0, config.StyleZ.Block1, config.StyleZ.Block2, config.StyleZ.Block3}
	for i, expected := range expectedStyles {
		if blocks[i].Style() != expected {
			t.Errorf("Block %d style = %v, want %v", i, blocks[i].Style(), expected)
		}
	}
}

func TestZ_Rotate(t *testing.T) {
	Init(Config{
		SpawnX: 5,
		SpawnY: 5,
		StyleZ: BlockStyles{Block0: "Z0", Block1: "Z1", Block2: "Z2", Block3: "Z3"},
		StyleI: BlockStyles{Block0: "I0", Block1: "I1", Block2: "I2", Block3: "I3"},
		StyleO: BlockStyles{Block0: "O0", Block1: "O1", Block2: "O2", Block3: "O3"},
		StyleS: BlockStyles{Block0: "S0", Block1: "S1", Block2: "S2", Block3: "S3"},
		StyleJ: BlockStyles{Block0: "J0", Block1: "J1", Block2: "J2", Block3: "J3"},
		StyleL: BlockStyles{Block0: "L0", Block1: "L1", Block2: "L2", Block3: "L3"},
		StyleT: BlockStyles{Block0: "T0", Block1: "T1", Block2: "T2", Block3: "T3"},
	})

	z := newZ()

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
			z = newZ()
			for r := 0; r < tt.rotation; r++ {
				z.Rotate()
			}

			blocks := z.Blocks()
			if len(blocks) != 4 {
				t.Errorf("Z piece should have 4 blocks after %d rotations, got %d", tt.rotation, len(blocks))
			}

			expectedCoords := [][]struct{ x, y int }{
				{{5, 5}, {4, 5}, {5, 6}, {6, 6}}, // rotation 0
				{{5, 5}, {5, 6}, {6, 5}, {6, 4}}, // rotation 1
				{{5, 5}, {6, 5}, {5, 4}, {4, 4}}, // rotation 2
				{{5, 5}, {5, 4}, {4, 5}, {4, 6}}, // rotation 3
				{{5, 5}, {4, 5}, {5, 6}, {6, 6}}, // rotation 4 (back to original)
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

func TestZ_CanRotate(t *testing.T) {
	Init(Config{
		SpawnX: 5,
		SpawnY: 5,
		StyleZ: BlockStyles{Block0: "Z0", Block1: "Z1", Block2: "Z2", Block3: "Z3"},
		StyleI: BlockStyles{Block0: "I0", Block1: "I1", Block2: "I2", Block3: "I3"},
		StyleO: BlockStyles{Block0: "O0", Block1: "O1", Block2: "O2", Block3: "O3"},
		StyleS: BlockStyles{Block0: "S0", Block1: "S1", Block2: "S2", Block3: "S3"},
		StyleJ: BlockStyles{Block0: "J0", Block1: "J1", Block2: "J2", Block3: "J3"},
		StyleL: BlockStyles{Block0: "L0", Block1: "L1", Block2: "L2", Block3: "L3"},
		StyleT: BlockStyles{Block0: "T0", Block1: "T1", Block2: "T2", Block3: "T3"},
	})

	tests := []struct {
		name             string
		boardHeight      int
		boardWidth       int
		stationaryBlocks []Block
		expected         bool
		setupZ           func() Tetro
	}{
		{
			name:             "can rotate in empty board",
			boardHeight:      20,
			boardWidth:       10,
			stationaryBlocks: []Block{},
			expected:         true,
			setupZ:           func() Tetro { return newZ() },
		},
		{
			name:             "can rotate with space",
			boardHeight:      20,
			boardWidth:       10,
			stationaryBlocks: []Block{},
			expected:         true,
			setupZ: func() Tetro {
				z := newZ()
				z.MoveRight()
				return z
			},
		},
		{
			name:        "can rotate with distant stationary blocks",
			boardHeight: 20,
			boardWidth:  10,
			stationaryBlocks: []Block{
				{x: 0, y: 0, style: "block"},
			},
			expected: true,
			setupZ:   func() Tetro { return newZ() },
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			z := tt.setupZ()
			result := z.CanRotate(tt.boardHeight, tt.boardWidth, tt.stationaryBlocks)
			if result != tt.expected {
				t.Errorf("CanRotate() = %v, want %v", result, tt.expected)
			}
		})
	}
}
