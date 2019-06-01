package hashtable_test

import (
	"testing"

	"github.com/wenjiax/go-algorithms/hashtable"
)

func TestListHashTable(t *testing.T) {
	h := hashtable.NewListHashTable(2)

	h.Set("a", 1)
	h.Set("b", 2)
	h.Set("c", 3)
	h.Set("d", 4)
	h.Set("e", 5)
	h.Set("f", 6)
	h.Set("g", 7)

	t.Log(h.Get("a"))
	t.Log(h.Get("b"))
	t.Log(h.Get("c"))
	t.Log(h.Get("d"))
	t.Log(h.Get("e"))
	t.Log(h.Get("f"))
	t.Log(h.Get("g"))
}
