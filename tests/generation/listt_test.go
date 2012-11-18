package listt

import (
	"testing"
)

func TestStructWithListsNew(t *testing.T) {
	emission := NewStructWithLists()

	if emission == nil {
		t.Errorf("NewStructWithLists emitted nil, not the struct.")
	}
}

func TestStructWithListsPushAndIter(t *testing.T) {
	sl := NewStructWithLists()

	// test i32 push and iter
	// TODO(GeertJohan): add int32 max and min.
	i32Values := []int32{-100, -42, -4, -2, -1, 0, 1, 2, 4, 42, 100}
	for _, i32 := range i32Values {
		sl.ListInt32.Push(i32)
	}

	i32Iter := sl.ListInt32.Iter()
	var read int32
	for num, expected := range i32Values {
		select {
		case read = <-i32Iter:
			if expected != read {
				t.Errorf("i32 push/iter failed on #%d. Expected: %d. Received: %d.", num, expected, read)
			}
			break
		default:
			t.Errorf("Failed to iterate over ListInt32. Expected another value but iter chan is empty. At #%d.", num)
			break
		}
	}
}
