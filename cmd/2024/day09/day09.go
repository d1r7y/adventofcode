/*
Copyright © 2021-2024 Cameron Esfahani
*/

package TwentyTwentyFour_day09

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/d1r7y/adventofcode/utilities"
	"github.com/spf13/cobra"
)

// Day09Cmd represents the day09 command
var Day09Cmd = &cobra.Command{
	Use:   "day09",
	Short: `Disk Fragmenter`,
	Run: func(cmd *cobra.Command, args []string) {
		df, err := os.Open(utilities.GetInputPath(cmd))
		if err != nil {
			log.Fatal(err)
		}

		defer df.Close()

		fileContent, err := io.ReadAll(df)
		if err != nil {
			log.Fatal(err)
		}

		if fileContent != nil {
			err = day(string(fileContent))
			if err != nil {
				log.Fatal(err)
			}
		}
	},
}

type BlockList []int

type File struct {
	ID     int
	Size   int
	Blocks BlockList
}

type Disk struct {
	Size  int
	Files []File

	BlockAllocations BlockList
}

const (
	Unallocated = -1
)

func (d *Disk) Describe() string {
	str := ""

	for _, b := range d.BlockAllocations {
		if b == Unallocated {
			str += "."
		} else {
			str += fmt.Sprintf("%d", b)
		}
	}

	return str
}

func (d *Disk) CompactBlocks() {

	getNextFreeBlock := func(currentFreeBlock int) (bool, int) {
		if currentFreeBlock == -1 {
			currentFreeBlock = 0
		} else {
			currentFreeBlock++
		}

		for ; currentFreeBlock < len(d.BlockAllocations)-1; currentFreeBlock++ {
			if d.BlockAllocations[currentFreeBlock] == Unallocated {
				return true, currentFreeBlock
			}
		}

		return false, currentFreeBlock
	}

	getFirstFreeBlock := func() (bool, int) {
		return getNextFreeBlock(-1)
	}

	allocateBlock := func(block int, id int) {
		d.BlockAllocations[block] = id
	}

	freeBlock := func(block int) {
		d.BlockAllocations[block] = Unallocated
	}

	valid, currentFreeBlock := getFirstFreeBlock()
	if !valid {
		return
	}

	// Move files in reverse order.  Never have to move the first file.
	for fileIndex := len(d.Files) - 1; fileIndex > 0; fileIndex-- {
		file := d.Files[fileIndex]

		// Move blocks in reverse order.
		for blockIndex := len(file.Blocks) - 1; blockIndex >= 0; blockIndex-- {
			// If currentFreeBlock is greater than the current file block, then we're done
			if currentFreeBlock > file.Blocks[blockIndex] {
				return
			}

			// Allocate the disk block.
			allocateBlock(currentFreeBlock, file.ID)

			// Free the current file block
			freeBlock(file.Blocks[blockIndex])

			// Update to the new file block.
			file.Blocks[blockIndex] = currentFreeBlock

			valid, currentFreeBlock = getNextFreeBlock(currentFreeBlock)

			if !valid {
				return
			}
		}
	}
}

func (d *Disk) CompactFiles() {

	type Range struct {
		Start  int
		Length int
	}

	getNextFreeRange := func(currentRange Range) (bool, Range) {
		var start int

		if currentRange.Start == -1 {
			start = 0
		} else {
			start = currentRange.Start + currentRange.Length
		}

		r := Range{}

		foundFreeBlock := false

		for ; start < len(d.BlockAllocations)-1; start++ {
			if d.BlockAllocations[start] == Unallocated {
				if !foundFreeBlock {
					foundFreeBlock = true
					r.Start = start
				}
				r.Length++
			} else {
				if foundFreeBlock {
					return true, r
				}
			}
		}

		if foundFreeBlock {
			return true, r
		}

		return false, r
	}

	getFirstFreeRange := func() (bool, Range) {
		r := Range{-1, 0}
		return getNextFreeRange(r)
	}

	findFreeRange := func(size int) (bool, Range) {
		for ok, r := getFirstFreeRange(); ok; {
			if size <= r.Length {
				r.Length = size
				return true, r
			}

			ok, r = getNextFreeRange(r)
		}

		return false, Range{}
	}

	allocateBlocks := func(r Range, id int) {
		for b := r.Start; b < r.Start+r.Length; b++ {
			d.BlockAllocations[b] = id
		}
	}

	freeBlocks := func(r Range) {
		for b := r.Start; b < r.Start+r.Length; b++ {
			d.BlockAllocations[b] = Unallocated
		}
	}

	// Move files in reverse order.  Never have to move the first file.
	for fileIndex := len(d.Files) - 1; fileIndex > 0; fileIndex-- {
		file := d.Files[fileIndex]

		ok, r := findFreeRange(file.Size)
		if !ok {
			continue
		}

		// If currentFreeBlock is greater than the current file block, then we're done
		if r.Start > file.Blocks[0] {
			continue
		}

		// Allocate the disk blocks.
		allocateBlocks(r, file.ID)

		// Free the current file blocks
		fileRange := Range{file.Blocks[0], file.Size}
		freeBlocks(fileRange)

		// Update to the new file blocks.
		for fb := 0; fb < file.Size; fb++ {
			file.Blocks[fb] = r.Start + fb
		}
	}
}

func (d *Disk) CalculateChecksum() int {
	checksum := 0

	for pos, id := range d.BlockAllocations {
		if id != Unallocated {
			checksum += pos * id
		}
	}

	return checksum
}

func ParseDisk(fileContent string) *Disk {
	disk := &Disk{}
	disk.Files = make([]File, 0)
	disk.BlockAllocations = make(BlockList, 0)

	diskSize := 0
	id := 0

	for i := 0; i < len(fileContent); {
		allocationCount := int(fileContent[i]) - int('0')
		freeCount := 0

		// Handle last item specially: there's no free space after the file allocation.
		if i < len(fileContent)-1 {
			freeCount = int(fileContent[i+1]) - int('0')
		}

		file := File{
			ID:     id,
			Size:   0,
			Blocks: make(BlockList, 0),
		}

		for b := 0; b < allocationCount; b++ {
			// Allocate the disk blocks to the file.
			file.Blocks = append(file.Blocks, diskSize+b)
			file.Size++
		}

		for b := 0; b < allocationCount; b++ {
			// Remember which file owns the disk block.
			disk.BlockAllocations = append(disk.BlockAllocations, id)
		}

		// Add unallocated range
		for b := 0; b < freeCount; b++ {
			disk.BlockAllocations = append(disk.BlockAllocations, Unallocated)
		}

		id++

		i += 2

		diskSize += file.Size + freeCount

		disk.Files = append(disk.Files, file)
	}

	disk.Size = diskSize

	return disk
}

func day(fileContents string) error {
	// Part 1: While The Historians quickly figure out how to pilot these things, you notice an amphipod
	// in the corner struggling with his computer. He's trying to make more contiguous free space by
	// compacting all of the files, but his program isn't working; you offer to help.
	//
	// The disk map uses a dense format to represent the layout of files and free space on the disk. The
	// digits alternate between indicating the length of a file and the length of free space.
	//
	// So, a disk map like 12345 would represent a one-block file, two blocks of free space, a three-block
	// file, four blocks of free space, and then a five-block file. A disk map like 90909 would represent
	// three nine-block files in a row (with no free space between them).
	//
	// Each file on disk also has an ID number based on the order of the files as they appear before they are
	// rearranged, starting with ID 0. So, the disk map 12345 has three files: a one-block file with ID 0, a
	// three-block file with ID 1, and a five-block file with ID 2.
	//
	// The amphipod would like to move file blocks one at a time from the end of the disk to the leftmost free
	// space block (until there are no gaps remaining between file blocks).
	//
	// The final step of this file-compacting process is to update the filesystem checksum. To calculate the
	// checksum, add up the result of multiplying each of these blocks' position with the file ID number it
	// contains. The leftmost block is in position 0. If a block contains free space, skip it instead.
	//
	// Compact the amphipod's hard drive using the process he requested. What is the resulting filesystem checksum?

	disk := ParseDisk(fileContents)
	disk.CompactBlocks()
	checksum := disk.CalculateChecksum()

	fmt.Printf("Filesystem checksum after block compaction: %d\n", checksum)

	// Part 2: Upon completion, two things immediately become clear. First, the disk definitely has a lot more
	// contiguous free space, just like the amphipod hoped. Second, the computer is running much more slowly!
	// Maybe introducing all of that file system fragmentation was a bad idea?
	//
	// The eager amphipod already has a new plan: rather than move individual blocks, he'd like to try compacting
	// the files on his disk by moving whole files instead.
	//
	// This time, attempt to move whole files to the leftmost span of free space blocks that could fit the file.
	// Attempt to move each file exactly once in order of decreasing file ID number starting with the file with the
	// highest file ID number. If there is no span of free space to the left of a file that is large enough to fit the
	// file, the file does not move.

	disk2 := ParseDisk(fileContents)
	disk2.CompactFiles()
	checksum2 := disk2.CalculateChecksum()

	fmt.Printf("Filesystem checksum after file compaction: %d\n", checksum2)

	return nil
}
