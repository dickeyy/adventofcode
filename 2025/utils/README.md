# 2025 Utils

A collection of utilities I use for solving AoC problems.

## math.go

- `SumNums` - Sum a slice of integers
- `Abs` - Absolute value of an integer
- `Factors` - Get all factors of an integer
- `GCD` - Get the greatest common divisor of two integers
- `LCM` - Get the least common multiple of two integers

## strings.go

- `GetIntsInString` - Get all integers in a string (returns a slice of ints)
- `AtoiNoErr` - Wrapper of `strconv.Atoi` that just discards the error return
- `IntArrayToString` - Convert a slice of integers to a string (with a separator)

## grid.go

This is a new one this year (2025). I debated about adding it as it does reduce a lot of boilerplate, almost too much, but I decided if I wrote it once I might as well be able to use it on multiple days.

- `Grid` - A 2D grid of bytes
- `Point` - A coordinate in the grid
- `CardinalDirections` - The 4 cardinal directions (no diagonals)
- `AllDirections` - All 8 directions
- `ParseGrid` - Create a grid from a string input
- `IsInBounds` - Check if a point is within the grid boundaries
- `Get` - Get the byte at a given position
- `Set` - Set the byte at a given position
- `Add` - Add two points
- `Neighbors` - Get all valid neighboring points in given directions
- `CardinalNeighbors` - Get the 4 cardinal direction neighbors
- `AllNeighbors` - Get all 8 direction neighbors
- `FindAll` - Find all positions where the cell matches the given value
- `FindFirst` - Find the first position where the cell matches a value
- `String` - Get a string representation of the grid
- `Clone` - Get a deep copy of the grid
- `ManhattanDistance` - Get the Manhattan distance between two points

## debug.go

- `Debug` - Pretty print any data type

## read-file.go

- `ReadFile` - Reads the input file and returns a string

## cli.go

Multiple CLI utils for flag parsing and timing.

- `Output` - This is the only one that matters. It prints the output of the solution in a pretty format.
