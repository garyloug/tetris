//go:build console

package main

import "github.com/garyloug/tetris/pkg/ui"

func getDefaultUIType() ui.UiType {
	return ui.Tcell
}
