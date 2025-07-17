The tetris package defines the tetris shapes.

The basic element in this package is the Block, defined in `block.go`.
A block knows its location (x, y) and its style - a UI specific detail that determines how it looks on screen.
A Block can also be moved up, down, left and right.

Next there is the Tetromino (shortened in code to Tetro), defined as an interface in `tetris.go`.
A Tetro is a collection of 4 Blocks that form a shape. Each shape, implementing the Tetro interface, is defined in its own `shapeX.go` file.
A Tetro knows its 4 blocks, its location (x, y) and its current rotation state.
Like a Block, a Tetro can also be moved up, down, left and right. It can also be rotated.

Common Tetro methods such as `MoveLeft`, `MoveRight`, and `MoveDown` are defined in the `tetris.go` file.
Since each shape has unique rotation logic, the `Rotate` methods are defined in the `shapeX.go` files.
