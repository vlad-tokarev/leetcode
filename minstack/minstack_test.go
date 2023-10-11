package minstack

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Push(val int) func(t *testing.T, stack *MinStack) {
	return func(t *testing.T, stack *MinStack) {
		stack.Push(val)
	}
}

func Pop() func(t *testing.T, stack *MinStack) {
	return func(t *testing.T, stack *MinStack) {
		stack.Pop()
	}
}

func GetMin() func(t *testing.T, stack *MinStack) {
	return func(t *testing.T, stack *MinStack) {
		stack.GetMin()
	}
}
func Top() func(t *testing.T, stack *MinStack) {
	return func(t *testing.T, stack *MinStack) {
		stack.Top()
	}
}

func AssertTop(val int) func(t *testing.T, stack *MinStack) {
	return func(t *testing.T, stack *MinStack) {
		assert.Equal(t, val, stack.Top())
	}
}

func AssertMin(val int) func(t *testing.T, stack *MinStack) {
	return func(t *testing.T, stack *MinStack) {
		assert.Equal(t, val, stack.GetMin())
	}
}

func TestStack(t *testing.T) {

	tests := [][]func(t *testing.T, stack *MinStack){
		{
			Push(-2),
			Push(0),
			Push(-3),
			AssertMin(-3),
			Pop(),
			AssertTop(0),
			AssertMin(-2),
		},
		{
			Push(-2),
			Push(0),
			Push(-1),
			AssertMin(-2),
			AssertTop(-1),
			Pop(),
			AssertMin(-2),
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("case_%d", i), func(t *testing.T) {
			stack := Constructor()
			for i, op := range tt {
				_ = i
				op(t, &stack)
			}
		})
	}
}

func TestHeap(t *testing.T) {
	h := FastRemovalHeap{
		data:       make([]int, 0),
		refToIndex: make(map[int]int),
		indexToRef: make(map[int]int),
	}

	h.Insert(10)
	h.Insert(20)
	h.Insert(5)
	h.Insert(150)
	ref300_1 := h.Insert(-300)
	ref300_2 := h.Insert(-300)

	assert.Equal(t, -300, h.GetMin())
	t.Logf("data: %v", h.data)

	h.Delete(ref300_1)
	assert.Equal(t, -300, h.GetMin())
	t.Logf("data: %v", h.data)

	h.Delete(ref300_2)
	assert.Equal(t, 5, h.GetMin())
	t.Logf("data: %v", h.data)

}
