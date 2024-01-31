# Go Page Replacement Algorithms

This is a Go program that implements four different page replacement algorithms: FIFO, LRU, OPT, and NFU.

## Files

- `main.go`: This is the main file that contains the implementation of the four page replacement algorithms.

## Algorithms

- `FIFO`: This function implements the First-In-First-Out page replacement algorithm.
- `LRU`: This function implements the Least Recently Used page replacement algorithm.
- `OPT`: This function implements the Optimal page replacement algorithm.
- `NFU`: This function implements the Not Frequently Used page replacement algorithm.

## Usage

1. Run the program: `go run main.go`
2. When prompted, enter the name of the file that contains the page numbers and the number of frames.

## Input File Format

The input file should contain two lines:
- The first line should contain the page numbers, separated by spaces.
- The second line should contain a single number, the number of frames.

## Output

The program will print the number of page replacements and the number of page hits for each algorithm.

## Note

This program assumes that the input file is correctly formatted and that the page numbers and number of frames are valid.
