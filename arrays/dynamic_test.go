package arrays_test

import (
	"testing"

	"github.com/wenjiax/go-algorithms/arrays"
)

func TestLazyInit(t *testing.T) {
	arr := &arrays.DynamicArray{}

	// Len: 1, Cap: 4, Data: [2]
	arr.InsertFirst(2)
	t.Log(arr)

	// Len: 2, Cap: 4, Data: [1 2]
	arr.InsertFirst(1)
	t.Log(arr)

	arr1 := new(arrays.DynamicArray)

	// <nil>
	t.Log(arr1.First())

	// Len: 0, Cap: 4, Data: []
	t.Log(arr1)
}

func TestDynamicArray(t *testing.T) {
	arr := arrays.NewDynamicArray(0)

	// Len: 0, Cap: 4, Data: []
	t.Log(arr)

	// <nil>
	t.Log(arr.First())
	t.Log(arr.Last())

	// Len: 2, Cap: 4, Data: [2 1]
	arr.InsertFirst(1)
	arr.InsertFirst(2)
	t.Log(arr)

	// Len: 4, Cap: 4, Data: [2 1 4 3]
	arr.InsertLast(4)
	arr.InsertLast(3)
	t.Log(arr)

	// Len: 5, Cap: 8, Data: [2 5 1 4 3]
	arr.Insert(1, 5)
	t.Log(arr)

	// Len: 5, Cap: 8, Data: [2 5 6 4 3]
	arr.Update(2, 6)
	t.Log(arr)

	// 2
	t.Log(arr.First())

	// 3
	t.Log(arr.Last())

	// 4
	t.Log(arr.Get(3))

	// Len: 4, Cap: 4, Data: [2 6 4 3]
	arr.Remove(1)
	t.Log(arr)

	// Len: 3, Cap: 4, Data: [2 4 3]
	arr.Remove(1)
	t.Log(arr)
}
