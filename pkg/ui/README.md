The ui package creates a layer of abstraction between the game logic and the UI.

The idea is the game logic can potentially be reused with different UIs, such as a console, windowed application, or even a web UI.
The initial implementation is a console UI using the brilliant [tcell](https://github.com/gdamore/tcell) library. It's a library that I'm using for the first time so implementation may not be perfect.

The UI interface is defined in `ui.go`. The UI is responsible for updating the game on screen as well as handelling user input.
In the initial console implementation, tcell is able to handle both of these tasks, but in future implementations these responsibilities may need to be separated.

Each UI implementation is also responsible for defining the styles for each Tetro type. These styles are provided to the game logic via the `GetBlockStyles` method.
The game logic then initializes the Tetro package with these styles, where the appropriate style is stored in each individual block.
These styles (how the block looks on screen) are then provided, along with the block coordinates, each time the screen is updated.
