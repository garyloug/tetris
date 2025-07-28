package ui

import (
	"testing"
)

func TestMockUI_Implementation(t *testing.T) {
	mockUI, cleanup := newMockUI()
	defer cleanup()

	// check mockUI implements UI interface
	var _ UI = mockUI

	// assert for non-UI field access
	mock := mockUI.(*MockUI)

	err := mock.Init(20, 10)
	if err != nil {
		t.Errorf("mockUI.Init() error = %v, want nil", err)
	}
	if mock.BoardHeight != 20 {
		t.Errorf("mockUI.boardHeight = %v, want %v", mock.BoardHeight, 20)
	}
	if mock.BoardWidth != 10 {
		t.Errorf("mockUI.boardWidth = %v, want %v", mock.BoardWidth, 10)
	}

	o, i, s, z, l, j, tStyle := mock.GetBlockStyles()
	expectedStyle := "mockStyle"
	if o != expectedStyle {
		t.Errorf("GetBlockStyles() o = %v, want %v", o, expectedStyle)
	}
	if i != expectedStyle {
		t.Errorf("GetBlockStyles() i = %v, want %v", i, expectedStyle)
	}
	if s != expectedStyle {
		t.Errorf("GetBlockStyles() s = %v, want %v", s, expectedStyle)
	}
	if z != expectedStyle {
		t.Errorf("GetBlockStyles() z = %v, want %v", z, expectedStyle)
	}
	if l != expectedStyle {
		t.Errorf("GetBlockStyles() l = %v, want %v", l, expectedStyle)
	}
	if j != expectedStyle {
		t.Errorf("GetBlockStyles() j = %v, want %v", j, expectedStyle)
	}
	if tStyle != expectedStyle {
		t.Errorf("GetBlockStyles() tStyle = %v, want %v", tStyle, expectedStyle)
	}

	keyChannel := mock.KeyPress()
	if keyChannel == nil {
		t.Error("mockUI.KeyPress() returned nil channel")
	}

	mock.Start()
	if !mock.Started {
		t.Error("mockUI.Start() did not set started flag")
	}

	mock.SendKeyPress(KeyDown)
	select {
	case key := <-keyChannel:
		if key != KeyDown {
			t.Errorf("Received key = %v, want %v", key, KeyDown)
		}
	default:
		t.Error("No key received on channel")
	}

	mock.Stop()
	if !mock.Stopped {
		t.Error("mockUI.Stop() did not set stopped flag")
	}
}
