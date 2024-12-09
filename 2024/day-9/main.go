package main

import (
	"github.com/dickeyy/adventofcode/2024/utils"
)

func main() {
	utils.ParseFlags()
	p := utils.GetPart()

	i := utils.ReadFile("../../inputs/2024/day-9/input.txt")
	utils.Output(day9(i, p))
}

// file represents a file on the disk with its ID and size
type File struct {
	ID   int
	Size int
}

// disk represents the current state of the disk
type Disk struct {
	Blocks []int // -1 represents free space, non-negative numbers represent file IDs
}

func day9(input string, part int) int {
	return compactDisk(input, part)
}

// parseDiskMap coverts the input string into a sequence of files and free spaces
func parseDiskMap(input string) ([]int, []int) {
	files := make([]int, 0)
	spaces := make([]int, 0)

	for i := 0; i < len(input); i++ {
		size := utils.AtoiNoErr(string(input[i]))
		if i%2 == 0 {
			files = append(files, size)
		} else {
			spaces = append(spaces, size)
		}
	}

	return files, spaces
}

// createInitialDisk creates the initial disk state with files and free spaces
func createInitialDisk(files, spaces []int) Disk {
	var blocks []int
	fileID := 0

	for i := 0; i < len(files); i++ {
		// add file blocks
		for j := 0; j < files[i]; j++ {
			blocks = append(blocks, fileID)
		}
		fileID++

		// add free space
		if i < len(spaces) {
			for j := 0; j < spaces[i]; j++ {
				blocks = append(blocks, -1)
			}
		}
	}

	return Disk{Blocks: blocks}
}

// findFirstFreeSpace finds the leftmost free space in the disk
func (d *Disk) findFirstFreeSpace() int {
	for i, block := range d.Blocks {
		if block == -1 {
			return i
		}
	}
	return -1
}

// findLastFileBlock finds the position of the rightmost block of any file
func (d *Disk) findLastFileBlock() int {
	for i := len(d.Blocks) - 1; i >= 0; i-- {
		if d.Blocks[i] != -1 {
			return i
		}
	}
	return -1
}

// moveOneBlock moves a single block from source to destination
func (d *Disk) moveOneBlock(fromIndex, toIndex int) {
	fileID := d.Blocks[fromIndex]
	d.Blocks[fromIndex] = -1
	d.Blocks[toIndex] = fileID
}

func (d *Disk) calculateChecksum() int {
	cs := 0
	for pos, fileID := range d.Blocks {
		if fileID != -1 {
			cs += pos * fileID
		}
	}
	return cs
}

// findFileSize finds the size of a file given any position within it
func (d *Disk) findFileSize(pos int) int {
	if pos < 0 || pos >= len(d.Blocks) || d.Blocks[pos] == -1 {
		return 0
	}

	fileID := d.Blocks[pos]
	start := pos
	for start >= 0 && d.Blocks[start] == fileID {
		start--
	}
	start++

	end := pos
	for end < len(d.Blocks) && d.Blocks[end] == fileID {
		end++
	}

	return end - start
}

// findFileStart finds the starting position of a file given a position within it
func (d *Disk) findFileStart(pos int) int {
	if pos < 0 || pos >= len(d.Blocks) || d.Blocks[pos] == -1 {
		return -1
	}

	fileID := d.Blocks[pos]
	start := pos
	for start >= 0 && d.Blocks[start] == fileID {
		start--
	}
	return start + 1
}

// findFreeSpaceSize finds size of continuous free space starting at a position
func (d *Disk) findFreeSpaceSize(pos int) int {
	size := 0
	for i := pos; i < len(d.Blocks) && d.Blocks[i] == -1; i++ {
		size++
	}
	return size
}

// moveWholeFile moves an entire file from source to destination
func (d *Disk) moveWholeFile(fromIndex, toIndex, size int) {
	fileID := d.Blocks[fromIndex]

	// Clear old location
	for i := 0; i < size; i++ {
		d.Blocks[fromIndex+i] = -1
	}

	// Place at new location
	for i := 0; i < size; i++ {
		d.Blocks[toIndex+i] = fileID
	}
}

// compactDisk performs the disk compaction process
func compactDisk(diskMap string, part int) int {
	// Parse the input
	files, spaces := parseDiskMap(diskMap)
	disk := createInitialDisk(files, spaces)

	if part == 1 {
		// Part 1: Original block-by-block movement
		for {
			freeSpace := disk.findFirstFreeSpace()
			if freeSpace == -1 {
				break // No more free space
			}

			lastBlock := disk.findLastFileBlock()
			if lastBlock == -1 || lastBlock <= freeSpace {
				break // No more blocks to move
			}

			// Move one block at a time
			disk.moveOneBlock(lastBlock, freeSpace)
		}
	} else {
		// Part 2: Move whole files in decreasing file ID order
		maxFileID := -1
		for _, block := range disk.Blocks {
			if block > maxFileID {
				maxFileID = block
			}
		}

		// Process files in decreasing order
		for fileID := maxFileID; fileID >= 0; fileID-- {
			// Find this file's position
			var filePos int = -1
			for i := len(disk.Blocks) - 1; i >= 0; i-- {
				if disk.Blocks[i] == fileID {
					filePos = disk.findFileStart(i)
					break
				}
			}

			if filePos == -1 {
				continue
			}

			fileSize := disk.findFileSize(filePos)

			// Find leftmost suitable free space
			var bestFreeSpace int = -1
			for i := 0; i < filePos; i++ {
				if disk.Blocks[i] == -1 {
					freeSize := disk.findFreeSpaceSize(i)
					if freeSize >= fileSize {
						bestFreeSpace = i
						break
					}
					i += freeSize - 1 // Skip to end of this free space
				}
			}

			// Move the file if we found suitable space
			if bestFreeSpace != -1 && bestFreeSpace < filePos {
				disk.moveWholeFile(filePos, bestFreeSpace, fileSize)
			}
		}
	}

	return disk.calculateChecksum()
}
