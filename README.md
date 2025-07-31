# tetris
A just for fun Tetris game written in Go.

The game supports both a windowed desktop UI and a console UI to be run in a terminal.
The UI is abstracted behind a custom interface, so in future it may be possible to plug additional UIs into the same game logic. Possibly a web UI.


Mage Commands:
- `mage build` - Build with Fyne GUI (requires C compiler)
- `mage buildConsole` - Build console-only version (no C compiler dependencies)
- `mage run` - Build and run the game
- `mage test` - Run unit tests
- `mage check` - Run formatting, vetting, linting, and tests
- `mage clean` - Remove built artifacts

Non Mage users can build and run the game with:
```bash
# Default build (requires C compiler)
go build -o tetris ./cmd
./tetris

# Console-only build (no C compiler dependencies)
go build -tags console -o tetris ./cmd
./tetris
```

Enjoy :)

![alt text](./docs/screengrab.png)
