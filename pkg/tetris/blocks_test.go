package tetris

import (
	"testing"
)

func TestBlock_Coordinates(t *testing.T) {
	tests := []struct {
		name      string
		block     Block
		expectedX int
		expectedY int
	}{
		{
			name:      "origin block",
			block:     Block{x: 0, y: 0, style: "test"},
			expectedX: 0,
			expectedY: 0,
		},
		{
			name:      "positive coordinates",
			block:     Block{x: 5, y: 10, style: "test"},
			expectedX: 5,
			expectedY: 10,
		},
		{
			name:      "negative coordinates",
			block:     Block{x: -3, y: -7, style: "test"},
			expectedX: -3,
			expectedY: -7,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x, y := tt.block.Coordinates()
			if x != tt.expectedX {
				t.Errorf("Coordinates() x = %v, want %v", x, tt.expectedX)
			}
			if y != tt.expectedY {
				t.Errorf("Coordinates() y = %v, want %v", y, tt.expectedY)
			}
		})
	}
}

func TestBlock_Style(t *testing.T) {
	tests := []struct {
		name     string
		block    Block
		expected any
	}{
		{
			name:     "string style",
			block:    Block{x: 0, y: 0, style: "red"},
			expected: "red",
		},
		{
			name:     "int style",
			block:    Block{x: 0, y: 0, style: 42},
			expected: 42,
		},
		{
			name:     "nil style",
			block:    Block{x: 0, y: 0, style: nil},
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			style := tt.block.Style()
			if style != tt.expected {
				t.Errorf("Style() = %v, want %v", style, tt.expected)
			}
		})
	}
}

func TestBlock_MoveDown(t *testing.T) {
	tests := []struct {
		name      string
		initial   Block
		expectedY int
	}{
		{
			name:      "move from origin",
			initial:   Block{x: 0, y: 0, style: "test"},
			expectedY: 1,
		},
		{
			name:      "move from positive y",
			initial:   Block{x: 5, y: 10, style: "test"},
			expectedY: 11,
		},
		{
			name:      "move from negative y",
			initial:   Block{x: 0, y: -5, style: "test"},
			expectedY: -4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			block := tt.initial
			originalX := block.x

			block.MoveDown()

			if block.y != tt.expectedY {
				t.Errorf("MoveDown() y = %v, want %v", block.y, tt.expectedY)
			}
			if block.x != originalX {
				t.Errorf("MoveDown() should not change x coordinate, got %v, want %v", block.x, originalX)
			}
		})
	}
}

func TestBlock_moveRight(t *testing.T) {
	tests := []struct {
		name      string
		initial   Block
		expectedX int
	}{
		{
			name:      "move from origin",
			initial:   Block{x: 0, y: 0, style: "test"},
			expectedX: 1,
		},
		{
			name:      "move from positive x",
			initial:   Block{x: 5, y: 10, style: "test"},
			expectedX: 6,
		},
		{
			name:      "move from negative x",
			initial:   Block{x: -3, y: 0, style: "test"},
			expectedX: -2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			block := tt.initial
			originalY := block.y

			block.moveRight()

			if block.x != tt.expectedX {
				t.Errorf("moveRight() x = %v, want %v", block.x, tt.expectedX)
			}
			if block.y != originalY {
				t.Errorf("moveRight() should not change y coordinate, got %v, want %v", block.y, originalY)
			}
		})
	}
}

func TestBlock_moveLeft(t *testing.T) {
	tests := []struct {
		name      string
		initial   Block
		expectedX int
	}{
		{
			name:      "move from origin",
			initial:   Block{x: 0, y: 0, style: "test"},
			expectedX: -1,
		},
		{
			name:      "move from positive x",
			initial:   Block{x: 5, y: 10, style: "test"},
			expectedX: 4,
		},
		{
			name:      "move from negative x",
			initial:   Block{x: -3, y: 0, style: "test"},
			expectedX: -4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			block := tt.initial
			originalY := block.y

			block.moveLeft()

			if block.x != tt.expectedX {
				t.Errorf("moveLeft() x = %v, want %v", block.x, tt.expectedX)
			}
			if block.y != originalY {
				t.Errorf("moveLeft() should not change y coordinate, got %v, want %v", block.y, originalY)
			}
		})
	}
}

func TestBlock_moveUp(t *testing.T) {
	tests := []struct {
		name      string
		initial   Block
		expectedY int
	}{
		{
			name:      "move from origin",
			initial:   Block{x: 0, y: 0, style: "test"},
			expectedY: -1,
		},
		{
			name:      "move from positive y",
			initial:   Block{x: 5, y: 10, style: "test"},
			expectedY: 9,
		},
		{
			name:      "move from negative y",
			initial:   Block{x: 0, y: -5, style: "test"},
			expectedY: -6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			block := tt.initial
			originalX := block.x

			block.moveUp()

			if block.y != tt.expectedY {
				t.Errorf("moveUp() y = %v, want %v", block.y, tt.expectedY)
			}
			if block.x != originalX {
				t.Errorf("moveUp() should not change x coordinate, got %v, want %v", block.x, originalX)
			}
		})
	}
}

func TestBlock_MultipleMoves(t *testing.T) {
	block := Block{x: 5, y: 5, style: "test"}

	block.MoveDown()
	block.MoveDown()
	block.moveRight()
	block.moveLeft()
	block.moveUp()

	expectedX, expectedY := 5, 6 // right+left=0, down+down+up=1
	actualX, actualY := block.Coordinates()

	if actualX != expectedX {
		t.Errorf("After multiple moves, x = %v, want %v", actualX, expectedX)
	}
	if actualY != expectedY {
		t.Errorf("After multiple moves, y = %v, want %v", actualY, expectedY)
	}
}
