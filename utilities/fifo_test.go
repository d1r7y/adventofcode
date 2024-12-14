/*
Copyright © 2021-2024 Cameron Esfahani
*/

package utilities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFifoInt(t *testing.T) {
	fifo := NewFIFO[int]()

	fifo.Push(1)
	fifo.Push(2)
	fifo.Push(3)

	assert.Equal(t, 1, fifo.Pop())
	assert.Equal(t, 2, fifo.Pop())
	assert.Equal(t, 3, fifo.Pop())

	fifo.PushList([]int{5, 6, 7, 8})

	assert.Equal(t, 5, fifo.Pop())
	assert.Equal(t, 6, fifo.Pop())
	assert.Equal(t, 7, fifo.Pop())
	assert.Equal(t, 8, fifo.Pop())
}

func TestFifoString(t *testing.T) {
	fifo := NewFIFO[string]()

	fifo.Push("1")
	fifo.Push("2")
	fifo.Push("3")

	assert.Equal(t, "1", fifo.Pop())
	assert.Equal(t, "2", fifo.Pop())
	assert.Equal(t, "3", fifo.Pop())

	fifo.PushList([]string{"5", "6", "7", "8"})

	assert.Equal(t, "5", fifo.Pop())
	assert.Equal(t, "6", fifo.Pop())
	assert.Equal(t, "7", fifo.Pop())
	assert.Equal(t, "8", fifo.Pop())
}

func TestFifoIsEmpty(t *testing.T) {
	fifo := NewFIFO[string]()
	assert.True(t, fifo.IsEmpty())

	fifo.Push("1")
	assert.False(t, fifo.IsEmpty())

	assert.Equal(t, "1", fifo.Pop())
	assert.True(t, fifo.IsEmpty())
}
