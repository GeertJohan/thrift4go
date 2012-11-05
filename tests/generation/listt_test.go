package listt

import (
	"testing"
)

func TestContainerOfEnumsNew(t *testing.T) {
	emission := NewWrappedLists()

	if emission == nil {
		t.Errorf("NewWrappedLists emitted nil, not the struct.")
	}
}
