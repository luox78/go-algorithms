package arrays_test

import (
	"testing"

	"github.com/wenjiax/go-algorithms/arrays"
)

func TestOrderedArray(t *testing.T) {
	arr := arrays.NewOrderedArray(4)

	// Len: 0, Cap: 4, Data: []
	t.Log(arr)

	// Len: 1, Cap: 4, Data: [5]
	arr.Add(5)
	t.Log(arr)

	// Len: 2, Cap: 4, Data: [4 5]
	arr.Add(4)
	t.Log(arr)

	// Len: 3, Cap: 4, Data: [4 5 6]
	arr.Add(6)
	t.Log(arr)

	// Len: 4, Cap: 4, Data: [4 4 5 6]
	arr.Add(4)
	t.Log(arr)

	// Len: 4, Cap: 4, Data: [4 4 5 6]
	arr.Add(1)
	t.Log(arr)

	// Len: 3, Cap: 4, Data: [4 4 5]
	arr.Del(3)
	t.Log(arr)

	// 4 <nil>
	t.Log(arr.Get(0))

	// Len: 3, Cap: 4, Data: [4 5 10]
	arr.Update(1, 10)
	t.Log(arr)
}
