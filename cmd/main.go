package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/garyloug/tetris/pkg/game"
	"github.com/garyloug/tetris/pkg/ui"
)

func main() {
	devMode := flag.Bool("dev", false, "Run with a development UI.")
	flag.Parse()

	var uiType ui.UiType
	if *devMode {
		uiType = ui.ConsoleDev
	} else {
		uiType = ui.Console
	}

	uiInstance, cleanup, err := ui.NewUI(uiType)
	if err != nil {
		panic(fmt.Sprintf("Failed to create UI: %v", err))
	}
	defer cleanup()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	game := game.NewGame(uiInstance)
	gameOver := game.Start()

	select {
	case <-gameOver:
	case <-sigChan:
	}
}
