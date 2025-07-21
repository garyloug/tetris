The game package implements the core logic of the game.

It manages the game state, handles user inputs, and sends updates to the UI.

The method `run` is the main loop in the game. On each iteration, it moves the active tetro down, checks for game over conditions, clears any completed lines, updates the game level and speed, and finally updates the UI based on everything it just processed.

A secondary loop, implemented in the method `userEvents`, watches for user inputs and processes them accordingly. It allows the player to move the tetro left, right, down, rotate it, pause the game, or quit.

When a tetro is no longer able to move down, its blocks are added to the game's list of stationary blocks. The next active (falling) tetro is then pulled from the front of the queue, while a brand new random tetro is spawned and added to the end of the queue.

As mentioned above, on each iteration the game will process its list of stationary blocks, checking for completed lines. Completed lines of blocks are removed, and the blocks above are moved down.

The score is updated based on how many lines were cleared by the falling tetro. More lines cleared together gives a better score.

Level is updated based on total cleared lines, and the game speed is adjusted according to the level - the active tetro falls faster as the level increases.

The game ends when the active tetro reaches the top of the board.