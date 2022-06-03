package talosEcsTests

import (
	"testing"

	"github.com/costaconrado/talosecs"
)

func TestDoubleInitNotAllowed(t *testing.T) {
	talosecs.Reset()

	talosecs.Init()

	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("Double init should fail")
		}
	}()

	talosecs.Init()
}
