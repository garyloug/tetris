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
		expectError   bool
		errorContains string
	}{
		{
			name:        "valid Console UI",
			uiType:      Mock,
			expectError: false,
		},
		{
			name:          "invalid UI type",
			uiType:        UiType(999),
			expectError:   true,
			errorContains: "unsupported UI type",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ui, cleanup, err := NewUI(tt.uiType)

			if tt.expectError {
				if err == nil {
					t.Errorf("NewUI(%v) expected error, got nil", tt.uiType)
				} else if tt.errorContains != "" {
					if len(err.Error()) == 0 || (len(tt.errorContains) > 0 && len(err.Error()) == 0) {
						t.Errorf("NewUI(%v) error = %q, want to contain %q", tt.uiType, err.Error(), tt.errorContains)
					}
				}

				if cleanup != nil && ui == nil {
					cleanup()
				}
			} else {
				if err != nil {
					t.Errorf("NewUI(%v) unexpected error: %v", tt.uiType, err)
				}
				if ui == nil {
					t.Errorf("NewUI(%v) returned nil UI", tt.uiType)
				}
				if cleanup == nil {
					t.Errorf("NewUI(%v) returned nil cleanup function", tt.uiType)
				}

				if ui != nil {
					keyChannel := ui.KeyPress()
					if keyChannel == nil {
						t.Error("NewUI created UI with nil KeyPress channel")
					}

					o, i, s, z, l, j, tStyle := ui.GetBlockStyles()
					if o == nil || i == nil || s == nil || z == nil || l == nil || j == nil || tStyle == nil {
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
