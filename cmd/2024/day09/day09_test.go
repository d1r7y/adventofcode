/*
Copyright © 2021-2024 Cameron Esfahani
*/

package TwentyTwentyFour_day09

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseDisk(t *testing.T) {
	type testCase struct {
		text         string
		expectedDisk *Disk
	}
	testCases := []testCase{
		{
			text: "12345",
			expectedDisk: &Disk{
				Size:             15,
				BlockAllocations: BlockList{0, Unallocated, Unallocated, 1, 1, 1, Unallocated, Unallocated, Unallocated, Unallocated, 2, 2, 2, 2, 2},
				Files: []File{
					{
						ID:     0,
						Size:   1,
						Blocks: BlockList{0},
					},
					{
						ID:     1,
						Size:   3,
						Blocks: BlockList{3, 4, 5},
					},
					{
						ID:     2,
						Size:   5,
						Blocks: BlockList{10, 11, 12, 13, 14},
					},
				},
			},
		},
		{
			text: "2333133121414131402",
			expectedDisk: &Disk{
				Size: 42,
				BlockAllocations: BlockList{
					0, 0, Unallocated, Unallocated, Unallocated, 1, 1, 1, Unallocated, Unallocated, Unallocated, 2, Unallocated, Unallocated, Unallocated,
					3, 3, 3, Unallocated, 4, 4, Unallocated, 5, 5, 5, 5, Unallocated, 6, 6, 6, 6, Unallocated, 7, 7, 7, Unallocated, 8, 8, 8, 8, 9, 9,
				},
				Files: []File{
					{
						ID:     0,
						Size:   2,
						Blocks: BlockList{0, 1},
					},
					{
						ID:     1,
						Size:   3,
						Blocks: BlockList{5, 6, 7},
					},
					{
						ID:     2,
						Size:   1,
						Blocks: BlockList{11},
					},
					{
						ID:     3,
						Size:   3,
						Blocks: BlockList{15, 16, 17},
					},
					{
						ID:     4,
						Size:   2,
						Blocks: BlockList{19, 20},
					},
					{
						ID:     5,
						Size:   4,
						Blocks: BlockList{22, 23, 24, 25},
					},
					{
						ID:     6,
						Size:   4,
						Blocks: BlockList{27, 28, 29, 30},
					},
					{
						ID:     7,
						Size:   3,
						Blocks: BlockList{32, 33, 34},
					},
					{
						ID:     8,
						Size:   4,
						Blocks: BlockList{36, 37, 38, 39},
					},
					{
						ID:     9,
						Size:   2,
						Blocks: BlockList{40, 41},
					},
				},
			},
		},
	}

	for _, test := range testCases {
		disk := ParseDisk(test.text)
		assert.Equal(t, test.expectedDisk, disk)
	}
}

func TestDescribe(t *testing.T) {
	type testCase struct {
		text                string
		expectedDescription string
	}
	testCases := []testCase{
		{
			text:                "12345",
			expectedDescription: "0..111....22222",
		},
		{
			text:                "2333133121414131402",
			expectedDescription: "00...111...2...333.44.5555.6666.777.888899",
		},
	}

	for _, test := range testCases {
		disk := ParseDisk(test.text)
		assert.Equal(t, test.expectedDescription, disk.Describe())
	}
}

func TestCompactBlocks(t *testing.T) {
	type testCase struct {
		text                string
		expectedDescription string
	}
	testCases := []testCase{
		{
			text:                "12345",
			expectedDescription: "022111222......",
		},
		{
			text:                "2333133121414131402",
			expectedDescription: "0099811188827773336446555566..............",
		},
	}

	for _, test := range testCases {
		disk := ParseDisk(test.text)
		disk.CompactBlocks()
		assert.Equal(t, test.expectedDescription, disk.Describe())
	}
}

func TestCompactFiles(t *testing.T) {
	type testCase struct {
		text                string
		expectedDescription string
	}
	testCases := []testCase{
		{
			text:                "12345",
			expectedDescription: "0..111....22222",
		},
		{
			text:                "2333133121414131402",
			expectedDescription: "00992111777.44.333....5555.6666.....8888..",
		},
	}

	for _, test := range testCases {
		disk := ParseDisk(test.text)
		disk.CompactFiles()
		assert.Equal(t, test.expectedDescription, disk.Describe())
	}
}

func TestCalculateChecksumBlocks(t *testing.T) {
	type testCase struct {
		text             string
		expectedChecksum int
	}
	testCases := []testCase{
		{
			text:             "12345",
			expectedChecksum: 60,
		},
		{
			text:             "2333133121414131402",
			expectedChecksum: 1928,
		},
	}

	for _, test := range testCases {
		disk := ParseDisk(test.text)
		disk.CompactBlocks()
		assert.Equal(t, test.expectedChecksum, disk.CalculateChecksum())
	}
}

func TestCalculateChecksumFiles(t *testing.T) {
	type testCase struct {
		text             string
		expectedChecksum int
	}
	testCases := []testCase{
		{
			text:             "12345",
			expectedChecksum: 132,
		},
		{
			text:             "2333133121414131402",
			expectedChecksum: 2858,
		},
	}

	for _, test := range testCases {
		disk := ParseDisk(test.text)
		disk.CompactFiles()
		assert.Equal(t, test.expectedChecksum, disk.CalculateChecksum())
	}
}
