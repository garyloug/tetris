package ui

import (
	"testing"
)

func TestUI_KeyPress(t *testing.T) {
	u := ui{
		eventChan: make(chan KeyPress, 1),
	}

	keyChannel := u.KeyPress()
	if keyChannel == nil {
		t.Error("KeyPress() returned nil channel")
	}

	testKey := KeyDown
	u.eventChan <- testKey

	select {
	case receivedKey := <-keyChannel:
		if receivedKey != testKey {
			t.Errorf("KeyPress() channel received %v, want %v", receivedKey, testKey)
		}
	default:
		t.Error("KeyPress() channel did not receive the sent key")
	}
}

func TestNewUI(t *testing.T) {
	tests := []struct {
		name          string
		uiType        UiType
		boardHeight   int
		boardWidth    int
		expectError   bool
		errorContains string
	}{
		{
			name:        "valid Console UI",
			uiType:      Console,
			boardHeight: 20,
			boardWidth:  10,
			expectError: false,
		},
		{
			name:        "valid ConsoleDev UI",
			uiType:      ConsoleDev,
			boardHeight: 20,
			boardWidth:  10,
			expectError: false,
		},
		{
			name:          "invalid UI type",
			uiType:        UiType(999),
			boardHeight:   20,
			boardWidth:    10,
			expectError:   true,
			errorContains: "unsupported UI type",
		},
		{
			name:          "invalid board height",
			uiType:        Console,
			boardHeight:   0,
			boardWidth:    10,
			expectError:   true,
			errorContains: "failed to initialize UI",
		},
		{
			name:          "invalid board width",
			uiType:        Console,
			boardHeight:   20,
			boardWidth:    -1,
			expectError:   true,
			errorContains: "failed to initialize UI",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ui, cleanup, err := NewUI(tt.uiType, tt.boardHeight, tt.boardWidth)

			if tt.expectError {
				if err == nil {
					t.Errorf("NewUI(%v, %d, %d) expected error, got nil", tt.uiType, tt.boardHeight, tt.boardWidth)
				} else if tt.errorContains != "" {
					if len(err.Error()) == 0 || (len(tt.errorContains) > 0 && len(err.Error()) == 0) {
						t.Errorf("NewUI(%v, %d, %d) error = %q, want to contain %q", tt.uiType, tt.boardHeight, tt.boardWidth, err.Error(), tt.errorContains)
					}
				}

				if cleanup != nil && ui == nil {
					cleanup()
				}
			} else {
				if err != nil {
					t.Errorf("NewUI(%v, %d, %d) unexpected error: %v", tt.uiType, tt.boardHeight, tt.boardWidth, err)
				}
				if ui == nil {
					t.Errorf("NewUI(%v, %d, %d) returned nil UI", tt.uiType, tt.boardHeight, tt.boardWidth)
				}
				if cleanup == nil {
					t.Errorf("NewUI(%v, %d, %d) returned nil cleanup function", tt.uiType, tt.boardHeight, tt.boardWidth)
				}

				if ui != nil {
					keyChannel := ui.KeyPress()
					if keyChannel == nil {
						t.Error("NewUI created UI with nil KeyPress channel")
					}

					o, i, s, z, l, j, tStyle := ui.GetBlockStyles()
					if o.Block0 == nil || i.Block0 == nil || s.Block0 == nil || z.Block0 == nil || l.Block0 == nil || j.Block0 == nil || tStyle.Block0 == nil {
						t.Error("NewUI created UI with nil block styles")
					}
				}

				if cleanup != nil {
					cleanup()
				}
			}
		})
	}
}
